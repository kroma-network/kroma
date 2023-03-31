package actions

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"

	"github.com/wemixkanvas/kanvas/bindings/bindings"
	"github.com/wemixkanvas/kanvas/components/node/eth"
	"github.com/wemixkanvas/kanvas/components/node/sources"
	"github.com/wemixkanvas/kanvas/components/node/testlog"
	chal "github.com/wemixkanvas/kanvas/components/validator/challenge"
	"github.com/wemixkanvas/kanvas/e2e/e2eutils"
)

func TestChallenger(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	dp.DeployConfig.FinalizationPeriodSeconds = 60 * 60 * 24
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlDebug)
	miner, propEngine, proposer := setupProposerTest(t, sd, log)

	rollupPropCl := proposer.RollupClient()
	batcher := NewL2Batcher(log, sd.RollupCfg, &BatcherCfg{
		MinL1TxSize: 0,
		MaxL1TxSize: 128_000,
		BatcherKey:  dp.Secrets.Batcher,
	}, rollupPropCl, miner.EthClient(), propEngine.EthClient())

	// setup mockup rpc for returning invalid output
	mockRPC := e2eutils.NewRPC(proposer.RPCClient())
	validator := NewL2Validator(t, log, &ValidatorCfg{
		OutputOracleAddr:  sd.DeploymentsL1.L2OutputOracleProxy,
		ValidatorKey:      dp.Secrets.Validator,
		AllowNonFinalized: false,
	}, miner.EthClient(), sources.NewRollupClient(mockRPC))

	asserter := NewL2Challenger(t, log, &ValidatorCfg{
		OutputOracleAddr:  sd.DeploymentsL1.L2OutputOracleProxy,
		ColosseumAddr:     sd.DeploymentsL1.ColosseumProxy,
		ValidatorKey:      dp.Secrets.Validator,
		AllowNonFinalized: false,
	}, miner.EthClient(), sources.NewRollupClient(mockRPC))

	challenger := NewL2Challenger(t, log, &ValidatorCfg{
		OutputOracleAddr:  sd.DeploymentsL1.L2OutputOracleProxy,
		ColosseumAddr:     sd.DeploymentsL1.ColosseumProxy,
		ValidatorKey:      dp.Secrets.Validator,
		AllowNonFinalized: false,
	}, miner.EthClient(), proposer.RollupClient())

	// L1 block
	miner.ActEmptyBlock(t)
	// L2 block
	proposer.ActL1HeadSignal(t)
	proposer.ActL2PipelineFull(t)
	proposer.ActBuildToL1Head(t)
	// submit and include in L1
	batcher.ActSubmitAll(t)
	includeL1Block(t, miner, dp.Addresses.Batcher)
	// finalize the first and second L1 blocks, including the batch
	miner.ActL1SafeNext(t)
	miner.ActL1SafeNext(t)
	miner.ActL1FinalizeNext(t)
	miner.ActL1FinalizeNext(t)
	// derive and see the L2 chain fully finalize
	proposer.ActL2PipelineFull(t)
	proposer.ActL1SafeSignal(t)
	proposer.ActL1FinalizedSignal(t)

	common.BytesToHash([]byte{})

	require.Equal(t, proposer.SyncStatus().UnsafeL2, proposer.SyncStatus().FinalizedL2)
	// create l2 output submission transactions until there is nothing left to submit
	for validator.CanSubmit(t) {
		// and submit it to L1
		validator.ActSubmitL2Output(t)
		// include output on L1
		includeL1Block(t, miner, dp.Addresses.Validator)
		// Check submission was successful
		receipt, err := miner.EthClient().TransactionReceipt(t.Ctx(), validator.LastSubmitL2OutputTx())
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "submission failed")
	}

	// check that L1 stored the expected output root
	outputOracleContract, err := bindings.NewL2OutputOracle(sd.DeploymentsL1.L2OutputOracleProxy, miner.EthClient())
	require.NoError(t, err)
	block := proposer.SyncStatus().FinalizedL2

	outputOnL1, err := outputOracleContract.GetL2OutputAfter(nil, new(big.Int).SetUint64(block.Number))
	require.NoError(t, err)
	require.Less(t, block.Time, outputOnL1.Timestamp.Uint64(), "output is registered with L1 timestamp of L2 tx output submission, past L2 block")
	outputComputed, err := proposer.RollupClient().OutputAtBlock(t.Ctx(), block.Number)
	require.NoError(t, err)
	require.NotEqual(t, eth.Bytes32(outputOnL1.OutputRoot), outputComputed.OutputRoot, "output roots must different")

	// submit create challenge tx
	txHash := challenger.ActCreateChallenge(t)

	// include tx on L1
	includeL1Block(t, miner, dp.Addresses.Validator)

	// Check whether the submission was successful
	receipt, err := miner.EthClient().TransactionReceipt(t.Ctx(), txHash)
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "failed to create challenge")

	// check challenge created
	colosseumContract, err := bindings.NewColosseum(sd.DeploymentsL1.ColosseumProxy, miner.EthClient())
	require.NoError(t, err)
	challengeId, err := colosseumContract.LatestChallengeId(nil)
	require.NoError(t, err)
	require.True(t, challengeId.Cmp(big.NewInt(0)) == 1, "challenge not found")

interaction:
	for {
		status, err := colosseumContract.GetStatusInProgress(nil)
		require.NoError(t, err)

		sender := dp.Addresses.Validator

		switch status {
		case chal.StatusChallengerTurn:
			// call bisect by challenger
			txHash = challenger.ActBisect(t)
		case chal.StatusAsserterTurn:
			sender = dp.Addresses.Validator
			challenge, err := colosseumContract.GetChallengeInProgress(nil)
			require.NoError(t, err)
			mockRPC.SetSegmentStart(challenge.SegStart)
			// call bisect by validator
			txHash = asserter.ActBisect(t)
		case chal.StatusAsserterTimeout:
			txHash = challenger.ActTimeout(t)
		case chal.StatusProveReady:
			txHash = challenger.ActProveFault(t)
		default:
			break interaction
		}

		// include tx on L1
		includeL1Block(t, miner, sender)

		// Check whether the submission was successful
		receipt, err = miner.EthClient().TransactionReceipt(t.Ctx(), txHash)
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "failed to progress interactive proof")
	}

	_, err = outputOracleContract.GetL2OutputAfter(nil, new(big.Int).SetUint64(block.Number))
	require.ErrorContains(t, err,
		"cannot get output for a block that has not been submitted",
		"output not deleted",
	)
}
