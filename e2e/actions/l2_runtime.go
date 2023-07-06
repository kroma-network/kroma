package actions

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/bindings/bindings"
	"github.com/kroma-network/kroma/components/node/eth"
	"github.com/kroma-network/kroma/components/node/sources"
	"github.com/kroma-network/kroma/components/node/testlog"
	"github.com/kroma-network/kroma/e2e"
	"github.com/kroma-network/kroma/e2e/e2eutils"
)

const defaultDepositAmount = 1_000

type Runtime struct {
	t                        StatefulTesting
	l                        log.Logger
	sd                       *e2eutils.SetupData
	dp                       *e2eutils.DeployParams
	miner                    *L1Miner
	propEngine               *L2Engine
	proposer                 *L2Proposer
	batcher                  *L2Batcher
	validator                *L2Validator
	challenger               *L2Validator
	guardian                 *L2Validator
	outputOracleContract     *bindings.L2OutputOracle
	colosseumContract        *bindings.Colosseum
	valPoolContract          *bindings.ValidatorPoolCaller
	targetInvalidBlockNumber uint64
	outputIndex              *big.Int
	outputOnL1               bindings.TypesCheckpointOutput
	txHash                   common.Hash
	receipt                  *types.Receipt
}

func defaultRuntime(gt *testing.T) Runtime {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	dp.DeployConfig.FinalizationPeriodSeconds = 60 * 60 * 24
	dp.DeployConfig.ColosseumDummyHash = common.HexToHash(e2e.DummyHashSepolia)
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	l := testlog.Logger(t, log.LvlDebug)
	rt := Runtime{
		t:  t,
		dp: dp,
		sd: sd,
		l:  l,
	}
	rt.miner, rt.propEngine, rt.proposer = setupProposerTest(rt.t, rt.sd, rt.l)
	rt.setupBatcher()

	return rt
}

func (rt *Runtime) setupBatcher() {
	rollupPropCl := rt.proposer.RollupClient()
	batcher := NewL2Batcher(rt.l, rt.sd.RollupCfg, &BatcherCfg{
		MinL1TxSize: 0,
		MaxL1TxSize: 128_000,
		BatcherKey:  rt.dp.Secrets.Batcher,
	}, rollupPropCl, rt.miner.EthClient(), rt.propEngine.EthClient())
	rt.batcher = batcher
}

func (rt *Runtime) setTargetInvalidBlockNumber(targetInvalidBlockNumber uint64) {
	rt.targetInvalidBlockNumber = targetInvalidBlockNumber
}

func (rt *Runtime) setupHonestValidator() {
	rt.validator = rt.honestValidator(rt.dp.Secrets.TrustedValidator)
}

func (rt *Runtime) setupMaliciousValidator() {
	rt.validator = rt.maliciousValidator(rt.dp.Secrets.TrustedValidator)
}

func (rt *Runtime) setupHonestChallenger() {
	rt.challenger = rt.honestValidator(rt.dp.Secrets.Challenger)
}

func (rt *Runtime) setupMaliciousChallenger() {
	rt.challenger = rt.maliciousValidator(rt.dp.Secrets.Challenger)
}

func (rt *Runtime) setupHonestGuardian() {
	rt.guardian = rt.honestValidator(rt.dp.Secrets.Challenger)
}

func (rt *Runtime) setupMaliciousGuardian() {
	rt.guardian = rt.maliciousValidator(rt.dp.Secrets.Challenger)
}

func (rt *Runtime) honestValidator(pk *ecdsa.PrivateKey) *L2Validator {
	// setup mockup rpc for returning valid output
	validatorRPC := e2eutils.NewHonestL2RPC(rt.proposer.RPCClient())
	validatorRollupClient := sources.NewRollupClient(validatorRPC)
	validator := NewL2Validator(rt.t, rt.l, &ValidatorCfg{
		OutputOracleAddr:    rt.sd.DeploymentsL1.L2OutputOracleProxy,
		ValidatorPoolAddr:   rt.sd.DeploymentsL1.ValidatorPoolProxy,
		ColosseumAddr:       rt.sd.DeploymentsL1.ColosseumProxy,
		SecurityCouncilAddr: rt.sd.DeploymentsL1.SecurityCouncilProxy,
		ValidatorKey:        pk,
		AllowNonFinalized:   false,
	}, rt.miner.EthClient(), validatorRollupClient)
	validatorRPC.SetTargetBlockNumber(rt.targetInvalidBlockNumber)
	return validator
}

func (rt *Runtime) maliciousValidator(pk *ecdsa.PrivateKey) *L2Validator {
	// setup mockup rpc for returning invalid output
	validatorRPC := e2eutils.NewMaliciousL2RPC(rt.proposer.RPCClient())
	validatorRollupClient := sources.NewRollupClient(validatorRPC)
	validator := NewL2Validator(rt.t, rt.l, &ValidatorCfg{
		OutputOracleAddr:    rt.sd.DeploymentsL1.L2OutputOracleProxy,
		ValidatorPoolAddr:   rt.sd.DeploymentsL1.ValidatorPoolProxy,
		ColosseumAddr:       rt.sd.DeploymentsL1.ColosseumProxy,
		SecurityCouncilAddr: rt.sd.DeploymentsL1.SecurityCouncilProxy,
		ValidatorKey:        pk,
		AllowNonFinalized:   false,
	}, rt.miner.EthClient(), validatorRollupClient)
	validatorRPC.SetTargetBlockNumber(rt.targetInvalidBlockNumber)
	return validator
}

func (rt *Runtime) bindChallengeContracts() {
	var err error
	// bind contracts
	rt.outputOracleContract, err = bindings.NewL2OutputOracle(rt.sd.DeploymentsL1.L2OutputOracleProxy, rt.miner.EthClient())
	require.NoError(rt.t, err)

	rt.colosseumContract, err = bindings.NewColosseum(rt.sd.DeploymentsL1.ColosseumProxy, rt.miner.EthClient())
	require.NoError(rt.t, err)

	rt.valPoolContract, err = bindings.NewValidatorPoolCaller(rt.sd.DeploymentsL1.ValidatorPoolProxy, rt.miner.EthClient())
	require.NoError(rt.t, err)
}

// setup challenge between asserter and challenger
func (rt *Runtime) setupChallenge() {
	// NOTE(chokobole): It is necessary to wait for one finalized (or safe if AllowNonFinalized
	// config is set) block to pass after each submission interval before submitting the output
	// root. For example, if the submission interval is set to 1800 blocks, the output root can
	// only be submitted at 1801 finalized blocks. In fact, the following code is designed to
	// create one or more finalized L2 blocks in order to pass the test. If Proto Dank Sharding
	// is introduced, the below code fix may no longer be necessary.
	for i := 0; i < 3; i++ {
		// L1 block
		rt.miner.ActEmptyBlock(rt.t)
		// L2 block
		rt.proposer.ActL1HeadSignal(rt.t)
		rt.proposer.ActL2PipelineFull(rt.t)
		rt.proposer.ActBuildToL1Head(rt.t)
		// submit and include in L1
		rt.batcher.ActSubmitAll(rt.t)
		rt.miner.includeL1Block(rt.t, rt.dp.Addresses.Batcher)
		// finalize the first and second L1 blocks, including the batch
		rt.miner.ActL1SafeNext(rt.t)
		rt.miner.ActL1SafeNext(rt.t)
		rt.miner.ActL1FinalizeNext(rt.t)
		rt.miner.ActL1FinalizeNext(rt.t)
		// derive and see the L2 chain fully finalize
		rt.proposer.ActL2PipelineFull(rt.t)
		rt.proposer.ActL1SafeSignal(rt.t)
		rt.proposer.ActL1FinalizedSignal(rt.t)
	}

	// deposit bond for validator
	rt.validator.ActDeposit(rt.t, defaultDepositAmount)
	rt.miner.includeL1Block(rt.t, rt.validator.address)

	// check validator balance increased
	bal, err := rt.valPoolContract.BalanceOf(nil, rt.validator.address)
	require.NoError(rt.t, err)
	require.Equal(rt.t, new(big.Int).SetUint64(defaultDepositAmount), bal)

	require.Equal(rt.t, rt.proposer.SyncStatus().UnsafeL2, rt.proposer.SyncStatus().FinalizedL2)

	// create l2 output submission transactions until there is nothing left to submit
	for rt.validator.CanSubmit(rt.t) {
		// and submit it to L1
		rt.validator.ActSubmitL2Output(rt.t)
		// include output on L1
		rt.miner.includeL1Block(rt.t, rt.validator.address)
		// Check submission was successful
		receipt, err := rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), rt.validator.LastSubmitL2OutputTx())
		require.NoError(rt.t, err)
		require.Equal(rt.t, types.ReceiptStatusSuccessful, receipt.Status, "submission failed")
	}

	// check that L1 stored the expected output root
	// NOTE(chokobole): Comment these 2 lines because of the reason above.
	// If Proto Dank Sharding is introduced, the below code fix may be restored.
	// block := proposer.SyncStatus().FinalizedL2
	// outputOnL1, err := outputOracleContract.GetL2OutputAfter(nil, new(big.Int).SetUint64(block.Number))
	targetBlockNum := big.NewInt(int64(rt.targetInvalidBlockNumber))
	rt.outputIndex, err = rt.outputOracleContract.GetL2OutputIndexAfter(nil, targetBlockNum)
	require.NoError(rt.t, err)
	rt.outputOnL1, err = rt.outputOracleContract.GetL2OutputAfter(nil, targetBlockNum)
	require.NoError(rt.t, err)
	block, err := rt.propEngine.EthClient().BlockByNumber(rt.t.Ctx(), targetBlockNum)
	require.NoError(rt.t, err)
	require.Less(rt.t, block.Time(), rt.outputOnL1.Timestamp.Uint64(), "output is registered with L1 timestamp of L2 tx output submission, past L2 block")
	outputComputed, err := rt.proposer.RollupClient().OutputAtBlock(rt.t.Ctx(), targetBlockNum.Uint64())
	require.NoError(rt.t, err)
	require.NotEqual(rt.t, eth.Bytes32(rt.outputOnL1.OutputRoot), outputComputed.OutputRoot, "output roots must different")

	// deposit bond for challenger
	rt.challenger.ActDeposit(rt.t, defaultDepositAmount)
	rt.miner.includeL1Block(rt.t, rt.challenger.address)

	// check bond amount before create challenge
	bond, err := rt.valPoolContract.GetBond(nil, rt.outputIndex)
	require.NoError(rt.t, err)
	require.Equal(rt.t, rt.dp.DeployConfig.ValidatorPoolMinBondAmount.ToInt(), bond.Amount)

	// submit create challenge tx
	rt.txHash = rt.challenger.ActCreateChallenge(rt.t, rt.outputIndex)

	// include tx on L1
	rt.miner.includeL1Block(rt.t, rt.challenger.address)

	// Check whether the submission was successful
	rt.receipt, err = rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), rt.txHash)
	require.NoError(rt.t, err)
	require.Equal(rt.t, types.ReceiptStatusSuccessful, rt.receipt.Status, "failed to create challenge")

	// check challenge created
	challenge, err := rt.colosseumContract.GetChallenge(nil, rt.outputIndex)
	require.NoError(rt.t, err)
	require.NotNil(rt.t, challenge, "challenge not found")

	// check bond amount doubled after create challenge
	bond, err = rt.valPoolContract.GetBond(nil, rt.outputIndex)
	require.NoError(rt.t, err)
	require.Equal(rt.t, big.NewInt(2*rt.dp.DeployConfig.ValidatorPoolMinBondAmount.ToInt().Int64()), bond.Amount)

	// check challenger balance decreased
	cBal, err := rt.valPoolContract.BalanceOf(nil, rt.challenger.address)
	require.NoError(rt.t, err)
	require.Equal(rt.t, new(big.Int).Sub(new(big.Int).SetInt64(defaultDepositAmount), rt.dp.DeployConfig.ValidatorPoolMinBondAmount.ToInt()), cBal)
}
