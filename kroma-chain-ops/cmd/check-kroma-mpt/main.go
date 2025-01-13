package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	oppredeploys "github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	op_service "github.com/ethereum-optimism/optimism/op-service"
	"github.com/ethereum-optimism/optimism/op-service/cliapp"
	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/dial"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/ethereum-optimism/optimism/op-service/opio"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/urfave/cli/v2"

	"github.com/kroma-network/kroma/kroma-bindings/bindings"
	"github.com/kroma-network/kroma/kroma-bindings/predeploys"
)

func main() {
	app := cli.NewApp()
	app.Name = "check-kroma-mpt"
	app.Usage = "Check Kroma MPT upgrade results."
	app.Description = "Check Kroma MPT upgrade results."
	app.Action = func(c *cli.Context) error {
		return errors.New("see sub-commands")
	}
	app.Writer = os.Stdout
	app.ErrWriter = os.Stderr
	app.Commands = []*cli.Command{
		{
			Name: "mpt",
			Subcommands: []*cli.Command{
				makeCommand("all", checkAll),
				makeCommand("is-system-tx", checkIsSystemTx),
				makeCommand("miner-addr", checkMinerAddr),
				makeCommand("fee-distribution", checkFeeDistribution),
				makeCommand("l1-block-addr-data", checkL1BlockAddrAndData),
				makeCommand("historical-rpc", checkHistoricalRPC),
				makeCommand("l1-block", checkL1Block),
				makeCommand("gpo", checkGPO),
				makeCommand("eip-6780-selfdestruct", checkSelfdestruct),
				makeCommand("upgrade-txs", checkUpgradeTxs),
			},
			Flags:  makeFlags(),
			Action: makeCommandAction(checkAll),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Application failed: %v\n", err)
		os.Exit(1)
	}
}

type actionEnv struct {
	log      log.Logger
	l1       *ethclient.Client
	l2       *ethclient.Client
	rollupCl *sources.RollupClient
	key      *ecdsa.PrivateKey
	addr     common.Address
}

type CheckAction func(ctx context.Context, env *actionEnv) error

var (
	prefix     = "CHECK_KROMA_MPT"
	EndpointL1 = &cli.StringFlag{
		Name:    "l1",
		Usage:   "L1 execution RPC endpoint",
		EnvVars: op_service.PrefixEnvVar(prefix, "L1"),
		Value:   "http://localhost:8545",
	}
	EndpointL2 = &cli.StringFlag{
		Name:    "l2",
		Usage:   "L2 execution RPC endpoint",
		EnvVars: op_service.PrefixEnvVar(prefix, "L2"),
		Value:   "http://localhost:9545",
	}
	EndpointRollup = &cli.StringFlag{
		Name:    "rollup",
		Usage:   "L2 rollup-node RPC endpoint",
		EnvVars: op_service.PrefixEnvVar(prefix, "ROLLUP"),
		Value:   "http://localhost:7545",
	}
	AccountKey = &cli.StringFlag{
		Name:    "account",
		Usage:   "Private key (hex-formatted string) of test account to perform test txs with",
		EnvVars: op_service.PrefixEnvVar(prefix, "ACCOUNT"),
	}
)

func makeFlags() []cli.Flag {
	flags := []cli.Flag{
		EndpointL1,
		EndpointL2,
		EndpointRollup,
		AccountKey,
	}
	return append(flags, oplog.CLIFlags(prefix)...)
}

func makeCommand(name string, fn CheckAction) *cli.Command {
	return &cli.Command{
		Name:   name,
		Action: makeCommandAction(fn),
		Flags:  cliapp.ProtectFlags(makeFlags()),
	}
}

func makeCommandAction(fn CheckAction) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		logCfg := oplog.ReadCLIConfig(c)
		logger := oplog.NewLogger(c.App.Writer, logCfg)

		c.Context = opio.CancelOnInterrupt(c.Context)
		l1Cl, err := ethclient.DialContext(c.Context, c.String(EndpointL1.Name))
		if err != nil {
			return fmt.Errorf("failed to dial L1 RPC: %w", err)
		}
		l2Cl, err := ethclient.DialContext(c.Context, c.String(EndpointL2.Name))
		if err != nil {
			return fmt.Errorf("failed to dial L2 RPC: %w", err)
		}
		rollupCl, err := dial.DialRollupClientWithTimeout(c.Context, time.Second*20, logger, c.String(EndpointRollup.Name))
		if err != nil {
			return fmt.Errorf("failed to dial rollup node RPC: %w", err)
		}
		key, err := crypto.HexToECDSA(c.String(AccountKey.Name))
		if err != nil {
			return fmt.Errorf("failed to parse test private key: %w", err)
		}
		if err := fn(c.Context, &actionEnv{
			log:      logger,
			l1:       l1Cl,
			l2:       l2Cl,
			rollupCl: rollupCl,
			key:      key,
			addr:     crypto.PubkeyToAddress(key.PublicKey),
		}); err != nil {
			return fmt.Errorf("command error: %w", err)
		}
		return nil
	}
}

type mptBlockWithParent struct {
	prevBlock *types.Block
	mptBlock  *types.Block
}

func getMPTBlockWithParent(ctx context.Context, env *actionEnv) (*mptBlockWithParent, error) {
	conf, err := env.rollupCl.RollupConfig(ctx)
	if err != nil {
		return nil, err
	}
	if conf.KromaMPTTime == nil {
		return nil, errors.New("kroma mpt time is required")
	}
	mptBlockNum, err := conf.TargetBlockNumber(*conf.KromaMPTTime)
	if err != nil {
		return nil, err
	}

	prevBlock, err := env.l2.BlockByNumber(ctx, big.NewInt(int64(mptBlockNum-1)))
	if err != nil {
		return nil, err
	}
	mptBlock, err := env.l2.BlockByNumber(ctx, big.NewInt(int64(mptBlockNum)))
	if err != nil {
		return nil, err
	}

	return &mptBlockWithParent{prevBlock, mptBlock}, nil
}

func checkIsSystemTx(ctx context.Context, env *actionEnv) error {
	blocks, err := getMPTBlockWithParent(ctx, env)
	if err != nil {
		return fmt.Errorf("failed to get mpt and parent block: %w", err)
	}

	for _, tx := range blocks.prevBlock.Transactions() {
		if tx.IsDepositTx() {
			depTxBytes, err := tx.MarshalBinary()
			if err != nil {
				return fmt.Errorf("marshal tx failed: %w", err)
			}
			isKromaDepTx, err := types.IsKromaDepositTx(depTxBytes[1:])
			if err != nil {
				return fmt.Errorf("check isKromaDepTx failed: %w", err)
			}
			if !isKromaDepTx {
				return errors.New("not kroma deposit tx before MPT")
			}
		}
	}

	for _, tx := range blocks.mptBlock.Transactions() {
		if tx.IsDepositTx() {
			depTxBytes, err := tx.MarshalBinary()
			if err != nil {
				return fmt.Errorf("marshal tx failed: %w", err)
			}
			isKromaDepTx, err := types.IsKromaDepositTx(depTxBytes[1:])
			if err != nil {
				return fmt.Errorf("check isKromaDepTx failed: %w", err)
			}
			if isKromaDepTx {
				return errors.New("kroma deposit tx after MPT")
			}
		}
	}

	env.log.Info("isSystemTransaction test: SUCCESS")
	return nil
}

func checkMinerAddr(ctx context.Context, env *actionEnv) error {
	blocks, err := getMPTBlockWithParent(ctx, env)
	if err != nil {
		return fmt.Errorf("failed to get mpt and parent block: %w", err)
	}

	if blocks.prevBlock.Coinbase().Cmp(common.Address{}) != 0 {
		return errors.New("coinbase address is not zero before MPT")
	}

	if blocks.mptBlock.Coinbase().Cmp(oppredeploys.SequencerFeeVaultAddr) != 0 {
		return errors.New("coinbase address is not sequencer fee vault after MPT")
	}

	env.log.Info("miner address test: SUCCESS")
	return nil
}

func checkFeeDistribution(ctx context.Context, env *actionEnv) error {
	to := common.Address{1, 2, 3, 5}
	receipt, err := execTx(ctx, &to, params.TxGas+100, 3*params.GWei, []byte("hello"), false, env)
	if err != nil {
		return fmt.Errorf("failed to execute tx: %w", err)
	}

	checkBalance := func(addr common.Address, prevBlockNum, blockNum *big.Int, increaseRequired bool) error {
		prevBal, err := env.l2.BalanceAt(ctx, addr, prevBlockNum)
		if err != nil {
			return err
		}
		bal, err := env.l2.BalanceAt(ctx, addr, blockNum)
		if err != nil {
			return err
		}
		if increaseRequired != (bal.Cmp(prevBal) == 1) {
			return fmt.Errorf("fee distribution unexpectedly, expected increasing: %t, but account: %s, prev: %d, after: %d",
				increaseRequired, addr, prevBal, bal)
		}
		return nil
	}

	// check fee distribution after MPT
	blockNum := receipt.BlockNumber
	prevBlockNum := new(big.Int).Sub(blockNum, common.Big1)
	err = checkBalance(predeploys.ProtocolVaultAddr, prevBlockNum, blockNum, false)
	if err != nil {
		return err
	}
	err = checkBalance(predeploys.L1FeeVaultAddr, prevBlockNum, blockNum, false)
	if err != nil {
		return err
	}
	err = checkBalance(oppredeploys.SequencerFeeVaultAddr, prevBlockNum, blockNum, true)
	if err != nil {
		return err
	}
	err = checkBalance(oppredeploys.BaseFeeVaultAddr, prevBlockNum, blockNum, true)
	if err != nil {
		return err
	}
	err = checkBalance(oppredeploys.L1FeeVaultAddr, prevBlockNum, blockNum, true)
	if err != nil {
		return err
	}

	env.log.Info("fee distribution test: SUCCESS")
	return nil
}

func checkL1BlockAddrAndData(ctx context.Context, env *actionEnv) error {
	blocks, err := getMPTBlockWithParent(ctx, env)
	if err != nil {
		return fmt.Errorf("failed to get mpt and parent block: %w", err)
	}

	prevL1InfoTx := blocks.prevBlock.Transactions()[0]
	mptL1InfoTx := blocks.mptBlock.Transactions()[0]

	if prevL1InfoTx.To().Cmp(predeploys.KromaL1BlockAddr) != 0 {
		return errors.New("l1 block addr is not kroma l1 block addr before MPT")
	}
	if mptL1InfoTx.To().Cmp(oppredeploys.L1BlockAddr) != 0 {
		return errors.New("l1 block addr is not op l1 block addr after MPT")
	}

	if len(prevL1InfoTx.Data()) != derive.L1InfoEcotoneLen {
		return fmt.Errorf("l1 info tx data length not matched before MPT, expected: %d, got: %d", derive.L1InfoEcotoneLen, len(prevL1InfoTx.Data()))
	}
	if len(mptL1InfoTx.Data()) != derive.L1InfoKromaMPTLen {
		return fmt.Errorf("l1 info tx data length not matched before MPT, expected: %d, got: %d", derive.L1InfoKromaMPTLen, len(mptL1InfoTx.Data()))
	}

	env.log.Info("L1Block address and tx data test: SUCCESS")
	return nil
}

func checkHistoricalRPC(ctx context.Context, env *actionEnv) error {
	blocks, err := getMPTBlockWithParent(ctx, env)
	if err != nil {
		return fmt.Errorf("failed to get mpt and parent block: %w", err)
	}

	rollupCfg, err := env.rollupCl.RollupConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve rollup config: %w", err)
	}
	l2RPC := client.NewBaseRPCClient(env.l2.Client())
	l2EthCl, err := sources.NewL2Client(l2RPC, env.log, nil, sources.L2ClientDefaultConfig(rollupCfg, false))
	if err != nil {
		return fmt.Errorf("failed to create eth client: %w", err)
	}

	genesisBlockNum := big.NewInt(rpc.EarliestBlockNumber.Int64())
	prevBlockNum := blocks.prevBlock.Number()
	_, err = env.l2.BalanceAt(ctx, env.addr, genesisBlockNum)
	if err != nil {
		return fmt.Errorf("failed to get balance at genesis: %w", err)
	}
	_, err = env.l2.BalanceAt(ctx, env.addr, prevBlockNum)
	if err != nil {
		return fmt.Errorf("failed to get balance at mpt previous block: %w", err)
	}

	_, err = l2EthCl.GetProof(ctx, predeploys.KromaL1BlockAddr, []common.Hash{}, "earliest")
	if err != nil {
		return fmt.Errorf("failed to get proof at genesis: %w", err)
	}
	_, err = l2EthCl.GetProof(ctx, predeploys.KromaL1BlockAddr, []common.Hash{}, blocks.prevBlock.Hash().Hex())
	if err != nil {
		return fmt.Errorf("failed to get proof at mpt previous block: %w", err)
	}

	_, err = env.l2.CodeAt(ctx, predeploys.KromaL1BlockAddr, genesisBlockNum)
	if err != nil {
		return fmt.Errorf("failed to get code at genesis: %w", err)
	}
	_, err = env.l2.CodeAt(ctx, predeploys.KromaL1BlockAddr, prevBlockNum)
	if err != nil {
		return fmt.Errorf("failed to get code at mpt previous block: %w", err)
	}

	_, err = l2EthCl.GetStorageAt(ctx, predeploys.KromaL1BlockAddr, common.Hash{}, "earliest")
	if err != nil {
		return fmt.Errorf("failed to get storage at genesis: %w", err)
	}
	_, err = l2EthCl.GetStorageAt(ctx, predeploys.KromaL1BlockAddr, common.Hash{}, blocks.prevBlock.Hash().Hex())
	if err != nil {
		return fmt.Errorf("failed to get storage at mpt previous block: %w", err)
	}

	msg := ethereum.CallMsg{
		From: env.addr,
		To:   &predeploys.KromaL1BlockAddr,
		Data: crypto.Keccak256([]byte("number()"))[:4],
	}
	_, err = env.l2.CallContract(ctx, msg, genesisBlockNum)
	if err != nil {
		return fmt.Errorf("failed to call at genesis: %w", err)
	}
	_, err = env.l2.CallContract(ctx, msg, prevBlockNum)
	if err != nil {
		return fmt.Errorf("failed to call at mpt previous block: %w", err)
	}

	_, err = env.l2.NonceAt(ctx, env.addr, genesisBlockNum)
	if err != nil {
		return fmt.Errorf("failed to get nonce at genesis: %w", err)
	}
	_, err = env.l2.NonceAt(ctx, env.addr, prevBlockNum)
	if err != nil {
		return fmt.Errorf("failed to get nonce at mpt previous block: %w", err)
	}

	env.log.Info("historical RPC test: SUCCESS")
	return nil
}

func checkL1Block(ctx context.Context, env *actionEnv) error {
	cl, err := bindings.NewL1Block(oppredeploys.L1BlockAddr, env.l2)
	if err != nil {
		return fmt.Errorf("failed to create bindings around L1Block contract: %w", err)
	}

	blobBaseFee, err := cl.BlobBaseFee(nil)
	if err != nil {
		return fmt.Errorf("failed to get blob basefee from L1Block contract: %w", err)
	}
	if big.NewInt(0).Cmp(blobBaseFee) == 0 {
		return errors.New("blob basefee must never be 0, EIP specifies minimum of 1")
	}

	rollupCfg, err := env.rollupCl.RollupConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve rollup config: %w", err)
	}
	l2RPC := client.NewBaseRPCClient(env.l2.Client())
	l2EthCl, err := sources.NewL2Client(l2RPC, env.log, nil, sources.L2ClientDefaultConfig(rollupCfg, false))
	if err != nil {
		return fmt.Errorf("failed to create eth client: %w", err)
	}

	prevValRewardScalar, err := l2EthCl.GetStorageAt(ctx, oppredeploys.L1BlockAddr, types.KromaL1BlobBaseFeeSlot, "latest")
	if err != nil {
		return fmt.Errorf("failed to get previous validatorRewardScalar storage: %w", err)
	}
	if big.NewInt(0).Cmp(prevValRewardScalar.Big()) != 0 {
		return errors.New("validatorRewardScalar must be 0")
	}

	env.log.Info("L1Block contract test: SUCCESS")
	return nil
}

func checkGPO(_ context.Context, env *actionEnv) error {
	cl, err := bindings.NewGasPriceOracle(predeploys.GasPriceOracleAddr, env.l2)
	if err != nil {
		return fmt.Errorf("failed to create bindings around GasPriceOracle contract: %w", err)
	}

	isKromaMPT, err := cl.IsKromaMPT(nil)
	if err != nil {
		return fmt.Errorf("failed to get Kroma MPT status: %w", err)
	}
	if !isKromaMPT {
		return fmt.Errorf("GPO is not set to Kroma MPT: %w", err)
	}

	l1BaseFee, err := cl.L1BaseFee(nil)
	if err != nil {
		return fmt.Errorf("failed to get l1 base fee: %w", err)
	}
	if l1BaseFee.Cmp(common.Big0) == 0 {
		return errors.New("l1 base fee should not be zero")
	}

	blobBaseFee, err := cl.BlobBaseFee(nil)
	if err != nil {
		return fmt.Errorf("failed to get blob base fee: %w", err)
	}
	if blobBaseFee.Cmp(common.Big0) == 0 {
		return errors.New("blob base fee should not be zero")
	}

	baseFeeScalar, err := cl.BaseFeeScalar(nil)
	if err != nil {
		return fmt.Errorf("failed to get base fee scalar: %w", err)
	}
	if baseFeeScalar == 0 {
		return errors.New("base fee scalar should not be zero")
	}

	blobBaseFeeScalar, err := cl.BlobBaseFeeScalar(nil)
	if err != nil {
		return fmt.Errorf("failed to get blob base fee scalar: %w", err)
	}
	if blobBaseFeeScalar == 0 {
		return errors.New("blob base fee scalar should not be zero")
	}

	env.log.Info("GasPriceOracle contract test: SUCCESS")
	return nil
}

// assuming a 0 (fail) or non-zero (success) on the stack, this performs a revert or self-destruct
func conditionalCode(data []byte) []byte {
	suffix := []byte{
		// add jump dest
		byte(vm.PUSH4),
		0xff, 0xff, 0xff, 0xff,
		byte(vm.JUMPI),
		// error case
		byte(vm.PUSH0),
		byte(vm.PUSH0),
		byte(vm.REVERT),
		// success case
		byte(vm.JUMPDEST),
		byte(vm.CALLER),
		byte(vm.SELFDESTRUCT),
		byte(vm.STOP),
	}
	binary.BigEndian.PutUint32(suffix[1:5], uint32(len(data))+9)
	out := make([]byte, 0, len(data)+len(suffix))
	out = append(out, data...)
	out = append(out, suffix...)
	return out
}

func execTx(
	ctx context.Context, to *common.Address, gas, value uint64, data []byte, expectRevert bool, env *actionEnv,
) (*types.Receipt, error) {
	nonce, err := env.l2.PendingNonceAt(ctx, env.addr)
	if err != nil {
		return nil, fmt.Errorf("pending nonce retrieval failed: %w", err)
	}
	head, err := env.l2.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve head header: %w", err)
	}

	tip := big.NewInt(params.GWei)
	maxFee := new(big.Int).Mul(head.BaseFee, big.NewInt(2))
	maxFee = maxFee.Add(maxFee, tip)

	chainID, err := env.l2.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chainID: %w", err)
	}
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID: chainID, Nonce: nonce,
		GasTipCap: tip, GasFeeCap: maxFee, Gas: gas, To: to, Value: big.NewInt(int64(value)), Data: data,
	})
	signer := types.NewCancunSigner(chainID)
	signedTx, err := types.SignTx(tx, signer, env.key)
	if err != nil {
		return nil, fmt.Errorf("failed to sign tx: %w", err)
	}

	env.log.Info("sending tx", "txhash", signedTx.Hash(), "to", to, "data", hexutil.Bytes(data))
	if err := env.l2.SendTransaction(ctx, signedTx); err != nil {
		return nil, fmt.Errorf("failed to send tx: %w", err)
	}
	for i := 0; i < 30; i++ {
		env.log.Info("checking confirmation...", "txhash", signedTx.Hash())
		receipt, err := env.l2.TransactionReceipt(ctx, signedTx.Hash())
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				env.log.Info("not found yet, waiting...")
				time.Sleep(time.Second)
				continue
			} else {
				return nil, fmt.Errorf("error while checking tx receipt: %w", err)
			}
		}
		if expectRevert {
			if receipt.Status == types.ReceiptStatusFailed {
				env.log.Info("tx reverted as expected", "txhash", signedTx.Hash())
				return receipt, nil
			} else {
				return nil, fmt.Errorf("tx %s unexpectedly completed without revert", signedTx.Hash())
			}
		} else {
			if receipt.Status == types.ReceiptStatusSuccessful {
				env.log.Info("tx confirmed", "txhash", signedTx.Hash())
				return receipt, nil
			} else {
				return nil, fmt.Errorf("tx %s failed", signedTx.Hash())
			}
		}
	}
	return nil, fmt.Errorf("failed to confirm tx: %s", signedTx.Hash())
}

func checkSelfdestruct(ctx context.Context, env *actionEnv) error {
	input := conditionalCode([]byte{
		// prepare code in memory
		byte(vm.PUSH2), // value
		byte(vm.CALLER),
		byte(vm.SELFDESTRUCT),
		byte(vm.PUSH1), // offset
		byte(vm.MSTORE),
		// create contract
		byte(vm.PUSH1), // size, just a 2 byte contract
		2,
		byte(vm.PUSH0),       // ETH value
		byte(vm.PUSH0),       // offset
		byte(vm.CREATE),      // pushes address on stack. Contract will immediately self-destruct
		byte(vm.EXTCODESIZE), // size should be 0
		byte(vm.ISZERO),      // check that it is
	})
	if _, err := execTx(ctx, nil, 500000, 0, input, false, env); err != nil {
		return err
	}

	env.log.Info("eip-6780 self-destruct test: SUCCESS")
	return nil
}

func checkUpgradeTxs(ctx context.Context, env *actionEnv) error {
	rollupCfg, err := env.rollupCl.RollupConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve rollup config: %w", err)
	}

	activationBlockNum := rollupCfg.Genesis.L2.Number +
		((*rollupCfg.KromaMPTTime - rollupCfg.Genesis.L2Time) / rollupCfg.BlockTime)
	env.log.Info("upgrade block num", "num", activationBlockNum)

	l2RPC := client.NewBaseRPCClient(env.l2.Client())
	l2EthCl, err := sources.NewL2Client(l2RPC, env.log, nil, sources.L2ClientDefaultConfig(rollupCfg, false))
	if err != nil {
		return fmt.Errorf("failed to create eth client: %w", err)
	}

	prevBlock, txs, err := l2EthCl.InfoAndTxsByNumber(ctx, activationBlockNum-1)
	if err != nil {
		return fmt.Errorf("failed to get activation previous block: %w", err)
	}

	if len(txs) != derive.KromaMPTUpgradeTxCount+1 {
		return fmt.Errorf("expected %d txs in Kroma MPT activation previous block, but got %d",
			derive.KromaMPTUpgradeTxCount+1, len(txs))
	}
	for i, tx := range txs {
		if !tx.IsDepositTx() {
			return fmt.Errorf("unexpected non-deposit tx in activation previous block, index %d, hash %s", i, tx.Hash())
		}
	}

	_, receipts, err := l2EthCl.FetchReceipts(ctx, prevBlock.Hash())
	if err != nil {
		return fmt.Errorf("failed to fetch receipts of activation previous block: %w", err)
	}
	for i, rec := range receipts {
		if rec.Status != types.ReceiptStatusSuccessful {
			return fmt.Errorf("failed tx receipt: %d", i)
		}
		switch i {
		case 1, 2, 3, 4, 5: // 5 implementation contracts deployment
			if rec.ContractAddress == (common.Address{}) {
				return fmt.Errorf("expected contract deployment, but got none")
			}
		case 6, 7, 8, 9, 10, 11: // proxy upgrades and setKromaMPT call
			if rec.ContractAddress != (common.Address{}) {
				return fmt.Errorf("unexpected contract deployment")
			}
		}
	}

	_, txs, err = l2EthCl.InfoAndTxsByNumber(ctx, activationBlockNum)
	if err != nil {
		return fmt.Errorf("failed to get activation block: %w", err)
	}

	if len(txs) != 1 {
		return fmt.Errorf("expected no txs other than system tx in Kroma MPT activation block, but got %d", len(txs))
	}

	env.log.Info("upgrade-txs receipts test: SUCCESS")
	return nil
}

func checkAll(ctx context.Context, env *actionEnv) error {
	if err := checkIsSystemTx(ctx, env); err != nil {
		return fmt.Errorf("is-system-tx error: %w", err)
	}
	if err := checkMinerAddr(ctx, env); err != nil {
		return fmt.Errorf("miner-addr error: %w", err)
	}
	if err := checkFeeDistribution(ctx, env); err != nil {
		return fmt.Errorf("fee-distribution error: %w", err)
	}
	if err := checkL1BlockAddrAndData(ctx, env); err != nil {
		return fmt.Errorf("l1-block-addr-data error: %w", err)
	}
	if err := checkHistoricalRPC(ctx, env); err != nil {
		return fmt.Errorf("historical-rpc error: %w", err)
	}
	if err := checkL1Block(ctx, env); err != nil {
		return fmt.Errorf("l1-block error: %w", err)
	}
	if err := checkGPO(ctx, env); err != nil {
		return fmt.Errorf("gpo error: %w", err)
	}
	if err := checkSelfdestruct(ctx, env); err != nil {
		return fmt.Errorf("eip-6780 selfdestruct error: %w", err)
	}
	if err := checkUpgradeTxs(ctx, env); err != nil {
		return fmt.Errorf("upgrade-txs error: %w", err)
	}

	env.log.Info("completed Kroma MPT feature tests successfully")
	return nil
}
