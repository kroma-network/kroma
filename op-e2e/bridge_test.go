package op_e2e

import (
	"context"
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"

	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/receipts"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/transactions"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/kroma-network/kroma/kroma-bindings/bindings"
	"github.com/kroma-network/kroma/kroma-bindings/predeploys"
	"github.com/kroma-network/kroma/op-e2e/e2eutils/geth"
)

// TestERC20BridgeDeposits tests the the L1StandardBridge bridge ERC20
// functionality.
func TestERC20BridgeDeposits(t *testing.T) {
	InitParallel(t)

	cfg := DefaultSystemConfig(t)

	sys, err := cfg.Start(t)
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	log := testlog.Logger(t, log.LevelInfo)
	log.Info("genesis", "l2", sys.RollupConfig.Genesis.L2, "l1", sys.RollupConfig.Genesis.L1, "l2_time", sys.RollupConfig.Genesis.L2Time)

	l1Client := sys.Clients["l1"]
	l2Client := sys.Clients["sequencer"]

	opts, err := bind.NewKeyedTransactorWithChainID(sys.Cfg.Secrets.Alice, cfg.L1ChainIDBig())
	require.Nil(t, err)

	// Deploy WETH9
	weth9Address, tx, WETH9, err := bindings.DeployWETH9(opts, l1Client)
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(context.Background(), l1Client, tx.Hash())
	require.NoError(t, err, "Waiting for deposit tx on L1")

	// Get some WETH
	opts.Value = big.NewInt(params.Ether)
	tx, err = WETH9.Fallback(opts, []byte{})
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(context.Background(), l1Client, tx.Hash())
	require.NoError(t, err)
	opts.Value = nil
	wethBalance, err := WETH9.BalanceOf(&bind.CallOpts{}, opts.From)
	require.NoError(t, err)
	require.Equal(t, big.NewInt(params.Ether), wethBalance)

	// Deploy L2 WETH9
	l2Opts, err := bind.NewKeyedTransactorWithChainID(sys.Cfg.Secrets.Alice, cfg.L2ChainIDBig())
	require.NoError(t, err)
	kromaMintableTokenFactory, err := bindings.NewKromaMintableERC20Factory(predeploys.KromaMintableERC20FactoryAddr, l2Client)
	require.NoError(t, err)
	tx, err = kromaMintableTokenFactory.CreateKromaMintableERC20(l2Opts, weth9Address, "L2-WETH", "L2-WETH")
	require.NoError(t, err)
	rcpt, err := wait.ForReceiptOK(context.Background(), l2Client, tx.Hash())
	require.NoError(t, err)

	event, err := receipts.FindLog(rcpt.Logs, kromaMintableTokenFactory.ParseKromaMintableERC20Created)
	require.NoError(t, err, "Should emit ERC20Created event")

	// Approve WETH9 with the bridge
	tx, err = WETH9.Approve(opts, cfg.L1Deployments.L1StandardBridgeProxy, new(big.Int).SetUint64(math.MaxUint64))
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(context.Background(), l1Client, tx.Hash())
	require.NoError(t, err)

	// Bridge the WETH9
	l1StandardBridge, err := bindings.NewL1StandardBridge(cfg.L1Deployments.L1StandardBridgeProxy, l1Client)
	require.NoError(t, err)
	tx, err = transactions.PadGasEstimate(opts, 1.1, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return l1StandardBridge.BridgeERC20(opts, weth9Address, event.LocalToken, big.NewInt(100), 100000, []byte{})
	})
	require.NoError(t, err)
	depositReceipt, err := wait.ForReceiptOK(context.Background(), l1Client, tx.Hash())
	require.NoError(t, err)

	t.Log("Deposit through L1StandardBridge", "gas used", depositReceipt.GasUsed)

	// compute the deposit transaction hash + poll for it
	portal, err := bindings.NewKromaPortal(cfg.L1Deployments.KromaPortalProxy, l1Client)
	require.NoError(t, err)

	depositEvent, err := receipts.FindLog(depositReceipt.Logs, portal.ParseTransactionDeposited)
	require.NoError(t, err, "Should emit deposit event")

	depositTx, err := derive.UnmarshalDepositLogEvent(&depositEvent.Raw)
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(context.Background(), l2Client, types.NewTx(depositTx).Hash())
	require.NoError(t, err)

	// Ensure that the deposit went through
	kromaMintableToken, err := bindings.NewKromaMintableERC20(event.LocalToken, l2Client)
	require.NoError(t, err)

	// Should have balance on L2
	l2Balance, err := kromaMintableToken.BalanceOf(&bind.CallOpts{}, opts.From)
	require.NoError(t, err)
	require.Equal(t, l2Balance, big.NewInt(100))
}

// TestBridgeGovernanceToken tests the L1StandardBridge bridge GovernanceToken
// functionality.
func TestBridgeGovernanceToken(t *testing.T) {
	InitParallel(t)

	cfg := DefaultSystemConfig(t)
	cfg.DeployConfig.MintManagerMintActivatedBlock = (*hexutil.Big)(new(big.Int).SetUint64(0))
	cfg.DeployConfig.MintManagerRecipients = []common.Address{cfg.Secrets.Addresses().Alice}
	cfg.DeployConfig.MintManagerShares = []uint64{100000}

	sys, err := cfg.Start(t)
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	log := testlog.Logger(t, log.LvlInfo)
	log.Info("genesis", "l2", sys.RollupConfig.Genesis.L2, "l1", sys.RollupConfig.Genesis.L1, "l2_time", sys.RollupConfig.Genesis.L2Time)

	l1Client := sys.Clients["l1"]
	l2Client := sys.Clients["sequencer"]

	l1TokenAddr := sys.Cfg.L1Deployments.L1GovernanceTokenProxy
	l1BridgeAddr := cfg.L1Deployments.L1StandardBridgeProxy
	l2TokenAddr := predeploys.GovernanceTokenAddr
	l2BridgeAddr := predeploys.L2StandardBridgeAddr

	l1Opts, err := bind.NewKeyedTransactorWithChainID(sys.Cfg.Secrets.Bob, cfg.L1ChainIDBig())
	require.Nil(t, err)
	l2Opts, err := bind.NewKeyedTransactorWithChainID(sys.Cfg.Secrets.Bob, cfg.L2ChainIDBig())
	require.Nil(t, err)

	// Init bridge contracts
	l1Bridge, err := bindings.NewL1StandardBridge(l1BridgeAddr, l1Client)
	require.NoError(t, err)
	l2Bridge, err := bindings.NewL2StandardBridge(l2BridgeAddr, l2Client)
	require.NoError(t, err)

	l1Token, err := bindings.NewGovernanceToken(l1TokenAddr, l1Client)
	require.NoError(t, err)
	l2Token, err := bindings.NewGovernanceToken(l2TokenAddr, l2Client)
	require.NoError(t, err)

	// Approve GovernanceToken with the bridge on L1 and L2
	tx, err := l1Token.Approve(l1Opts, l1BridgeAddr, new(big.Int).SetUint64(math.MaxUint64))
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(context.Background(), l1Client, tx.Hash())
	require.NoError(t, err)
	tx, err = l2Token.Approve(l2Opts, l2BridgeAddr, new(big.Int).SetUint64(math.MaxUint64))
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(context.Background(), l2Client, tx.Hash())
	require.NoError(t, err)

	// Wait until Alice have enough tokens on L2
	_, err = geth.WaitForBlock(big.NewInt(10), l2Client, 20*time.Second)
	require.NoError(t, err)
	aliceL2Opts, err := bind.NewKeyedTransactorWithChainID(cfg.Secrets.Alice, cfg.L2ChainIDBig())
	require.Nil(t, err)
	bridgeAmount, err := l2Token.BalanceOf(&bind.CallOpts{}, aliceL2Opts.From)
	require.NoError(t, err)
	require.NotZero(t, bridgeAmount.Uint64())

	// Send Alice's tokens to Bob
	tx, err = transactions.PadGasEstimate(aliceL2Opts, 1.1, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return l2Token.Transfer(opts, l2Opts.From, bridgeAmount)
	})
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(context.Background(), l2Client, tx.Hash())
	require.NoError(t, err)

	bobL1Balance, err := l1Token.BalanceOf(&bind.CallOpts{}, l1Opts.From)
	require.NoError(t, err)
	bobL2Balance, err := l2Token.BalanceOf(&bind.CallOpts{}, l2Opts.From)
	require.NoError(t, err)
	require.Equal(t, bobL2Balance, bridgeAmount)

	l1Supply, err := l1Token.TotalSupply(&bind.CallOpts{})
	require.NoError(t, err)

	// Withdraw GovernanceToken to L1
	tx, err = transactions.PadGasEstimate(l2Opts, 1.1, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return l2Bridge.BridgeERC20(opts, l2TokenAddr, l1TokenAddr, bridgeAmount, 100000, []byte{})
	})
	require.NoError(t, err)
	receipt, err := wait.ForReceiptOK(context.Background(), l2Client, tx.Hash())
	require.NoError(t, err)
	proveReceipt, finalizeReceipt := ProveAndFinalizeWithdrawal(t, cfg, sys, "verifier", cfg.Secrets.Bob, receipt)
	require.Equal(t, types.ReceiptStatusSuccessful, proveReceipt.Status)
	require.Equal(t, types.ReceiptStatusSuccessful, finalizeReceipt.Status)

	// Withdrawal complete, Bob's token balance increases on L1.
	newBobL1Balance, err := l1Token.BalanceOf(&bind.CallOpts{}, l1Opts.From)
	require.NoError(t, err)
	require.Equal(t, bobL1Balance.Add(bobL1Balance, bridgeAmount).Cmp(newBobL1Balance), 0)
	// Bob's token balance decreases on L2.
	newBobL2Balance, err := l2Token.BalanceOf(&bind.CallOpts{}, l2Opts.From)
	require.NoError(t, err)
	require.Equal(t, bobL2Balance.Sub(bobL2Balance, bridgeAmount).Cmp(newBobL2Balance), 0)
	// Total supply increases on L1.
	newL1Supply, err := l1Token.TotalSupply(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, l1Supply.Add(l1Supply, bridgeAmount).Cmp(newL1Supply), 0)
	// It's difficult to verify the exact decrease in token total supply on L2,
	// as the issuance amount increases with each block

	// Deposit GovernanceToken to L2
	tx, err = transactions.PadGasEstimate(l1Opts, 1.1, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return l1Bridge.BridgeERC20(opts, l1TokenAddr, l2TokenAddr, bridgeAmount, 100000, []byte{})
	})
	require.NoError(t, err)
	depositReceipt, err := wait.ForReceiptOK(context.Background(), l1Client, tx.Hash())
	require.NoError(t, err)

	t.Log("Deposit through L1StandardBridge", "gas used", depositReceipt.GasUsed)

	// compute the deposit transaction hash + poll for it
	portal, err := bindings.NewKromaPortal(cfg.L1Deployments.KromaPortalProxy, l1Client)
	require.NoError(t, err)

	depIt, err := portal.FilterTransactionDeposited(&bind.FilterOpts{Start: 0}, nil, nil, nil)
	require.NoError(t, err)
	var depositEvent *bindings.KromaPortalTransactionDeposited
	for depIt.Next() {
		depositEvent = depIt.Event
	}
	require.NotNil(t, depositEvent)

	depositTx, err := derive.UnmarshalDepositLogEvent(&depositEvent.Raw)
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(context.Background(), l2Client, types.NewTx(depositTx).Hash())
	require.NoError(t, err)

	// Deposit complete, Bob's token balance decreases on L1.
	newBobL1Balance, err = l1Token.BalanceOf(&bind.CallOpts{}, l1Opts.From)
	require.NoError(t, err)
	require.Equal(t, bobL1Balance.Sub(bobL1Balance, bridgeAmount).Cmp(newBobL1Balance), 0)
	// Bob's token balance increases on L2.
	newBobL2Balance, err = l2Token.BalanceOf(&bind.CallOpts{}, l2Opts.From)
	require.NoError(t, err)
	require.Equal(t, bobL2Balance.Add(bobL2Balance, bridgeAmount).Cmp(newBobL2Balance), 0)
	// Total supply decreases on L1.
	newL1Supply, err = l1Token.TotalSupply(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, l1Supply.Sub(l1Supply, bridgeAmount).Cmp(newL1Supply), 0)
	// It's difficult to verify the exact increase in token total supply on L2,
	// as the issuance amount increases with each block
}
