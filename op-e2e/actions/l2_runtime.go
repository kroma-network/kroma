package actions

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/kroma-bindings/bindings"
	val "github.com/kroma-network/kroma/kroma-validator"
	valhelper "github.com/kroma-network/kroma/op-e2e/e2eutils/validator"
	"github.com/kroma-network/kroma/op-e2e/testdata"
)

const defaultDepositAmount = 1_000

var defaultValPoolTerminationIndex = common.Big2

// These definitions are moved from l1_replica_test.go file.
var defaultRollupTestParams = &e2eutils.TestParams{
	MaxSequencerDrift:   40,
	SequencerWindowSize: 120,
	ChannelTimeout:      120,
	L1BlockTime:         15,
}

var defaultAlloc = &e2eutils.AllocParams{PrefundTestUsers: true}

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
	colosseumContract        *bindings.MockColosseum
	securityCouncilContract  *bindings.SecurityCouncil
	valPoolContract          *bindings.ValidatorPoolCaller
	valMgrContract           *bindings.ValidatorManagerCaller
	assetMgrContract         *bindings.AssetManagerCaller
	assetTokenContract       *bindings.GovernanceTokenCaller
	targetInvalidBlockNumber uint64
	outputIndex              *big.Int
	outputOnL1               bindings.TypesCheckpointOutput
	txHash                   common.Hash
	receipt                  *types.Receipt
	l1BlockDelta             uint64
}

type SetupSequencerTestFunc = func(t Testing, sd *e2eutils.SetupData, log log.Logger) (*L1Miner, *L2Engine, *L2Sequencer)

// defaultRuntime is currently only used for l2_challenger_test.
func defaultRuntime(gt *testing.T, setupSequencerTest SetupSequencerTestFunc, deltaTimeOffset *hexutil.Uint64) Runtime {
	t := NewDefaultTesting(gt)
	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	dp.DeployConfig.L2GenesisDeltaTimeOffset = deltaTimeOffset
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	l := testlog.Logger(t, log.LvlDebug)
	rt := Runtime{
		t:            t,
		dp:           dp,
		sd:           sd,
		l:            l,
		l1BlockDelta: 6,
	}

	rt.miner, rt.seqEngine, rt.sequencer = setupSequencerTest(rt.t, rt.sd, rt.l)
	rt.setupBatcher(dp)

	return rt
}

func (rt *Runtime) setupBatcher(dp *e2eutils.DeployParams) {
	rollupSeqCl := rt.sequencer.RollupClient()
	batcher := NewL2Batcher(rt.l, rt.sd.RollupCfg, DefaultBatcherCfg(dp),
		rollupSeqCl, rt.miner.EthClient(), rt.seqEngine.EthClient(), rt.seqEngine.EngineClient(rt.t, rt.sd.RollupCfg))
	rt.batcher = batcher
}

func (rt *Runtime) setTargetInvalidBlockNumber(targetInvalidBlockNumber uint64) {
	rt.targetInvalidBlockNumber = targetInvalidBlockNumber
}

func (rt *Runtime) setupHonestValidator(setInvalidBlockNumber bool) {
	rt.validator = rt.setupValidator(rt.dp.Secrets.TrustedValidator, setInvalidBlockNumber, false)
}

func (rt *Runtime) setupMaliciousValidator() {
	rt.validator = rt.setupValidator(rt.dp.Secrets.TrustedValidator, true, true)
}

func (rt *Runtime) setupHonestChallenger1() {
	rt.challenger1 = rt.setupValidator(rt.dp.Secrets.Challenger1, true, false)
}

func (rt *Runtime) setupHonestChallenger2() {
	rt.challenger2 = rt.setupValidator(rt.dp.Secrets.Challenger2, true, false)
}

func (rt *Runtime) setupMaliciousChallenger1() {
	rt.challenger1 = rt.setupValidator(rt.dp.Secrets.Challenger1, true, true)
}

func (rt *Runtime) setupMaliciousChallenger2() {
	rt.challenger2 = rt.setupValidator(rt.dp.Secrets.Challenger2, true, true)
}

func (rt *Runtime) setupHonestGuardian() {
	rt.guardian = rt.setupValidator(rt.dp.Secrets.Guardian, true, false)
}

func (rt *Runtime) setupMaliciousGuardian() {
	rt.guardian = rt.setupValidator(rt.dp.Secrets.Guardian, true, true)
}

func (rt *Runtime) setupValidator(pk *ecdsa.PrivateKey, setInvalidBlockNumber bool, isMalicious bool) *L2Validator {
	var validatorRollupClient *sources.RollupClient
	if isMalicious {
		// setup mockup rpc for returning invalid output
		validatorRPC, err := e2eutils.NewMaliciousL2RPC(rt.sequencer.RPCClient(), testdata.DefaultProofType)
		require.NoError(rt.t, err)
		validatorRPC.SetTargetBlockNumber(rt.targetInvalidBlockNumber)
		validatorRollupClient = sources.NewRollupClient(validatorRPC)
	} else {
		// setup mockup rpc for returning valid output
		validatorRPC, err := e2eutils.NewHonestL2RPC(rt.sequencer.RPCClient(), testdata.DefaultProofType)
		require.NoError(rt.t, err)
		if setInvalidBlockNumber {
			validatorRPC.SetTargetBlockNumber(rt.targetInvalidBlockNumber)
		}
		validatorRollupClient = sources.NewRollupClient(validatorRPC)
	}

	validator := NewL2Validator(rt.t, rt.l, &ValidatorCfg{
		OutputOracleAddr:     rt.sd.DeploymentsL1.L2OutputOracleProxy,
		ValidatorPoolAddr:    rt.sd.DeploymentsL1.ValidatorPoolProxy,
		ValidatorManagerAddr: rt.sd.DeploymentsL1.ValidatorManagerProxy,
		AssetManagerAddr:     rt.sd.DeploymentsL1.AssetManagerProxy,
		ColosseumAddr:        rt.sd.DeploymentsL1.ColosseumProxy,
		SecurityCouncilAddr:  rt.sd.DeploymentsL1.SecurityCouncilProxy,
		ValidatorKey:         pk,
		AllowNonFinalized:    false,
	}, rt.miner.EthClient(), rt.seqEngine.EthClient(), validatorRollupClient)
	return validator
}

func (rt *Runtime) bindContracts() {
	var err error
	// bind contracts
	rt.outputOracleContract, err = bindings.NewL2OutputOracle(rt.sd.DeploymentsL1.L2OutputOracleProxy, rt.miner.EthClient())
	require.NoError(rt.t, err)

	rt.colosseumContract, err = bindings.NewMockColosseum(rt.sd.DeploymentsL1.ColosseumProxy, rt.miner.EthClient())
	require.NoError(rt.t, err)

	rt.securityCouncilContract, err = bindings.NewSecurityCouncil(rt.sd.DeploymentsL1.SecurityCouncilProxy, rt.miner.EthClient())
	require.NoError(rt.t, err)

	rt.valPoolContract, err = bindings.NewValidatorPoolCaller(rt.sd.DeploymentsL1.ValidatorPoolProxy, rt.miner.EthClient())
	require.NoError(rt.t, err)

	rt.valMgrContract, err = bindings.NewValidatorManagerCaller(rt.sd.DeploymentsL1.ValidatorManagerProxy, rt.miner.EthClient())
	require.NoError(rt.t, err)

	rt.assetMgrContract, err = bindings.NewAssetManagerCaller(rt.sd.DeploymentsL1.AssetManagerProxy, rt.miner.EthClient())
	require.NoError(rt.t, err)

	assetTokenAddr, err := rt.assetMgrContract.ASSETTOKEN(nil)
	require.NoError(rt.t, err)
	rt.assetTokenContract, err = bindings.NewGovernanceTokenCaller(assetTokenAddr, rt.miner.EthClient())
	require.NoError(rt.t, err)
}

// assertRedeployValPoolToTerminate redeploys and upgrades ValidatorPool to change the termination index.
// It also asserts that the deploying and upgrade tx is successful.
func (rt *Runtime) assertRedeployValPoolToTerminate(newTerminationIndex *big.Int) {
	deployTx, upgradeTx, err := e2eutils.RedeployValPoolToTerminate(
		newTerminationIndex,
		rt.miner.EthClient(),
		rt.dp.Secrets,
		rt.sd.RollupCfg.L1ChainID,
		rt.sd.DeploymentsL1,
		rt.dp.DeployConfig,
	)
	require.NoError(rt.t, err)

	// Check deploy tx submission was successful
	rt.includeL1BlockByTx(deployTx.Hash())
	receipt, err := rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), deployTx.Hash())
	require.NoError(rt.t, err)
	require.Equal(rt.t, types.ReceiptStatusSuccessful, receipt.Status, "deploy tx submission failed")

	// Check upgrade tx submission was successful
	rt.includeL1BlockByTx(upgradeTx.Hash())
	receipt, err = rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), upgradeTx.Hash())
	require.NoError(rt.t, err)
	require.Equal(rt.t, types.ReceiptStatusSuccessful, receipt.Status, "upgrade tx submission failed")
}

// setupOutputSubmitted sets output submission by validator.
func (rt *Runtime) setupOutputSubmitted(version uint8) {
	// NOTE(chokobole): It is necessary to wait for one finalized (or safe if AllowNonFinalized
	// config is set) block to pass after each submission interval before submitting the output
	// root. For example, if the submission interval is set to 1800 blocks, the output root can
	// only be submitted at 1801 finalized blocks. In fact, the following code is designed to
	// create one or more finalized L2 blocks in order to pass the test. If Proto Dank Sharding
	// is introduced, the below code fix may no longer be necessary.
	rt.proceedWithBlocks(3)

	rt.depositToValPool(rt.validator)
	if version == valhelper.ValidatorV2 {
		rt.registerToValMgr(rt.validator)
	}

	// create l2 output submission transactions until there is nothing left to submit
	for {
		waitTime := rt.validator.CalculateWaitTime(rt.t)
		if waitTime > 0 {
			break
		}
		rt.submitL2Output()
	}
}

// setupChallenge sets challenge by challenger.
func (rt *Runtime) setupChallenge(challenger *L2Validator, version uint8) {
	// check that the output root that L1 stores is different from challenger's output root
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

	if version == valhelper.ValidatorV1 {
		rt.depositToValPool(challenger)

		// check bond amount before create challenge
		bond, err := rt.valPoolContract.GetBond(nil, rt.outputIndex)
		require.NoError(rt.t, err)
		require.Equal(rt.t, rt.dp.DeployConfig.ValidatorPoolRequiredBondAmount.ToInt(), bond.Amount)
	} else if version == valhelper.ValidatorV2 {
		rt.registerToValMgr(challenger)

		// check bond amount before create challenge
		bond, err := rt.assetMgrContract.TotalValidatorKroBonded(nil, challenger.address)
		require.NoError(rt.t, err)
		require.Equal(rt.t, uint64(0), bond.Uint64())
	}

	// submit create challenge tx
	rt.txHash = challenger.ActCreateChallenge(rt.t, rt.outputIndex)

	// include tx on L1
	rt.includeL1BlockBySender(challenger.address)

	// Check whether the submission was successful
	rt.receipt, err = rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), rt.txHash)
	require.NoError(rt.t, err)
	require.Equal(rt.t, types.ReceiptStatusSuccessful, rt.receipt.Status, "failed to create challenge")

	// check challenge created
	challenge, err := rt.colosseumContract.GetChallenge(nil, rt.outputIndex, challenger.address)
	require.NoError(rt.t, err)
	require.NotNil(rt.t, challenge, "challenge not found")

	if version == valhelper.ValidatorV1 {
		// check pending bond amount after create challenge
		pendingBond, err := rt.valPoolContract.GetPendingBond(nil, rt.outputIndex, challenger.address)
		require.NoError(rt.t, err)
		require.Equal(rt.t, pendingBond, rt.dp.DeployConfig.ValidatorPoolRequiredBondAmount.ToInt())

		// check challenger balance decreased
		cBal, err := rt.valPoolContract.BalanceOf(nil, challenger.address)
		require.NoError(rt.t, err)
		require.Equal(rt.t, new(big.Int).Sub(new(big.Int).SetInt64(defaultDepositAmount), rt.dp.DeployConfig.ValidatorPoolRequiredBondAmount.ToInt()), cBal)
	} else if version == valhelper.ValidatorV2 {
		// check bond amount after create challenge
		bond, err := rt.assetMgrContract.TotalValidatorKroBonded(nil, challenger.address)
		require.NoError(rt.t, err)
		require.Equal(rt.t, rt.dp.DeployConfig.AssetManagerBondAmount.ToInt().Uint64(), bond.Uint64())
	}
}

func (rt *Runtime) depositToValPool(validator *L2Validator) {
	// deposit bond for validator
	validator.ActDeposit(rt.t, defaultDepositAmount)
	rt.includeL1BlockBySender(validator.address)

	// check validator balance increased
	bal, err := rt.valPoolContract.BalanceOf(nil, validator.address)
	require.NoError(rt.t, err)
	require.Equal(rt.t, new(big.Int).SetUint64(defaultDepositAmount), bal)
}

func (rt *Runtime) registerToValMgr(validator *L2Validator) {
	minActivateAmount := rt.dp.DeployConfig.ValidatorManagerMinActivateAmount.ToInt()
	minActivateAmount = new(big.Int).Mul(minActivateAmount, common.Big256)

	// approve governance token
	validator.ActApprove(rt.t, minActivateAmount)
	rt.includeL1BlockBySender(validator.address)

	// register validator
	validator.ActRegisterValidator(rt.t, minActivateAmount)
	rt.includeL1BlockBySender(validator.address)

	// check validator status is active
	status := validator.getValidatorStatus(rt.t)
	require.Equal(rt.t, val.StatusActive, status)
}

// proceedWithBlocks proceeds n blocks.
func (rt *Runtime) proceedWithBlocks(n int) {
	for i := 0; i < n; i++ {
		// L1 block
		rt.miner.ActEmptyBlock(rt.t)
		// L2 block
		rt.sequencer.ActL1HeadSignal(rt.t)
		rt.sequencer.ActL2PipelineFull(rt.t)
		rt.sequencer.ActBuildToL1Head(rt.t)
		// submit and include in L1
		rt.batcher.ActSubmitAll(rt.t)
		rt.includeL1BlockBySender(rt.dp.Addresses.Batcher)
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
}

func (rt *Runtime) submitL2Output() {
	// submit to L1
	rt.validator.ActSubmitL2Output(rt.t)
	// include output on L1
	rt.includeL1BlockBySender(rt.validator.address)
	// Check submission was successful
	receipt, err := rt.miner.EthClient().TransactionReceipt(rt.t.Ctx(), rt.validator.LastSubmitL2OutputTx())
	require.NoError(rt.t, err)
	require.Equal(rt.t, types.ReceiptStatusSuccessful, receipt.Status, "submission failed")
}

func (rt *Runtime) fetchValidatorStatus(validator *L2Validator) (uint8, bool, *big.Int, *big.Int, *big.Int) {
	valStatus := validator.getValidatorStatus(rt.t)
	inJail := validator.isInJail(rt.t)
	slashingAmount, err := rt.assetMgrContract.BONDAMOUNT(nil)
	require.NoError(rt.t, err)
	validatorAsset, err := rt.assetMgrContract.TotalValidatorKro(nil, validator.address)
	require.NoError(rt.t, err)
	validatorAssetBonded, err := rt.assetMgrContract.TotalValidatorKroBonded(nil, validator.address)
	require.NoError(rt.t, err)

	return valStatus, inJail, validatorAsset, validatorAssetBonded, slashingAmount
}

func (rt *Runtime) includeL1BlockBySender(from common.Address) {
	rt.miner.includeL1BlockBySender(rt.t, from, rt.l1BlockDelta)
}

func (rt *Runtime) includeL1BlockByTx(txHash common.Hash) {
	rt.miner.includeL1BlockByTx(rt.t, txHash, rt.l1BlockDelta)
}
