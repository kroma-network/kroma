package actions

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/optimism/op-node/sources"
	"github.com/ethereum-optimism/optimism/op-node/testlog"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	e2e "github.com/ethereum-optimism/optimism/op-e2e"
)

const defaultDepositAmount = 1_000

type Runtime struct {
	t                        StatefulTesting
	l                        log.Logger
	sd                       *e2eutils.SetupData
	dp                       *e2eutils.DeployParams
	miner                    *L1Miner
	seqEngine                *L2Engine
	sequencer                *L2Sequencer
	batcher                  *L2Batcher
	validator                *L2Validator
	challenger1              *L2Validator
	challenger2              *L2Validator
	guardian                 *L2Validator
	outputOracleContract     *bindings.L2OutputOracle
	colosseumContract        *bindings.Colosseum
	securityCouncilContract  *bindings.SecurityCouncil
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
	dp.DeployConfig.ColosseumCreationPeriodSeconds = 60 * 60 * 20
	dp.DeployConfig.ColosseumDummyHash = common.HexToHash(e2e.DummyHashDev)
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	l := testlog.Logger(t, log.LvlDebug)
	rt := Runtime{
		t:  t,
		dp: dp,
		sd: sd,
		l:  l,
	}
	rt.miner, rt.seqEngine, rt.sequencer = setupSequencerTest(rt.t, rt.sd, rt.l)
	rt.setupBatcher()

	return rt
}

func (rt *Runtime) setupBatcher() {
	rollupSeqCl := rt.sequencer.RollupClient()
	batcher := NewL2Batcher(rt.l, rt.sd.RollupCfg, &BatcherCfg{
		MinL1TxSize: 0,
		MaxL1TxSize: 128_000,
		BatcherKey:  rt.dp.Secrets.Batcher,
	}, rollupSeqCl, rt.miner.EthClient(), rt.seqEngine.EthClient())
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

func (rt *Runtime) setupHonestChallenger1() {
	rt.challenger1 = rt.honestValidator(rt.dp.Secrets.Challenger1)
}

func (rt *Runtime) setupHonestChallenger2() {
	rt.challenger2 = rt.honestValidator(rt.dp.Secrets.Challenger2)
}

func (rt *Runtime) setupMaliciousChallenger1() {
	rt.challenger1 = rt.maliciousValidator(rt.dp.Secrets.Challenger1)
}

func (rt *Runtime) setupMaliciousChallenger2() {
	rt.challenger2 = rt.maliciousValidator(rt.dp.Secrets.Challenger2)
}

func (rt *Runtime) setupHonestGuardian() {
	rt.guardian = rt.honestValidator(rt.dp.Secrets.Challenger1)
}

func (rt *Runtime) setupMaliciousGuardian() {
	rt.guardian = rt.maliciousValidator(rt.dp.Secrets.Challenger1)
}

func (rt *Runtime) honestValidator(pk *ecdsa.PrivateKey) *L2Validator {
	// setup mockup rpc for returning valid output
	validatorRPC := e2eutils.NewHonestL2RPC(rt.sequencer.RPCClient())
	validatorRollupClient := sources.NewRollupClient(validatorRPC)
	validator := NewL2Validator(rt.t, rt.l, &ValidatorCfg{
		OutputOracleAddr:    rt.sd.DeploymentsL1.L2OutputOracleProxy,
		ValidatorPoolAddr:   rt.sd.DeploymentsL1.ValidatorPoolProxy,
		ColosseumAddr:       rt.sd.DeploymentsL1.ColosseumProxy,
		SecurityCouncilAddr: rt.sd.DeploymentsL1.SecurityCouncilProxy,
		ValidatorKey:        pk,
		AllowNonFinalized:   false,
	}, rt.miner.EthClient(), rt.seqEngine.EthClient(), validatorRollupClient)
	validatorRPC.SetTargetBlockNumber(rt.targetInvalidBlockNumber)
	return validator
}

func (rt *Runtime) maliciousValidator(pk *ecdsa.PrivateKey) *L2Validator {
	// setup mockup rpc for returning invalid output
	validatorRPC := e2eutils.NewMaliciousL2RPC(rt.sequencer.RPCClient())
	validatorRollupClient := sources.NewRollupClient(validatorRPC)
	validator := NewL2Validator(rt.t, rt.l, &ValidatorCfg{
		OutputOracleAddr:    rt.sd.DeploymentsL1.L2OutputOracleProxy,
		ValidatorPoolAddr:   rt.sd.DeploymentsL1.ValidatorPoolProxy,
		ColosseumAddr:       rt.sd.DeploymentsL1.ColosseumProxy,
		SecurityCouncilAddr: rt.sd.DeploymentsL1.SecurityCouncilProxy,
		ValidatorKey:        pk,
		AllowNonFinalized:   false,
	}, rt.miner.EthClient(), rt.seqEngine.EthClient(), validatorRollupClient)
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

	rt.securityCouncilContract, err = bindings.NewSecurityCouncil(rt.sd.DeploymentsL1.SecurityCouncilProxy, rt.miner.EthClient())
	require.NoError(rt.t, err)

	rt.valPoolContract, err = bindings.NewValidatorPoolCaller(rt.sd.DeploymentsL1.ValidatorPoolProxy, rt.miner.EthClient())
	require.NoError(rt.t, err)
}

// setupOutputSubmitted sets output submission by validator
func (rt *Runtime) setupOutputSubmitted() {
	// NOTE(chokobole): It is necessary to wait for one finalized (or safe if AllowNonFinalized
	// config is set) block to pass after each submission interval before submitting the output
	// root. For example, if the submission interval is set to 1800 blocks, the output root can
	// only be submitted at 1801 finalized blocks. In fact, the following code is designed to
	// create one or more finalized L2 blocks in order to pass the test. If Proto Dank Sharding
	// is introduced, the below code fix may no longer be necessary.
	for i := 0; i < 5; i++ {
		// L1 block
		rt.miner.ActEmptyBlock(rt.t)
		// L2 block
		rt.sequencer.ActL1HeadSignal(rt.t)
		rt.sequencer.ActL2PipelineFull(rt.t)
		rt.sequencer.ActBuildToL1Head(rt.t)
		// submit and include in L1
		rt.batcher.ActSubmitAll(rt.t)
		rt.miner.includeL1Block(rt.t, rt.dp.Addresses.Batcher)
		// finalize the first and second L1 blocks, including the batch
		rt.miner.ActL1SafeNext(rt.t)
		rt.miner.ActL1SafeNext(rt.t)
		rt.miner.ActL1FinalizeNext(rt.t)
		rt.miner.ActL1FinalizeNext(rt.t)
		// derive and see the L2 chain fully finalize
		rt.sequencer.ActL2PipelineFull(rt.t)
		rt.sequencer.ActL1SafeSignal(rt.t)
		rt.sequencer.ActL1FinalizedSignal(rt.t)
	}

	// deposit bond for validator
	rt.validator.ActDeposit(rt.t, defaultDepositAmount)
	rt.miner.includeL1Block(rt.t, rt.validator.address)

	// check validator balance increased
	bal, err := rt.valPoolContract.BalanceOf(nil, rt.validator.address)
	require.NoError(rt.t, err)
	require.Equal(rt.t, new(big.Int).SetUint64(defaultDepositAmount), bal)

	require.Equal(rt.t, rt.sequencer.SyncStatus().UnsafeL2, rt.sequencer.SyncStatus().FinalizedL2)

	// create l2 output submission transactions until there is nothing left to submit
	for {
		waitTime := rt.validator.CalculateWaitTime(rt.t)
		if waitTime > 0 {
			break
		}
		// and submit it to L1
		rt.validator.ActSubmitL2Output(rt.t)
		// include output on L1
		rt.miner.includeL1Block(rt.t, rt.validator.address)
		// Check submission was successful
		receipt, err := rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), rt.validator.LastSubmitL2OutputTx())
		require.NoError(rt.t, err)
		require.Equal(rt.t, types.ReceiptStatusSuccessful, receipt.Status, "submission failed")
	}
}

// setupChallenge sets challenge by challenger
func (rt *Runtime) setupChallenge(challenger *L2Validator) {
	// check that the output root that L1 stores is different from challenger's output root
	// NOTE(chokobole): Comment these 2 lines because of the reason above.
	// If Proto Dank Sharding is introduced, the below code fix may be restored.
	// block := sequencer.SyncStatus().FinalizedL2
	// outputOnL1, err := outputOracleContract.GetL2OutputAfter(nil, new(big.Int).SetUint64(block.Number))
	targetBlockNum := big.NewInt(int64(rt.targetInvalidBlockNumber))
	var err error
	rt.outputIndex, err = rt.outputOracleContract.GetL2OutputIndexAfter(nil, targetBlockNum)
	require.NoError(rt.t, err)
	rt.outputOnL1, err = rt.outputOracleContract.GetL2OutputAfter(nil, targetBlockNum)
	require.NoError(rt.t, err)
	block, err := rt.seqEngine.EthClient().BlockByNumber(rt.t.Ctx(), targetBlockNum)
	require.NoError(rt.t, err)
	require.Less(rt.t, block.Time(), rt.outputOnL1.Timestamp.Uint64(), "output is registered with L1 timestamp of L2 tx output submission, past L2 block")
	outputComputed := challenger.fetchOutput(rt.t, rt.outputOnL1.L2BlockNumber)
	require.NotEqual(rt.t, eth.Bytes32(rt.outputOnL1.OutputRoot), outputComputed.OutputRoot, "output roots must different")

	// deposit bond for challenger
	challenger.ActDeposit(rt.t, defaultDepositAmount)
	rt.miner.includeL1Block(rt.t, challenger.address)

	// check bond amount before create challenge
	bond, err := rt.valPoolContract.GetBond(nil, rt.outputIndex)
	require.NoError(rt.t, err)
	require.Equal(rt.t, rt.dp.DeployConfig.ValidatorPoolRequiredBondAmount.ToInt(), bond.Amount)

	// submit create challenge tx
	rt.txHash = challenger.ActCreateChallenge(rt.t, rt.outputIndex)

	// include tx on L1
	rt.miner.includeL1Block(rt.t, challenger.address)

	// Check whether the submission was successful
	rt.receipt, err = rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), rt.txHash)
	require.NoError(rt.t, err)
	require.Equal(rt.t, types.ReceiptStatusSuccessful, rt.receipt.Status, "failed to create challenge")

	// check challenge created
	challenge, err := rt.colosseumContract.GetChallenge(nil, rt.outputIndex, challenger.address)
	require.NoError(rt.t, err)
	require.NotNil(rt.t, challenge, "challenge not found")

	// check pending bond amount after create challenge
	pendingBond, err := rt.valPoolContract.GetPendingBond(nil, rt.outputIndex, challenger.address)
	require.NoError(rt.t, err)
	require.Equal(rt.t, pendingBond, rt.dp.DeployConfig.ValidatorPoolRequiredBondAmount.ToInt())

	// check challenger balance decreased
	cBal, err := rt.valPoolContract.BalanceOf(nil, challenger.address)
	require.NoError(rt.t, err)
	require.Equal(rt.t, new(big.Int).Sub(new(big.Int).SetInt64(defaultDepositAmount), rt.dp.DeployConfig.ValidatorPoolRequiredBondAmount.ToInt()), cBal)
}

// IsCreationEnded checks if the creation period of rt.outputIndex output is ended
func (rt *Runtime) IsCreationEnded() bool {
	output, err := rt.outputOracleContract.GetL2Output(nil, rt.outputIndex)
	require.NoError(rt.t, err)

	ended := output.Timestamp.Uint64() + rt.dp.DeployConfig.ColosseumCreationPeriodSeconds
	isEnded := rt.miner.l1Chain.CurrentBlock().Time > ended
	return isEnded
}

func (rt *Runtime) SetCreationPeriod(period uint64) {
	rt.dp.DeployConfig.ColosseumCreationPeriodSeconds = period
}
