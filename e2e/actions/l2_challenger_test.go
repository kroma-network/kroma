package actions

import (
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
	chal "github.com/kroma-network/kroma/components/validator/challenge"
	"github.com/kroma-network/kroma/e2e"
	"github.com/kroma-network/kroma/e2e/e2eutils"
	"github.com/kroma-network/kroma/e2e/testdata"
)

func TestChallenger(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	dp.DeployConfig.FinalizationPeriodSeconds = 60 * 60 * 24
	dp.DeployConfig.ColosseumDummyHash = common.HexToHash(e2e.DummyHashSepolia)
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	l := testlog.Logger(t, log.LvlDebug)
	miner, propEngine, proposer := setupProposerTest(t, sd, l)
	var validatorInitialAmount uint64 = 1_000
	var challengerInitialAmount uint64 = 1_000

	rollupPropCl := proposer.RollupClient()
	batcher := NewL2Batcher(l, sd.RollupCfg, &BatcherCfg{
		MinL1TxSize: 0,
		MaxL1TxSize: 128_000,
		BatcherKey:  dp.Secrets.Batcher,
	}, rollupPropCl, miner.EthClient(), propEngine.EthClient())

	// setup mockup rpc for returning invalid output
	validatorRPC := e2eutils.NewMaliciousL2RPC(proposer.RPCClient())
	validatorRollupClient := sources.NewRollupClient(validatorRPC)
	validator := NewL2Validator(t, l, &ValidatorCfg{
		OutputOracleAddr:  sd.DeploymentsL1.L2OutputOracleProxy,
		ValidatorPoolAddr: sd.DeploymentsL1.ValidatorPoolProxy,
		ColosseumAddr:     sd.DeploymentsL1.ColosseumProxy,
		ValidatorKey:      dp.Secrets.TrustedValidator,
		AllowNonFinalized: false,
	}, miner.EthClient(), validatorRollupClient)

	challengerRPC := e2eutils.NewHonestL2RPC(proposer.RPCClient())
	challengerRollupClient := sources.NewRollupClient(challengerRPC)
	challenger := NewL2Validator(t, l, &ValidatorCfg{
		OutputOracleAddr:  sd.DeploymentsL1.L2OutputOracleProxy,
		ValidatorPoolAddr: sd.DeploymentsL1.ValidatorPoolProxy,
		ColosseumAddr:     sd.DeploymentsL1.ColosseumProxy,
		ValidatorKey:      dp.Secrets.Challenger,
		AllowNonFinalized: false,
	}, miner.EthClient(), challengerRollupClient)

	guardianRPC := e2eutils.NewHonestL2RPC(proposer.RPCClient())
	guardianRollupClient := sources.NewRollupClient(guardianRPC)
	guardian := NewL2Validator(t, l, &ValidatorCfg{
		OutputOracleAddr:    sd.DeploymentsL1.L2OutputOracleProxy,
		SecurityCouncilAddr: sd.DeploymentsL1.SecurityCouncilProxy,
		ValidatorKey:        dp.Secrets.Challenger,
		AllowNonFinalized:   false,
	}, miner.EthClient(), guardianRollupClient)

	validatorRPC.SetTargetBlockNumber(testdata.TargetBlockNumber)
	challengerRPC.SetTargetBlockNumber(testdata.TargetBlockNumber)
	guardianRPC.SetTargetBlockNumber(testdata.TargetBlockNumber)

	// bind contracts
	outputOracleContract, err := bindings.NewL2OutputOracle(sd.DeploymentsL1.L2OutputOracleProxy, miner.EthClient())
	require.NoError(t, err)

	colosseumContract, err := bindings.NewColosseum(sd.DeploymentsL1.ColosseumProxy, miner.EthClient())
	require.NoError(t, err)

	valPoolContract, err := bindings.NewValidatorPoolCaller(sd.DeploymentsL1.ValidatorPoolProxy, miner.EthClient())
	require.NoError(t, err)

	// NOTE(chokobole): It is necessary to wait for one finalized (or safe if AllowNonFinalized
	// config is set) block to pass after each submission interval before submitting the output
	// root. For example, if the submission interval is set to 1800 blocks, the output root can
	// only be submitted at 1801 finalized blocks. In fact, the following code is designed to
	// create one or more finalized L2 blocks in order to pass the test. If Proto Dank Sharding
	// is introduced, the below code fix may no longer be necessary.
	for i := 0; i < 3; i++ {
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
	}

	// deposit bond for validator
	validator.ActDeposit(t, validatorInitialAmount)
	includeL1Block(t, miner, validator.address)

	// check validator balance increased
	bal, err := valPoolContract.BalanceOf(nil, validator.address)
	require.NoError(t, err)
	require.Equal(t, new(big.Int).SetUint64(validatorInitialAmount), bal)

	require.Equal(t, proposer.SyncStatus().UnsafeL2, proposer.SyncStatus().FinalizedL2)

	// create l2 output submission transactions until there is nothing left to submit
	for validator.CanSubmit(t) {
		// and submit it to L1
		validator.ActSubmitL2Output(t)
		// include output on L1
		includeL1Block(t, miner, validator.address)
		// Check submission was successful
		receipt, err := miner.EthClient().TransactionReceipt(t.Ctx(), validator.LastSubmitL2OutputTx())
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "submission failed")
	}

	// check that L1 stored the expected output root
	// NOTE(chokobole): Comment these 2 lines because of the reason above.
	// If Proto Dank Sharding is introduced, the below code fix may be restored.
	// block := proposer.SyncStatus().FinalizedL2
	// outputOnL1, err := outputOracleContract.GetL2OutputAfter(nil, new(big.Int).SetUint64(block.Number))
	targetBlockNum := big.NewInt(int64(testdata.TargetBlockNumber))
	outputIndex, err := outputOracleContract.GetL2OutputIndexAfter(nil, targetBlockNum)
	require.NoError(t, err)
	outputOnL1, err := outputOracleContract.GetL2OutputAfter(nil, targetBlockNum)
	require.NoError(t, err)
	block, err := propEngine.EthClient().BlockByNumber(t.Ctx(), targetBlockNum)
	require.NoError(t, err)
	require.Less(t, block.Time(), outputOnL1.Timestamp.Uint64(), "output is registered with L1 timestamp of L2 tx output submission, past L2 block")
	outputComputed, err := proposer.RollupClient().OutputAtBlock(t.Ctx(), targetBlockNum.Uint64())
	require.NoError(t, err)
	require.NotEqual(t, eth.Bytes32(outputOnL1.OutputRoot), outputComputed.OutputRoot, "output roots must different")

	// deposit bond for challenger
	challenger.ActDeposit(t, challengerInitialAmount)
	includeL1Block(t, miner, challenger.address)

	// check bond amount before create challenge
	bond, err := valPoolContract.GetBond(nil, outputIndex)
	require.NoError(t, err)
	require.Equal(t, 0, big.NewInt(dp.DeployConfig.ValidatorPoolMinBondAmount.ToInt().Int64()).Cmp(bond.Amount))

	// submit create challenge tx
	txHash := challenger.ActCreateChallenge(t, outputIndex)

	// include tx on L1
	includeL1Block(t, miner, challenger.address)

	// Check whether the submission was successful
	receipt, err := miner.EthClient().TransactionReceipt(t.Ctx(), txHash)
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "failed to create challenge")

	// check challenge created
	challenge, err := colosseumContract.GetChallenge(nil, outputIndex)
	require.NoError(t, err)
	require.NotNil(t, challenge, "challenge not found")

	// check bond amount doubled
	bond, err = valPoolContract.GetBond(nil, outputIndex)
	require.NoError(t, err)
	require.Equal(t, big.NewInt(2*dp.DeployConfig.ValidatorPoolMinBondAmount.ToInt().Int64()), bond.Amount)

interaction:
	for {
		status, err := colosseumContract.GetStatus(nil, outputIndex)
		require.NoError(t, err)

		switch status {
		case chal.StatusChallengerTurn:
			// call bisect by challenger
			txHash = challenger.ActBisect(t, outputIndex)
			includeL1Block(t, miner, challenger.address)
		case chal.StatusAsserterTurn:
			// call bisect by validator
			txHash = validator.ActBisect(t, outputIndex)
			includeL1Block(t, miner, validator.address)
		case chal.StatusAsserterTimeout:
			// not expected
		case chal.StatusReadyToProve:
			txHash = challenger.ActProveFault(t, outputIndex, false)
			includeL1Block(t, miner, challenger.address)
		case chal.StatusProven:
			// validate l2 output submitted by challenger
			outputBlockNum := outputOnL1.L2BlockNumber.Uint64()
			output := challenger.ActOutputAtBlockSafe(t, outputBlockNum)
			isValid := guardian.ActValidateL2Output(t, output.OutputRoot, outputBlockNum)
			require.True(t, isValid)
			txHash = guardian.ActConfirmTransaction(t, big.NewInt(0))
			includeL1Block(t, miner, guardian.address)

			// check challenger bond amount decreased
			cBal, err := valPoolContract.BalanceOf(nil, challenger.address)
			require.NoError(t, err)
			require.Equal(t, new(big.Int).SetUint64(challengerInitialAmount-1), cBal)
		default:
			break interaction
		}

		// Check whether the submission was successful
		receipt, err = miner.EthClient().TransactionReceipt(t.Ctx(), txHash)
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "failed to progress interactive proof")
	}

	// Check the status of challenge is StatusApproved(7)
	status, err := colosseumContract.GetStatus(nil, outputIndex)
	require.NoError(t, err)
	require.Equal(t, chal.StatusApproved, status)
}

func TestChallengerChallengerBisectTimeout(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	dp.DeployConfig.FinalizationPeriodSeconds = 60 * 60 * 24
	dp.DeployConfig.ColosseumDummyHash = common.HexToHash(e2e.DummyHashSepolia)
	dp.DeployConfig.ColosseumProvingTimeout = 1
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	l := testlog.Logger(t, log.LvlDebug)
	miner, propEngine, proposer := setupProposerTest(t, sd, l)
	var validatorInitialAmount uint64 = 1_000
	var challengerInitialAmount uint64 = 1_000

	rollupPropCl := proposer.RollupClient()
	batcher := NewL2Batcher(l, sd.RollupCfg, &BatcherCfg{
		MinL1TxSize: 0,
		MaxL1TxSize: 128_000,
		BatcherKey:  dp.Secrets.Batcher,
	}, rollupPropCl, miner.EthClient(), propEngine.EthClient())

	// setup mockup rpc for returning invalid output
	validatorRPC := e2eutils.NewMaliciousL2RPC(proposer.RPCClient())
	validatorRollupClient := sources.NewRollupClient(validatorRPC)
	validator := NewL2Validator(t, l, &ValidatorCfg{
		OutputOracleAddr:  sd.DeploymentsL1.L2OutputOracleProxy,
		ValidatorPoolAddr: sd.DeploymentsL1.ValidatorPoolProxy,
		ColosseumAddr:     sd.DeploymentsL1.ColosseumProxy,
		ValidatorKey:      dp.Secrets.TrustedValidator,
		AllowNonFinalized: false,
	}, miner.EthClient(), validatorRollupClient)

	challengerRPC := e2eutils.NewHonestL2RPC(proposer.RPCClient())
	challengerRollupClient := sources.NewRollupClient(challengerRPC)
	challenger := NewL2Validator(t, l, &ValidatorCfg{
		OutputOracleAddr:  sd.DeploymentsL1.L2OutputOracleProxy,
		ValidatorPoolAddr: sd.DeploymentsL1.ValidatorPoolProxy,
		ColosseumAddr:     sd.DeploymentsL1.ColosseumProxy,
		ValidatorKey:      dp.Secrets.Challenger,
		AllowNonFinalized: false,
	}, miner.EthClient(), challengerRollupClient)

	// bind contracts
	outputOracleContract, err := bindings.NewL2OutputOracle(sd.DeploymentsL1.L2OutputOracleProxy, miner.EthClient())
	require.NoError(t, err)

	colosseumContract, err := bindings.NewColosseum(sd.DeploymentsL1.ColosseumProxy, miner.EthClient())
	require.NoError(t, err)

	valPoolContract, err := bindings.NewValidatorPoolCaller(sd.DeploymentsL1.ValidatorPoolProxy, miner.EthClient())
	require.NoError(t, err)

	// NOTE(chokobole): After the Blue hard fork, it is necessary to wait for one finalized
	// (or safe if AllowNonFinalized config is set) block to pass after each submission interval
	// before submitting the output root.
	// For example, if the submission interval is set to 1800 blocks, before the Blue hard fork,
	// the output root could be submitted at 1800 finalized blocks. However, after the update,
	// the output root can only be submitted at 1801 finalized blocks.
	// In fact, the following code is designed to create one or more finalized L2 blocks
	// in order to pass the test after the Blue hard fork.
	// If Proto Dank Sharding is introduced, the below code fix may no longer be necessary.
	for i := 0; i < 3; i++ {
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
	}

	// deposit bond for validator
	validator.ActDeposit(t, validatorInitialAmount)
	includeL1Block(t, miner, validator.address)

	// check validator balance increased
	bal, err := valPoolContract.BalanceOf(nil, validator.address)
	require.NoError(t, err)
	require.Equal(t, new(big.Int).SetUint64(validatorInitialAmount), bal)

	require.Equal(t, proposer.SyncStatus().UnsafeL2, proposer.SyncStatus().FinalizedL2)

	validatorRPC.SetTargetBlockNumber(testdata.TargetBlockNumber)
	challengerRPC.SetTargetBlockNumber(testdata.TargetBlockNumber)
	// create l2 output submission transactions until there is nothing left to submit
	for validator.CanSubmit(t) {
		// and submit it to L1
		validator.ActSubmitL2Output(t)
		// include output on L1
		includeL1Block(t, miner, validator.address)
		// Check submission was successful
		receipt, err := miner.EthClient().TransactionReceipt(t.Ctx(), validator.LastSubmitL2OutputTx())
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "submission failed")
	}

	// check that L1 stored the expected output root
	// NOTE(chokobole): Comment these 2 lines because of the reason above.
	// If Proto Dank Sharding is introduced, the below code fix may be restored.
	// block := proposer.SyncStatus().FinalizedL2
	// outputOnL1, err := outputOracleContract.GetL2OutputAfter(nil, new(big.Int).SetUint64(block.Number))
	targetBlockNum := big.NewInt(int64(testdata.TargetBlockNumber))
	outputIndex, err := outputOracleContract.GetL2OutputIndexAfter(nil, targetBlockNum)
	require.NoError(t, err)
	outputOnL1, err := outputOracleContract.GetL2OutputAfter(nil, targetBlockNum)
	require.NoError(t, err)
	block, err := propEngine.EthClient().BlockByNumber(t.Ctx(), targetBlockNum)
	require.NoError(t, err)
	require.Less(t, block.Time(), outputOnL1.Timestamp.Uint64(), "output is registered with L1 timestamp of L2 tx output submission, past L2 block")
	outputComputed, err := proposer.RollupClient().OutputAtBlock(t.Ctx(), targetBlockNum.Uint64())
	require.NoError(t, err)
	require.NotEqual(t, eth.Bytes32(outputOnL1.OutputRoot), outputComputed.OutputRoot, "output roots must different")

	// deposit bond for challenger
	challenger.ActDeposit(t, challengerInitialAmount)
	includeL1Block(t, miner, challenger.address)

	// check bond amount before create challenge
	bond, err := valPoolContract.GetBond(nil, outputIndex)
	require.NoError(t, err)
	require.Equal(t, big.NewInt(dp.DeployConfig.ValidatorPoolMinBondAmount.ToInt().Int64()), bond.Amount)

	// submit create challenge tx
	txHash := challenger.ActCreateChallenge(t, outputIndex)

	// include tx on L1
	includeL1Block(t, miner, challenger.address)

	// Check whether the submission was successful
	receipt, err := miner.EthClient().TransactionReceipt(t.Ctx(), txHash)
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "failed to create challenge")

	challenge, err := colosseumContract.GetChallenge(nil, outputIndex)
	require.NoError(t, err)
	require.NotNil(t, challenge, "challenge not found")

	// check bond amount doubled
	bond, err = valPoolContract.GetBond(nil, outputIndex)
	require.NoError(t, err)
	require.Equal(t, big.NewInt(2*dp.DeployConfig.ValidatorPoolMinBondAmount.ToInt().Int64()), bond.Amount)

interaction:
	for {
		status, err := colosseumContract.GetStatus(nil, outputIndex)
		require.NoError(t, err)

		switch status {
		case chal.StatusChallengerTurn:
			// do nothing to trigger challenger timeout
			miner.ActEmptyBlock(t)
		case chal.StatusAsserterTurn:
			// call bisect by validator
			txHash = validator.ActBisect(t, outputIndex)
			includeL1Block(t, miner, validator.address)
		case chal.StatusAsserterTimeout, chal.StatusReadyToProve:
			// do nothing to trigger challenger timeout
			miner.ActEmptyBlock(t)
		case chal.StatusChallengerTimeout:
			// check challenger bond amount decreased
			cBal, err := valPoolContract.BalanceOf(nil, challenger.address)
			require.NoError(t, err)
			require.Equal(t, new(big.Int).SetUint64(challengerInitialAmount-1), cBal)
			break interaction
		case chal.StatusProven:
			// not expected
		default:
			break interaction
		}

		// Check whether the submission was successful
		receipt, err = miner.EthClient().TransactionReceipt(t.Ctx(), txHash)
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "failed to progress interactive proof")
	}

	// Check the status of challenge is StatusApproved(7)
	status, err := colosseumContract.GetStatus(nil, outputIndex)
	require.NoError(t, err)
	require.Equal(t, chal.StatusChallengerTimeout, status)
}

func TestChallengerChallengerProvingTimeout(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	dp.DeployConfig.FinalizationPeriodSeconds = 60 * 60 * 24
	dp.DeployConfig.ColosseumDummyHash = common.HexToHash(e2e.DummyHashSepolia)
	dp.DeployConfig.ColosseumProvingTimeout = 1
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	l := testlog.Logger(t, log.LvlDebug)
	miner, propEngine, proposer := setupProposerTest(t, sd, l)
	var validatorInitialAmount uint64 = 1_000
	var challengerInitialAmount uint64 = 1_000

	rollupPropCl := proposer.RollupClient()
	batcher := NewL2Batcher(l, sd.RollupCfg, &BatcherCfg{
		MinL1TxSize: 0,
		MaxL1TxSize: 128_000,
		BatcherKey:  dp.Secrets.Batcher,
	}, rollupPropCl, miner.EthClient(), propEngine.EthClient())

	// setup mockup rpc for returning invalid output
	validatorRPC := e2eutils.NewMaliciousL2RPC(proposer.RPCClient())
	validatorRollupClient := sources.NewRollupClient(validatorRPC)
	validator := NewL2Validator(t, l, &ValidatorCfg{
		OutputOracleAddr:  sd.DeploymentsL1.L2OutputOracleProxy,
		ValidatorPoolAddr: sd.DeploymentsL1.ValidatorPoolProxy,
		ColosseumAddr:     sd.DeploymentsL1.ColosseumProxy,
		ValidatorKey:      dp.Secrets.TrustedValidator,
		AllowNonFinalized: false,
	}, miner.EthClient(), validatorRollupClient)

	challengerRPC := e2eutils.NewHonestL2RPC(proposer.RPCClient())
	challengerRollupClient := sources.NewRollupClient(challengerRPC)
	challenger := NewL2Validator(t, l, &ValidatorCfg{
		OutputOracleAddr:  sd.DeploymentsL1.L2OutputOracleProxy,
		ValidatorPoolAddr: sd.DeploymentsL1.ValidatorPoolProxy,
		ColosseumAddr:     sd.DeploymentsL1.ColosseumProxy,
		ValidatorKey:      dp.Secrets.Challenger,
		AllowNonFinalized: false,
	}, miner.EthClient(), challengerRollupClient)

	// bind contracts
	outputOracleContract, err := bindings.NewL2OutputOracle(sd.DeploymentsL1.L2OutputOracleProxy, miner.EthClient())
	require.NoError(t, err)

	colosseumContract, err := bindings.NewColosseum(sd.DeploymentsL1.ColosseumProxy, miner.EthClient())
	require.NoError(t, err)

	valPoolContract, err := bindings.NewValidatorPoolCaller(sd.DeploymentsL1.ValidatorPoolProxy, miner.EthClient())
	require.NoError(t, err)

	// NOTE(chokobole): After the Blue hard fork, it is necessary to wait for one finalized
	// (or safe if AllowNonFinalized config is set) block to pass after each submission interval
	// before submitting the output root.
	// For example, if the submission interval is set to 1800 blocks, before the Blue hard fork,
	// the output root could be submitted at 1800 finalized blocks. However, after the update,
	// the output root can only be submitted at 1801 finalized blocks.
	// In fact, the following code is designed to create one or more finalized L2 blocks
	// in order to pass the test after the Blue hard fork.
	// If Proto Dank Sharding is introduced, the below code fix may no longer be necessary.
	for i := 0; i < 3; i++ {
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
	}

	// deposit bond for validator
	validator.ActDeposit(t, validatorInitialAmount)
	includeL1Block(t, miner, validator.address)

	// check validator balance increased
	bal, err := valPoolContract.BalanceOf(nil, validator.address)
	require.NoError(t, err)
	require.Equal(t, new(big.Int).SetUint64(validatorInitialAmount), bal)

	require.Equal(t, proposer.SyncStatus().UnsafeL2, proposer.SyncStatus().FinalizedL2)

	validatorRPC.SetTargetBlockNumber(testdata.TargetBlockNumber)
	challengerRPC.SetTargetBlockNumber(testdata.TargetBlockNumber)
	// create l2 output submission transactions until there is nothing left to submit
	for validator.CanSubmit(t) {
		// and submit it to L1
		validator.ActSubmitL2Output(t)
		// include output on L1
		includeL1Block(t, miner, validator.address)
		// Check submission was successful
		receipt, err := miner.EthClient().TransactionReceipt(t.Ctx(), validator.LastSubmitL2OutputTx())
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "submission failed")
	}

	// check that L1 stored the expected output root
	// NOTE(chokobole): Comment these 2 lines because of the reason above.
	// If Proto Dank Sharding is introduced, the below code fix may be restored.
	// block := proposer.SyncStatus().FinalizedL2
	// outputOnL1, err := outputOracleContract.GetL2OutputAfter(nil, new(big.Int).SetUint64(block.Number))
	targetBlockNum := big.NewInt(int64(testdata.TargetBlockNumber))
	outputIndex, err := outputOracleContract.GetL2OutputIndexAfter(nil, targetBlockNum)
	require.NoError(t, err)
	outputOnL1, err := outputOracleContract.GetL2OutputAfter(nil, targetBlockNum)
	require.NoError(t, err)
	block, err := propEngine.EthClient().BlockByNumber(t.Ctx(), targetBlockNum)
	require.NoError(t, err)
	require.Less(t, block.Time(), outputOnL1.Timestamp.Uint64(), "output is registered with L1 timestamp of L2 tx output submission, past L2 block")
	outputComputed, err := proposer.RollupClient().OutputAtBlock(t.Ctx(), targetBlockNum.Uint64())
	require.NoError(t, err)
	require.NotEqual(t, eth.Bytes32(outputOnL1.OutputRoot), outputComputed.OutputRoot, "output roots must different")

	// deposit bond for challenger
	challenger.ActDeposit(t, challengerInitialAmount)
	includeL1Block(t, miner, challenger.address)

	// check bond amount before create challenge
	bond, err := valPoolContract.GetBond(nil, outputIndex)
	require.NoError(t, err)
	require.Equal(t, big.NewInt(dp.DeployConfig.ValidatorPoolMinBondAmount.ToInt().Int64()), bond.Amount)

	// submit create challenge tx
	txHash := challenger.ActCreateChallenge(t, outputIndex)

	// include tx on L1
	includeL1Block(t, miner, challenger.address)

	// Check whether the submission was successful
	receipt, err := miner.EthClient().TransactionReceipt(t.Ctx(), txHash)
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "failed to create challenge")

	challenge, err := colosseumContract.GetChallenge(nil, outputIndex)
	require.NoError(t, err)
	require.NotNil(t, challenge, "challenge not found")

	// check bond amount doubled
	bond, err = valPoolContract.GetBond(nil, outputIndex)
	require.NoError(t, err)
	require.Equal(t, big.NewInt(2*dp.DeployConfig.ValidatorPoolMinBondAmount.ToInt().Int64()), bond.Amount)

interaction:
	for {
		status, err := colosseumContract.GetStatus(nil, outputIndex)
		require.NoError(t, err)

		switch status {
		case chal.StatusChallengerTurn:
			// call bisect by challenger
			txHash = challenger.ActBisect(t, outputIndex)
			includeL1Block(t, miner, challenger.address)
		case chal.StatusAsserterTurn:
			// call bisect by validator
			txHash = validator.ActBisect(t, outputIndex)
			includeL1Block(t, miner, validator.address)
		case chal.StatusAsserterTimeout, chal.StatusReadyToProve:
			// do nothing to trigger challenger timeout
			miner.ActEmptyBlock(t)
		case chal.StatusChallengerTimeout:
			// check challenger bond amount decreased
			cBal, err := valPoolContract.BalanceOf(nil, challenger.address)
			require.NoError(t, err)
			require.Equal(t, new(big.Int).SetUint64(challengerInitialAmount-1), cBal)
			break interaction
		case chal.StatusProven:
			// not expected
		default:
			break interaction
		}

		// Check whether the submission was successful
		receipt, err = miner.EthClient().TransactionReceipt(t.Ctx(), txHash)
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "failed to progress interactive proof")
	}

	// Check the status of challenge is StatusApproved(7)
	status, err := colosseumContract.GetStatus(nil, outputIndex)
	require.NoError(t, err)
	require.Equal(t, chal.StatusChallengerTimeout, status)
}

func TestChallengerInvalidProofFail(gt *testing.T) {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	dp.DeployConfig.FinalizationPeriodSeconds = 60 * 60 * 24
	dp.DeployConfig.ColosseumDummyHash = common.HexToHash(e2e.DummyHashSepolia)
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	l := testlog.Logger(t, log.LvlDebug)
	miner, propEngine, proposer := setupProposerTest(t, sd, l)

	rollupPropCl := proposer.RollupClient()
	batcher := NewL2Batcher(l, sd.RollupCfg, &BatcherCfg{
		MinL1TxSize: 0,
		MaxL1TxSize: 128_000,
		BatcherKey:  dp.Secrets.Batcher,
	}, rollupPropCl, miner.EthClient(), propEngine.EthClient())

	validatorRPC := e2eutils.NewMaliciousL2RPC(proposer.RPCClient())
	validatorRollupClient := sources.NewRollupClient(validatorRPC)
	validator := NewL2Validator(t, l, &ValidatorCfg{
		OutputOracleAddr:  sd.DeploymentsL1.L2OutputOracleProxy,
		ValidatorPoolAddr: sd.DeploymentsL1.ValidatorPoolProxy,
		ColosseumAddr:     sd.DeploymentsL1.ColosseumProxy,
		ValidatorKey:      dp.Secrets.TrustedValidator,
		AllowNonFinalized: false,
	}, miner.EthClient(), validatorRollupClient)

	challengerRPC := e2eutils.NewHonestL2RPC(proposer.RPCClient())
	challengerRollupClient := sources.NewRollupClient(challengerRPC)
	challenger := NewL2Validator(t, l, &ValidatorCfg{
		OutputOracleAddr:  sd.DeploymentsL1.L2OutputOracleProxy,
		ValidatorPoolAddr: sd.DeploymentsL1.ValidatorPoolProxy,
		ColosseumAddr:     sd.DeploymentsL1.ColosseumProxy,
		ValidatorKey:      dp.Secrets.Challenger,
		AllowNonFinalized: false,
	}, miner.EthClient(), challengerRollupClient)

	guardianRPC := e2eutils.NewMaliciousL2RPC(proposer.RPCClient())
	guardianRollupClient := sources.NewRollupClient(guardianRPC)
	guardian := NewL2Validator(t, l, &ValidatorCfg{
		OutputOracleAddr:    sd.DeploymentsL1.L2OutputOracleProxy,
		SecurityCouncilAddr: sd.DeploymentsL1.SecurityCouncilProxy,
		ValidatorKey:        dp.Secrets.Challenger,
		AllowNonFinalized:   false,
	}, miner.EthClient(), guardianRollupClient)

	// NOTE(chokobole): It is necessary to wait for one finalized (or safe if AllowNonFinalized
	// config is set) block to pass after each submission interval before submitting the output
	// root. For example, if the submission interval is set to 1800 blocks, the output root can
	// only be submitted at 1801 finalized blocks. In fact, the following code is designed to
	// create one or more finalized L2 blocks in order to pass the test. If Proto Dank Sharding
	// is introduced, the below code fix may no longer be necessary.
	for i := 0; i < 3; i++ {
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
	}

	// deposit bond for validator
	validator.ActDeposit(t, 1_000)
	includeL1Block(t, miner, validator.address)

	require.Equal(t, proposer.SyncStatus().UnsafeL2, proposer.SyncStatus().FinalizedL2)

	outputOracleContract, err := bindings.NewL2OutputOracle(sd.DeploymentsL1.L2OutputOracleProxy, miner.EthClient())
	require.NoError(t, err)

	validatorRPC.SetTargetBlockNumber(testdata.TargetBlockNumber)
	challengerRPC.SetTargetBlockNumber(testdata.TargetBlockNumber)
	guardianRPC.SetTargetBlockNumber(testdata.TargetBlockNumber)
	// create l2 output submission transactions until there is nothing left to submit
	for validator.CanSubmit(t) {
		// and submit it to L1
		validator.ActSubmitL2Output(t)
		// include output on L1
		includeL1Block(t, miner, validator.address)
		// Check submission was successful
		receipt, err := miner.EthClient().TransactionReceipt(t.Ctx(), validator.LastSubmitL2OutputTx())
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "submission failed")
	}

	// check that L1 stored the expected output root
	// NOTE(chokobole): Comment these 2 lines because of the reason above about the Blue hard fork.
	// If Proto Dank Sharding is introduced, the below code fix may be restored.
	// block := proposer.SyncStatus().FinalizedL2
	// outputOnL1, err := outputOracleContract.GetL2OutputAfter(nil, new(big.Int).SetUint64(block.Number))
	blockNum := big.NewInt(int64(testdata.TargetBlockNumber))
	outputIndex, err := outputOracleContract.GetL2OutputIndexAfter(nil, blockNum)
	require.NoError(t, err)
	outputOnL1, err := outputOracleContract.GetL2OutputAfter(nil, blockNum)
	require.NoError(t, err)
	block, err := propEngine.EthClient().BlockByNumber(t.Ctx(), blockNum)
	require.NoError(t, err)
	require.Less(t, block.Time(), outputOnL1.Timestamp.Uint64(), "output is registered with L1 timestamp of L2 tx output submission, past L2 block")
	outputComputed, err := challengerRollupClient.OutputAtBlock(t.Ctx(), blockNum.Uint64())
	require.NoError(t, err)
	require.NotEqual(t, eth.Bytes32(outputOnL1.OutputRoot), outputComputed.OutputRoot, "output roots must different")

	// deposit bond for challenger
	challenger.ActDeposit(t, 1_000)
	includeL1Block(t, miner, challenger.address)

	// submit create challenge tx
	txHash := challenger.ActCreateChallenge(t, outputIndex)

	// include tx on L1
	includeL1Block(t, miner, challenger.address)

	// Check whether the submission was successful
	receipt, err := miner.EthClient().TransactionReceipt(t.Ctx(), txHash)
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "failed to create challenge")

	// check challenge created
	colosseumContract, err := bindings.NewColosseum(sd.DeploymentsL1.ColosseumProxy, miner.EthClient())
	require.NoError(t, err)
	challenge, err := colosseumContract.GetChallenge(nil, outputIndex)
	require.NoError(t, err)
	require.NotNil(t, challenge, "challenge not found")

interaction:
	for {
		status, err := colosseumContract.GetStatus(nil, outputIndex)
		require.NoError(t, err)

		switch status {
		case chal.StatusChallengerTurn:
			// call bisect by challenger
			txHash = challenger.ActBisect(t, outputIndex)
			includeL1Block(t, miner, challenger.address)
		case chal.StatusAsserterTurn:
			// call bisect by validator
			txHash = validator.ActBisect(t, outputIndex)
			includeL1Block(t, miner, validator.address)
		case chal.StatusAsserterTimeout, chal.StatusReadyToProve:
			txHash = challenger.ActProveFault(t, outputIndex, false)
			includeL1Block(t, miner, challenger.address)
		case chal.StatusProven:
			// validate l2 output submitted by challenger
			outputBlockNum := outputOnL1.L2BlockNumber.Uint64()
			output := challenger.ActOutputAtBlockSafe(t, outputBlockNum)
			isValid := guardian.ActValidateL2Output(t, output.OutputRoot, outputBlockNum)
			require.False(t, isValid)
			break interaction
		default:
			break interaction
		}

		// Check whether the submission was successful
		receipt, err = miner.EthClient().TransactionReceipt(t.Ctx(), txHash)
		require.NoError(t, err)
		require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status, "failed to progress interactive proof")
	}

	// Check the status of challenge is StatusProven(7)
	status, err := colosseumContract.GetStatus(nil, outputIndex)
	require.NoError(t, err)
	require.Equal(t, chal.StatusProven, status)
}
