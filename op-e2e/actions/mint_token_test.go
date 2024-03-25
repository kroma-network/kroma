package actions

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-node/rollup/sync"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/kroma-network/kroma/kroma-bindings/bindings"
	"github.com/kroma-network/kroma/kroma-bindings/predeploys"
)

var floorUnit = big.NewInt(1e18 / 1e8)

func TestMintToken(t *testing.T) {
	tests := []struct {
		name string
		f    func(gt *testing.T)
	}{
		{"BeforeActivation", BeforeActivation},
		{"ActivatedAtGenesis", ActivatedAtGenesis},
		{"ActivatedAfterGenesis", ActivatedAfterGenesis},
		{"UntilExhausted", UntilExhausted},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			test.f(t)
		})
	}
}

func setupTestEnv(t StatefulTesting, dp *e2eutils.DeployParams) (*L1Miner, *L2Sequencer, *L2Verifier, *ethclient.Client, *L2Batcher) {
	sd := e2eutils.Setup(t, dp, defaultAlloc)
	log := testlog.Logger(t, log.LvlDebug)

	miner, seqEngine, seq := setupSequencerTest(t, sd, log)
	verifEngine, verifier := setupVerifier(t, sd, log, miner.L1Client(t, sd.RollupCfg), &sync.Config{})
	rollupSeqCl := seq.RollupClient()
	verifCl := verifEngine.EthClient()

	batcher := NewL2Batcher(log, sd.RollupCfg, &BatcherCfg{
		MinL1TxSize: 0,
		MaxL1TxSize: 128_000,
		BatcherKey:  dp.Secrets.Batcher,
	}, rollupSeqCl, miner.EthClient(), seqEngine.EthClient(), seqEngine.EngineClient(t, sd.RollupCfg))

	seq.ActL2PipelineFull(t)
	verifier.ActL2PipelineFull(t)

	return miner, seq, verifier, verifCl, batcher
}

func BeforeActivation(gt *testing.T) {
	t := NewDefaultTesting(gt)

	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	dp.DeployConfig.MintManagerMintActivatedBlock = nil

	miner, seq, verifier, verifCl, batcher := setupTestEnv(t, dp)

	miner.ActEmptyBlock(t)
	seq.ActL1HeadSignal(t)
	seq.ActBuildToL1Head(t)
	batcher.ActSubmitAll(t)

	miner.ActL1StartBlock(12)(t)
	miner.ActL1IncludeTx(dp.Addresses.Batcher)(t)
	miner.ActL1EndBlock(t)

	verifier.ActL2PipelineFull(t)

	infoTx, err := verifCl.TransactionInBlock(t.Ctx(), verifier.L2Safe().Hash, 0)
	require.NoError(t, err)
	receipt, err := verifCl.TransactionReceipt(t.Ctx(), infoTx.Hash())
	require.NoError(t, err)
	mintEvents, err := parseMintEvents(receipt, verifCl)
	require.NoError(t, err)
	require.Zero(t, len(mintEvents), "mint event exists before activation")
}

func ActivatedAtGenesis(gt *testing.T) {
	t := NewDefaultTesting(gt)

	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	dp.DeployConfig.MintManagerMintActivatedBlock = (*hexutil.Big)(new(big.Int).SetUint64(0))

	miner, seq, verifier, verifCl, batcher := setupTestEnv(t, dp)

	miner.ActEmptyBlock(t)
	seq.ActL1HeadSignal(t)
	seq.ActBuildToL1Head(t)
	batcher.ActSubmitAll(t)
	miner.ActL1StartBlock(12)(t)
	miner.ActL1IncludeTx(dp.Addresses.Batcher)(t)
	miner.ActL1EndBlock(t)
	verifier.ActL2PipelineFull(t)

	infoTx, err := verifCl.TransactionInBlock(t.Ctx(), verifier.L2Safe().Hash, 0)
	require.NoError(t, err)
	receipt, err := verifCl.TransactionReceipt(t.Ctx(), infoTx.Hash())
	require.NoError(t, err)
	mintEvents, err := parseMintEvents(receipt, verifCl)
	require.NoError(t, err)
	require.Equal(t, len(mintEvents), len(dp.DeployConfig.MintManagerRecipients))

	epoch, _ := epochAndOffset(dp, receipt.BlockNumber)
	mintAmount := mintAmountPerBlock(dp, epoch)
	validateMint(t, dp, mintEvents, mintAmount)
}

func ActivatedAfterGenesis(gt *testing.T) {
	t := NewDefaultTesting(gt)

	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	offset := uint64(24)
	dp.DeployConfig.MintManagerMintActivatedBlock = (*hexutil.Big)(new(big.Int).SetUint64(offset))

	miner, seq, verifier, verifCl, batcher := setupTestEnv(t, dp)

	// Build the L1 chain to reach the activation block.
	miner.ActEmptyBlock(t)
	for i := 0; i < 2; i++ {
		seq.ActL1HeadSignal(t)
		seq.ActBuildToL1Head(t)
		batcher.ActSubmitAll(t)
		miner.ActL1StartBlock(12)(t)
		miner.ActL1IncludeTx(dp.Addresses.Batcher)(t)
		miner.ActL1EndBlock(t)
	}
	verifier.ActL2PipelineFull(t)

	infoTx, err := verifCl.TransactionInBlock(t.Ctx(), verifier.L2Safe().Hash, 0)
	require.NoError(t, err)
	receipt, err := verifCl.TransactionReceipt(t.Ctx(), infoTx.Hash())
	require.NoError(t, err)
	mintEvents, err := parseMintEvents(receipt, verifCl)
	require.NoError(t, err)
	require.Equal(t, len(mintEvents), len(dp.DeployConfig.MintManagerRecipients))

	// Check the initial issuance was minted correctly.
	epoch, offset := epochAndOffset(dp, receipt.BlockNumber)
	mintPerBlock := new(big.Int).Set(dp.DeployConfig.MintManagerInitMintPerBlock.ToInt())
	mintAmount := new(big.Int).Set(common.Big0)
	for i := uint64(1); i < epoch; i++ {
		mintPerBlock = mintAmountPerBlock(dp, epoch)
		blocks := new(big.Int).SetUint64(dp.DeployConfig.MintManagerSlidingWindowBlocks)
		mintAmount.Add(mintAmount, new(big.Int).Mul(mintPerBlock, blocks))
	}
	if offset > 0 {
		blocks := new(big.Int).SetUint64(offset)
		mintAmount.Add(mintAmount, new(big.Int).Mul(mintPerBlock, blocks))
	}
	validateMint(t, dp, mintEvents, mintAmount)

	// Build the L1 chain to verify that tokens are minted correctly after the initial issuance and with each block.
	miner.ActEmptyBlock(t)
	seq.ActL1HeadSignal(t)
	seq.ActBuildToL1Head(t)
	batcher.ActSubmitAll(t)
	miner.ActL1StartBlock(12)(t)
	miner.ActL1IncludeTx(dp.Addresses.Batcher)(t)
	miner.ActL1EndBlock(t)
	verifier.ActL2PipelineFull(t)

	infoTx, err = verifCl.TransactionInBlock(t.Ctx(), verifier.L2Safe().Hash, 0)
	require.NoError(t, err)
	receipt, err = verifCl.TransactionReceipt(t.Ctx(), infoTx.Hash())
	require.NoError(t, err)
	mintEvents, err = parseMintEvents(receipt, verifCl)
	require.NoError(t, err)
	require.Equal(t, len(mintEvents), len(dp.DeployConfig.MintManagerRecipients))

	epoch, _ = epochAndOffset(dp, receipt.BlockNumber)
	mintAmount = mintAmountPerBlock(dp, epoch)
	validateMint(t, dp, mintEvents, mintAmount)
}

func UntilExhausted(gt *testing.T) {
	t := NewDefaultTesting(gt)

	recipients := make([]common.Address, 20)
	shares := make([]uint64, len(recipients))
	for i := range recipients {
		recipients[i] = common.Address{0: 0xff, 19: byte(i + 1)}
		shares[i] = uint64(100000 / len(recipients))
	}

	dp := e2eutils.MakeDeployParams(t, defaultRollupTestParams)
	dp.DeployConfig.MintManagerMintActivatedBlock = (*hexutil.Big)(new(big.Int).SetUint64(0))
	dp.DeployConfig.MintManagerSlidingWindowBlocks = 1
	dp.DeployConfig.MintManagerDecayingFactor = 90000
	dp.DeployConfig.MintManagerRecipients = recipients
	dp.DeployConfig.MintManagerShares = shares

	miner, seq, verifier, verifCl, batcher := setupTestEnv(t, dp)

	mintManager, err := bindings.NewMintManagerCaller(predeploys.MintManagerAddr, verifCl)
	require.NoError(t, err)

	exhaustedAt := uint64(1)
	for {
		epoch, _ := epochAndOffset(dp, new(big.Int).SetUint64(exhaustedAt))
		mintAmount := mintAmountPerBlock(dp, epoch)
		if mintAmount.Uint64() == 0 {
			exhaustedAt--
			break
		}
		exhaustedAt++
		require.Less(gt, exhaustedAt, uint64(500), "exhausting block number is too large to test")
	}

	l1BlockNumber := exhaustedAt * dp.DeployConfig.L2BlockTime / dp.DeployConfig.L1BlockTime
	for i := uint64(1); i < l1BlockNumber; i++ {
		miner.ActEmptyBlock(t)
		seq.ActL1HeadSignal(t)
		seq.ActBuildToL1Head(t)
		batcher.ActSubmitAll(t)
		miner.ActL1StartBlock(12)(t)
		miner.ActL1IncludeTx(dp.Addresses.Batcher)(t)
		miner.ActL1EndBlock(t)
	}

	seq.ActL2PipelineFull(t)
	verifier.ActL2PipelineFull(t)

	// Ensure that the amount of tokens minted decreases over time,
	// and verify that SystemTxGas is not exceeded as gas consumption increases as the epoch increases.
	var prevMintAmount *big.Int
	for i := uint64(1); i < exhaustedAt; i++ {
		mintAmount, err := mintManager.MintAmountPerBlock(&bind.CallOpts{}, new(big.Int).SetUint64(i))
		require.NoError(t, err)
		if i > 1 {
			require.Equal(t, mintAmount.Cmp(prevMintAmount), -1)
		}
		prevMintAmount = mintAmount

		infoTx, err := verifCl.TransactionInBlock(t.Ctx(), seq.L2Unsafe().Hash, 0)
		require.NoError(t, err)
		receipt, err := verifCl.TransactionReceipt(t.Ctx(), infoTx.Hash())
		require.NoError(t, err)
		require.Equal(t, receipt.Status, types.ReceiptStatusSuccessful)
		require.Less(t, receipt.GasUsed, uint64(derive.RegolithSystemTxGas))
		mintEvents, err := parseMintEvents(receipt, verifCl)
		require.NoError(t, err)
		validateMint(t, dp, mintEvents, mintAmount)
	}

	// Ensure that the token is exhausted.
	infoTx, err := verifCl.TransactionInBlock(t.Ctx(), seq.L2Unsafe().Hash, 0)
	require.NoError(t, err)
	receipt, err := verifCl.TransactionReceipt(t.Ctx(), infoTx.Hash())
	require.NoError(t, err)
	require.Equal(t, receipt.Status, types.ReceiptStatusSuccessful)
	require.Less(t, receipt.GasUsed, uint64(derive.RegolithSystemTxGas))
	mintEvents, err := parseMintEvents(receipt, verifCl)
	require.NoError(t, err)
	require.Zero(t, len(mintEvents))
}

func parseMintEvents(receipt *types.Receipt, backend bind.ContractBackend) ([]*bindings.GovernanceTokenTransfer, error) {
	governanceToken, err := bindings.NewGovernanceToken(predeploys.GovernanceTokenAddr, backend)
	if err != nil {
		return nil, err
	}

	evts := make([]*bindings.GovernanceTokenTransfer, 0)
	for _, l := range receipt.Logs {
		evt, err := governanceToken.ParseTransfer(*l)
		if err != nil || evt.From.Cmp(common.Address{}) != 0 {
			continue
		}
		evts = append(evts, evt)
	}

	return evts, nil
}

func validateMint(t StatefulTesting, dp *e2eutils.DeployParams, events []*bindings.GovernanceTokenTransfer, amount *big.Int) {
	for i, evt := range events {
		expectedTo := dp.DeployConfig.MintManagerRecipients[i]
		require.Equal(t, expectedTo, evt.To)
		shares := new(big.Int).SetUint64(dp.DeployConfig.MintManagerShares[i])
		expectedValue := new(big.Int)
		expectedValue.Mul(shares, amount)
		expectedValue.Div(expectedValue, big.NewInt(1e5))
		require.Equal(t, expectedValue, evt.Value)
	}
}

func epochAndOffset(dp *e2eutils.DeployParams, blockNumber *big.Int) (uint64, uint64) {
	epoch := (blockNumber.Uint64()-1)/dp.DeployConfig.MintManagerSlidingWindowBlocks + 1
	offset := (blockNumber.Uint64()-1)%dp.DeployConfig.MintManagerSlidingWindowBlocks + 1

	return epoch, offset
}

func mintAmountPerBlock(dp *e2eutils.DeployParams, epoch uint64) *big.Int {
	decayingFactor := new(big.Int).SetUint64(dp.DeployConfig.MintManagerDecayingFactor)
	decayingDenom := new(big.Int).SetUint64(1e5)

	amount := new(big.Int).Set(dp.DeployConfig.MintManagerInitMintPerBlock.ToInt())
	for i := uint64(1); i < epoch; i++ {
		amount.Mul(amount, decayingFactor).Div(amount, decayingDenom)
		amount.Div(amount, floorUnit).Mul(amount, floorUnit)
	}

	return amount
}
