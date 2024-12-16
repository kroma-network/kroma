package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/holiman/uint256"
	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus/misc/eip4844"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"

	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	op_service "github.com/ethereum-optimism/optimism/op-service"
	"github.com/ethereum-optimism/optimism/op-service/cliapp"
	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/dial"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/ethereum-optimism/optimism/op-service/opio"
	"github.com/ethereum-optimism/optimism/op-service/retry"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/kroma-network/kroma/kroma-bindings/bindings"
	"github.com/kroma-network/kroma/kroma-bindings/predeploys"
)

func main() {
	app := cli.NewApp()
	app.Name = "check-ecotone"
	app.Usage = "Check Ecotone upgrade results."
	app.Description = "Check Ecotone upgrade results."
	app.Action = func(c *cli.Context) error {
		return errors.New("see sub-commands")
	}
	app.Writer = os.Stdout
	app.ErrWriter = os.Stderr
	app.Commands = []*cli.Command{
		{
			Name: "cancun",
			Subcommands: []*cli.Command{
				makeCommand("eip-1153-tstore", checkEIP1153),
				makeCommand("eip-4844-blobhash", checkBlobDataHash),
				makeCommand("eip-4844-precompile", check4844Precompile),
				makeCommand("eip-5656-mcopy", checkMcopy),
				/* [Kroma: START]
				makeCommand("eip-6780-selfdestruct", checkSelfdestruct),
				[Kroma: END] */
				makeCommand("eip-4844-blobtx", checkBlobTxDenial),
				makeCommand("eip-4788-root", checkBeaconBlockRoot),
				makeCommand("eip-4788-contract", check4788Contract),
				makeCommand("multicall3-contract", checkMulticall3Contract),
				makeCommand("all", checkAllCancun),
			},
			Flags:  makeFlags(),
			Action: makeCommandAction(checkAllCancun),
		},
		makeCommand("upgrade", checkUpgradeTxs),
		{
			Name: "contracts",
			Subcommands: []*cli.Command{
				makeCommand("l1block", checkL1Block),
				makeCommand("gpo", checkGPO),
			},
		},
		makeCommand("fees", checkL1Fees),
		makeCommand("all", checkALL),
		{
			Name: "gen-key",
			Action: func(c *cli.Context) error {
				key, err := crypto.GenerateKey()
				if err != nil {
					return err
				}
				fmt.Println("address: " + crypto.PubkeyToAddress(key.PublicKey).String())
				return crypto.SaveECDSA("hotkey.txt", key)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Application failed: %v\n", err)
		os.Exit(1)
	}
}

type actionEnv struct {
	log       log.Logger
	l1        *ethclient.Client
	l2        *ethclient.Client
	rollupCl  *sources.RollupClient
	key       *ecdsa.PrivateKey
	addr      common.Address
	gasUsed   uint64
	l1GasUsed uint64
}

func (ae *actionEnv) RecordGasUsed(rec *types.Receipt) {
	ae.gasUsed += rec.GasUsed
	ae.l1GasUsed += rec.L1GasUsed.Uint64()
	ae.log.Debug("Recorded tx receipt gas", "gas_used", rec.GasUsed, "l1_gas_used", rec.L1GasUsed)
}

type CheckAction func(ctx context.Context, env *actionEnv) error

var (
	prefix     = "CHECK_ECOTONE"
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
		/* [Kroma: START]
		byte(vm.SELFDESTRUCT),
		[Kroma: END] */
		byte(vm.STOP),
	}
	binary.BigEndian.PutUint32(suffix[1:5], uint32(len(data))+9)
	out := make([]byte, 0, len(data)+len(suffix))
	out = append(out, data...)
	out = append(out, suffix...)
	return out
}

func checkEIP1153(ctx context.Context, env *actionEnv) error {
	input := conditionalCode([]byte{
		// store 0xc0ffee at 0x42
		byte(vm.PUSH3),
		0xc0, 0xff, 0xee,
		byte(vm.PUSH1),
		0x42,
		byte(vm.TSTORE),
		// retrieve it
		byte(vm.PUSH1),
		0x42,
		byte(vm.TLOAD),
		// check value
		byte(vm.PUSH3),
		0xc0, 0xff, 0xee,
		byte(vm.EQ),
	})
	if err := execTx(ctx, nil, input, false, env); err != nil {
		return err
	}
	env.log.Info("eip-1153 transient storage test: success")
	return nil
}

func checkBlobDataHash(ctx context.Context, env *actionEnv) error {
	// revert on non-blob tx
	input := []byte{
		byte(vm.BLOBHASH),
	}
	if err := execTx(ctx, nil, input, true, env); err != nil {
		return err
	}
	env.log.Info("4844 blob-data-hash test: success")
	return nil
}

// Deploy a contract that calls the EIP-4844 blob verification precompile with valid proof,
// and check if it verifies the proof successfully.
func check4844Precompile(ctx context.Context, env *actionEnv) error {
	var x eth.Blob
	if err := x.FromData(eth.Data("remember ethers phoenix")); err != nil {
		return fmt.Errorf("failed to construct blob: %w", err)
	}
	commitment, err := x.ComputeKZGCommitment()
	if err != nil {
		return fmt.Errorf("failed to compute commitment: %w", err)
	}
	point := kzg4844.Point{}
	proof, claim, err := kzg4844.ComputeProof(kzg4844.Blob(x), point)
	if err != nil {
		return fmt.Errorf("failed to compute proof: %w", err)
	}
	versionedHash := eth.KZGToVersionedHash(commitment)

	var inner []byte
	mstore32 := func(v []byte, offset uint8) {
		if len(v) != 32 {
			panic("invalid v")
		}
		inner = append(inner, byte(vm.PUSH32))
		inner = append(inner, v...)
		inner = append(inner, byte(vm.PUSH1), offset, byte(vm.MSTORE))
	}

	// prepare input in memory, following EIP-4844:
	//    versioned_hash = input[:32]
	//    z = input[32:64]
	//    y = input[64:96]
	//    commitment = input[96:144]
	//    proof = input[144:192]
	//
	// the call verify p(z) = y
	mstore32(versionedHash[:], 0) // versioned hash
	mstore32(point[:], 32)        // z
	mstore32(claim[:], 64)        // y
	mstore32(commitment[0:32], 96)
	mstore32(append(append([]byte{}, commitment[32:48]...), proof[0:16]...), 96+32)
	mstore32(proof[16:48], 144+16)
	env.log.Info(fmt.Sprintf("4844 precompile test: verifying p(%x) == %x for commitment %x proof %x", point[:], claim[:], commitment[:], proof[:]))

	inner = append(inner, []byte{
		byte(vm.PUSH0), // retSize
		byte(vm.PUSH0), // retOffset
		byte(vm.PUSH1), // argsSize
		byte(192),
		byte(vm.PUSH0), // argsOffset
		byte(vm.PUSH1), // precompile address
		byte(0x0A),
		byte(vm.GAS),        // gas (supply all gas)
		byte(vm.STATICCALL), // run call
		// 0 will be on stack if call reverts, 1 otherwise
		// now get the return-data size, and multiply it. To ensure size == 0 counts as failed call.
		byte(vm.RETURNDATASIZE),
		byte(vm.MUL),
	}...)

	input := conditionalCode(inner)
	if err := execTx(ctx, nil, input, false, env); err != nil {
		return fmt.Errorf("precompile check failed: %w", err)
	}
	env.log.Info("eip-4844 precompile test: success")
	return nil
}

func check4788Contract(ctx context.Context, env *actionEnv) error {
	head, err := env.l2.HeaderByNumber(ctx, nil)
	if err != nil {
		return err
	}
	// reverts on all-0 time input
	if err := execTx(ctx, &predeploys.EIP4788ContractAddr, make([]byte, 32), true, env); err != nil {
		return fmt.Errorf("expected revert on empty input: %w", err)
	}

	conf, err := env.rollupCl.RollupConfig(ctx)
	if err != nil {
		return fmt.Errorf("config retrieval failed: %w", err)
	}
	t := head.Time
	alignment := head.Time % conf.BlockTime
	for i := 0; i < 20; i++ {
		ti := t - uint64(i)
		if !conf.IsEcotone(ti) {
			continue
		}
		env.log.Info("Beacon block root query timestamp", "query_timestamp", ti)
		// revert when timestamp doesn't exist (when not aligned it won't exist),
		// or when we call it at the activation block and the contract was newly deployed
		// (the beacon block root is processed at the start of the block,
		// but the contract might not exist yet, during activation).
		revert := ti%conf.BlockTime != alignment
		if conf.IsEcotoneActivationBlock(ti) {
			// if the contract already existed, then we deployed the nonce=1 contract during upgrade
			code, err := env.l2.CodeAt(ctx, crypto.CreateAddress(derive.EIP4788From, 1), nil)
			if err != nil {
				return fmt.Errorf("failed to check code: %w", err)
			}
			revert = revert || len(code) == 0
		}
		input := new(uint256.Int).SetUint64(ti).Bytes32()
		if err := execTx(ctx, &predeploys.EIP4788ContractAddr, input[:], revert, env); err != nil {
			return fmt.Errorf("failed at t = %d", ti)
		}
	}
	env.log.Info("eip-4788 beacon block-roots contract test: success")
	return nil
}

// [Kroma: START]
func checkMulticall3Contract(ctx context.Context, env *actionEnv) error {
	multicall3Address := common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")
	multicall3Code := common.Hex2Bytes("6080604052600436106100f35760003560e01c80634d2301cc1161008a578063a8b0574e11610059578063a8b0574e1461025a578063bce38bd714610275578063c3077fa914610288578063ee82ac5e1461029b57600080fd5b80634d2301cc146101ec57806372425d9d1461022157806382ad56cb1461023457806386d516e81461024757600080fd5b80633408e470116100c65780633408e47014610191578063399542e9146101a45780633e64a696146101c657806342cbb15c146101d957600080fd5b80630f28c97d146100f8578063174dea711461011a578063252dba421461013a57806327e86d6e1461015b575b600080fd5b34801561010457600080fd5b50425b6040519081526020015b60405180910390f35b61012d610128366004610a85565b6102ba565b6040516101119190610bbe565b61014d610148366004610a85565b6104ef565b604051610111929190610bd8565b34801561016757600080fd5b50437fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0140610107565b34801561019d57600080fd5b5046610107565b6101b76101b2366004610c60565b610690565b60405161011193929190610cba565b3480156101d257600080fd5b5048610107565b3480156101e557600080fd5b5043610107565b3480156101f857600080fd5b50610107610207366004610ce2565b73ffffffffffffffffffffffffffffffffffffffff163190565b34801561022d57600080fd5b5044610107565b61012d610242366004610a85565b6106ab565b34801561025357600080fd5b5045610107565b34801561026657600080fd5b50604051418152602001610111565b61012d610283366004610c60565b61085a565b6101b7610296366004610a85565b610a1a565b3480156102a757600080fd5b506101076102b6366004610d18565b4090565b60606000828067ffffffffffffffff8111156102d8576102d8610d31565b60405190808252806020026020018201604052801561031e57816020015b6040805180820190915260008152606060208201528152602001906001900390816102f65790505b5092503660005b8281101561047757600085828151811061034157610341610d60565b6020026020010151905087878381811061035d5761035d610d60565b905060200281019061036f9190610d8f565b6040810135958601959093506103886020850185610ce2565b73ffffffffffffffffffffffffffffffffffffffff16816103ac6060870187610dcd565b6040516103ba929190610e32565b60006040518083038185875af1925050503d80600081146103f7576040519150601f19603f3d011682016040523d82523d6000602084013e6103fc565b606091505b50602080850191909152901515808452908501351761046d577f08c379a000000000000000000000000000000000000000000000000000000000600052602060045260176024527f4d756c746963616c6c333a2063616c6c206661696c656400000000000000000060445260846000fd5b5050600101610325565b508234146104e6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f4d756c746963616c6c333a2076616c7565206d69736d6174636800000000000060448201526064015b60405180910390fd5b50505092915050565b436060828067ffffffffffffffff81111561050c5761050c610d31565b60405190808252806020026020018201604052801561053f57816020015b606081526020019060019003908161052a5790505b5091503660005b8281101561068657600087878381811061056257610562610d60565b90506020028101906105749190610e42565b92506105836020840184610ce2565b73ffffffffffffffffffffffffffffffffffffffff166105a66020850185610dcd565b6040516105b4929190610e32565b6000604051808303816000865af19150503d80600081146105f1576040519150601f19603f3d011682016040523d82523d6000602084013e6105f6565b606091505b5086848151811061060957610609610d60565b602090810291909101015290508061067d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f4d756c746963616c6c333a2063616c6c206661696c656400000000000000000060448201526064016104dd565b50600101610546565b5050509250929050565b43804060606106a086868661085a565b905093509350939050565b6060818067ffffffffffffffff8111156106c7576106c7610d31565b60405190808252806020026020018201604052801561070d57816020015b6040805180820190915260008152606060208201528152602001906001900390816106e55790505b5091503660005b828110156104e657600084828151811061073057610730610d60565b6020026020010151905086868381811061074c5761074c610d60565b905060200281019061075e9190610e76565b925061076d6020840184610ce2565b73ffffffffffffffffffffffffffffffffffffffff166107906040850185610dcd565b60405161079e929190610e32565b6000604051808303816000865af19150503d80600081146107db576040519150601f19603f3d011682016040523d82523d6000602084013e6107e0565b606091505b506020808401919091529015158083529084013517610851577f08c379a000000000000000000000000000000000000000000000000000000000600052602060045260176024527f4d756c746963616c6c333a2063616c6c206661696c656400000000000000000060445260646000fd5b50600101610714565b6060818067ffffffffffffffff81111561087657610876610d31565b6040519080825280602002602001820160405280156108bc57816020015b6040805180820190915260008152606060208201528152602001906001900390816108945790505b5091503660005b82811015610a105760008482815181106108df576108df610d60565b602002602001015190508686838181106108fb576108fb610d60565b905060200281019061090d9190610e42565b925061091c6020840184610ce2565b73ffffffffffffffffffffffffffffffffffffffff1661093f6020850185610dcd565b60405161094d929190610e32565b6000604051808303816000865af19150503d806000811461098a576040519150601f19603f3d011682016040523d82523d6000602084013e61098f565b606091505b506020830152151581528715610a07578051610a07576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f4d756c746963616c6c333a2063616c6c206661696c656400000000000000000060448201526064016104dd565b506001016108c3565b5050509392505050565b6000806060610a2b60018686610690565b919790965090945092505050565b60008083601f840112610a4b57600080fd5b50813567ffffffffffffffff811115610a6357600080fd5b6020830191508360208260051b8501011115610a7e57600080fd5b9250929050565b60008060208385031215610a9857600080fd5b823567ffffffffffffffff811115610aaf57600080fd5b610abb85828601610a39565b90969095509350505050565b6000815180845260005b81811015610aed57602081850181015186830182015201610ad1565b81811115610aff576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600082825180855260208086019550808260051b84010181860160005b84811015610bb1578583037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001895281518051151584528401516040858501819052610b9d81860183610ac7565b9a86019a9450505090830190600101610b4f565b5090979650505050505050565b602081526000610bd16020830184610b32565b9392505050565b600060408201848352602060408185015281855180845260608601915060608160051b870101935082870160005b82811015610c52577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa0888703018452610c40868351610ac7565b95509284019290840190600101610c06565b509398975050505050505050565b600080600060408486031215610c7557600080fd5b83358015158114610c8557600080fd5b9250602084013567ffffffffffffffff811115610ca157600080fd5b610cad86828701610a39565b9497909650939450505050565b838152826020820152606060408201526000610cd96060830184610b32565b95945050505050565b600060208284031215610cf457600080fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610bd157600080fd5b600060208284031215610d2a57600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81833603018112610dc357600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610e0257600080fd5b83018035915067ffffffffffffffff821115610e1d57600080fd5b602001915036819003821315610a7e57600080fd5b8183823760009101908152919050565b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc1833603018112610dc357600080fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa1833603018112610dc357600080fdfea2646970667358221220bb2b5c71a328032f97c676ae39a1ec2148d3e5d6f73d95e9b17910152d61f16264736f6c634300080c0033")
	multicall3CodeHash := crypto.Keccak256Hash(multicall3Code)

	multicall3Contract, err := env.l2.CodeAt(ctx, multicall3Address, nil)
	if err != nil {
		return fmt.Errorf("failed to retrieve multicall3 contract code: %w", err)
	}
	codeHash := crypto.Keccak256Hash(multicall3Contract)
	if codeHash != multicall3CodeHash {
		return fmt.Errorf("unexpected multiucall3 contract code: %v", err)
	}
	env.log.Info("multicall3 contract test: success")
	return nil
}

// [Kroma: END]

func checkMcopy(ctx context.Context, env *actionEnv) error {
	input := conditionalCode([]byte{
		// push info & mstore it
		byte(vm.PUSH3),
		0xc0, 0xff, 0xee,
		byte(vm.PUSH0), // store at 0
		byte(vm.MSTORE),
		// copy the memory
		byte(vm.PUSH1), // length
		0x2,            // only copy the C0FF part
		byte(vm.PUSH1), // src
		32 - 3,         // right-aligned bytes3
		byte(vm.PUSH1), // dst
		0x42,
		byte(vm.MCOPY),
		byte(vm.PUSH1),  // copy from destination
		0x42 - (32 - 3), // a little to the left, so it's left-padded
		byte(vm.MLOAD),  // load the memory from copied location
		// check if it matches, with zero 3rd byte
		byte(vm.PUSH3),
		0xc0, 0xff, 0x00,
		byte(vm.EQ),
	})
	if err := execTx(ctx, nil, input, false, env); err != nil {
		return err
	}
	env.log.Info("eip-5656 mcopy test: success")
	return nil
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
	if err := execTx(ctx, nil, input, false, env); err != nil {
		return err
	}
	env.log.Info("eip-6780 self-destruct test: success")
	return nil
}

func execTx(ctx context.Context, to *common.Address, data []byte, expectRevert bool, env *actionEnv) error {
	nonce, err := env.l2.PendingNonceAt(ctx, env.addr)
	if err != nil {
		return fmt.Errorf("pending nonce retrieval failed: %w", err)
	}
	head, err := env.l2.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("failed to retrieve head header: %w", err)
	}

	tip := big.NewInt(params.GWei)
	maxFee := new(big.Int).Mul(head.BaseFee, big.NewInt(2))
	maxFee = maxFee.Add(maxFee, tip)

	chainID, err := env.l2.ChainID(ctx)
	if err != nil {
		return fmt.Errorf("failed to get chainID: %w", err)
	}
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID: chainID, Nonce: nonce,
		GasTipCap: tip, GasFeeCap: maxFee, Gas: 500000, To: to, Data: data,
	})
	signer := types.NewCancunSigner(chainID)
	signedTx, err := types.SignTx(tx, signer, env.key)
	if err != nil {
		return fmt.Errorf("failed to sign tx: %w", err)
	}

	env.log.Info("sending tx", "txhash", signedTx.Hash(), "to", to, "data", hexutil.Bytes(data))
	if err := env.l2.SendTransaction(ctx, signedTx); err != nil {
		return fmt.Errorf("failed to send tx: %w", err)
	}
	for i := 0; i < 30; i++ {
		env.log.Info("checking confirmation...", "txhash", signedTx.Hash())
		receipt, err := env.l2.TransactionReceipt(context.Background(), signedTx.Hash())
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				env.log.Info("not found yet, waiting...")
				time.Sleep(time.Second)
				continue
			} else {
				return fmt.Errorf("error while checking tx receipt: %w", err)
			}
		}
		env.RecordGasUsed(receipt)
		if expectRevert {
			if receipt.Status == types.ReceiptStatusFailed {
				env.log.Info("tx reverted as expected", "txhash", signedTx.Hash())
				return nil
			} else {
				return fmt.Errorf("tx %s unexpectedly completed without revert", signedTx.Hash())
			}
		} else {
			if receipt.Status == types.ReceiptStatusSuccessful {
				env.log.Info("tx confirmed", "txhash", signedTx.Hash())
				return nil
			} else {
				return fmt.Errorf("tx %s failed", signedTx.Hash())
			}
		}
	}
	return fmt.Errorf("failed to confirm tx: %s", signedTx.Hash())
}

func checkBlobTxDenial(ctx context.Context, env *actionEnv) error {
	// verify we cannot submit a blob tx to RPC
	var blob eth.Blob
	_, err := rand.Read(blob[:])
	if err != nil {
		return fmt.Errorf("failed randomnes: %w", err)
	}
	// get the field-elements into a valid range
	for i := 0; i < 4096; i++ {
		blob[32*i] &= 0b0011_1111
	}
	sidecar, blobHashes, err := txmgr.MakeSidecar([]*eth.Blob{&blob})
	if err != nil {
		return fmt.Errorf("failed to make sidecar: %w", err)
	}
	latestHeader, err := env.l1.HeaderByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to get header: %w", err)
	}
	if latestHeader.ExcessBlobGas == nil {
		return fmt.Errorf("the L1 block %s (time %d) is not ecotone yet", latestHeader.Hash(), latestHeader.Time)
	}
	blobBaseFee := eip4844.CalcBlobFee(*latestHeader.ExcessBlobGas)
	blobFeeCap := new(uint256.Int).Mul(uint256.NewInt(2), uint256.MustFromBig(blobBaseFee))
	if blobFeeCap.Lt(uint256.NewInt(params.GWei)) { // ensure we meet 1 gwei geth tx-pool minimum
		blobFeeCap = uint256.NewInt(params.GWei)
	}
	gasTipCap := big.NewInt(2 * params.GWei)
	gasFeeCap := new(big.Int).Add(gasTipCap, new(big.Int).Mul(latestHeader.BaseFee, big.NewInt(2)))

	nonce, err := env.l2.PendingNonceAt(ctx, env.addr)
	if err != nil {
		return fmt.Errorf("failed to get pending nonce: %w", err)
	}
	rollupCfg, err := env.rollupCl.RollupConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve rollup config: %w", err)
	}
	txData := &types.BlobTx{
		To:         rollupCfg.BatchInboxAddress,
		Data:       nil,
		Gas:        params.TxGas, // intrinsic gas only
		BlobHashes: blobHashes,
		Sidecar:    sidecar,
		ChainID:    uint256.MustFromBig(rollupCfg.L2ChainID),
		GasTipCap:  uint256.MustFromBig(gasTipCap),
		GasFeeCap:  uint256.MustFromBig(gasFeeCap),
		BlobFeeCap: blobFeeCap,
		Value:      uint256.NewInt(0),
		Nonce:      nonce,
	}
	// bypass signer filter by creating it manually and using the L2 chain ID
	signer := types.NewCancunSigner(rollupCfg.L2ChainID)
	tx, err := types.SignNewTx(env.key, signer, txData)
	if err != nil {
		return fmt.Errorf("failed to sign blob tx: %w", err)
	}
	err = env.l2.SendTransaction(ctx, tx)
	if err == nil {
		return errors.New("expected tx error, but got none")
	}
	if !strings.Contains(err.Error(), "transaction type not supported") {
		return fmt.Errorf("unexpected tx submission error: %w", err)
	}
	env.log.Info("blob-tx denial test: success")
	return nil
}

func checkBeaconBlockRoot(ctx context.Context, env *actionEnv) error {
	latest, err := env.l2.HeaderByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to get latest block: %w", err)
	}
	if latest.ParentBeaconRoot == nil {
		return fmt.Errorf("block %d misses beacon block root", latest.Number)
	}
	beaconBlockRootsContract, err := env.l2.CodeAt(ctx, predeploys.EIP4788ContractAddr, nil)
	if err != nil {
		return fmt.Errorf("failed to retrieve beacon block root contract code: %w", err)
	}
	codeHash := crypto.Keccak256Hash(beaconBlockRootsContract)
	if codeHash != predeploys.EIP4788ContractCodeHash {
		return fmt.Errorf("unexpected 4788 contract code: %w", err)
	}

	rollupCfg, err := env.rollupCl.RollupConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve rollup config: %w", err)
	}
	l2RPC := client.NewBaseRPCClient(env.l2.Client())
	l2EthCl, err := sources.NewL2Client(l2RPC, env.log, nil,
		sources.L2ClientDefaultConfig(rollupCfg, false))
	if err != nil {
		return fmt.Errorf("failed to create eth client")
	}
	result, err := l2EthCl.GetProof(ctx, predeploys.EIP4788ContractAddr, nil, eth.Unsafe)
	if err != nil {
		return fmt.Errorf("failed to get account proof to inspect storage-root")
	}
	if result.StorageHash == types.EmptyRootHash {
		return fmt.Errorf("expected contract storage to be set, but got none (%s)",
			result.StorageHash)
	}

	payload, err := l2EthCl.PayloadByLabel(ctx, eth.Unsafe)
	if err != nil {
		return fmt.Errorf("failed to get head ref: %w", err)
	}
	if payload.ParentBeaconBlockRoot == nil {
		return fmt.Errorf("payload %s misses parent beacon block root", payload.ExecutionPayload.ID())
	}
	headRef, err := derive.PayloadToBlockRef(rollupCfg, payload.ExecutionPayload)
	if err != nil {
		return fmt.Errorf("failed to convert to block-ref: %w", err)
	}
	l1Header, err := retry.Do(ctx, 5, retry.Fixed(time.Second*12), func() (*types.Header, error) {
		env.log.Info("retrieving L1 origin...", "l1", headRef.L1Origin)
		return env.l1.HeaderByHash(ctx, headRef.L1Origin.Hash)
	})
	if err != nil {
		return fmt.Errorf("failed to retrieve L1 origin %s of L2 block %s: %w", headRef.L1Origin, headRef, err)
	}
	var l1ParentBeaconBlockRoot common.Hash // zero before Dencun activates on L1
	if l1Header.ParentBeaconRoot != nil {
		l1ParentBeaconBlockRoot = *l1Header.ParentBeaconRoot
	}
	if l1ParentBeaconBlockRoot != *payload.ParentBeaconBlockRoot {
		return fmt.Errorf("parent beacon block root mismatch, L1: %s, L2: %s",
			l1ParentBeaconBlockRoot, *payload.ParentBeaconBlockRoot)
	}
	env.log.Info("beacon-block-root block-content test: success")
	return nil
}

func checkAllCancun(ctx context.Context, env *actionEnv) error {
	if err := checkEIP1153(ctx, env); err != nil {
		return fmt.Errorf("eip-1153 error: %w", err)
	}
	if err := checkBlobDataHash(ctx, env); err != nil {
		return fmt.Errorf("eip-4844 blobhash error: %w", err)
	}
	if err := check4844Precompile(ctx, env); err != nil {
		return fmt.Errorf("eip-4844 precompile error: %w", err)
	}
	if err := checkMcopy(ctx, env); err != nil {
		return fmt.Errorf("eip-5656 mcopy error: %w", err)
	}
	/* [Kroma: START]
	if err := checkSelfdestruct(ctx, env); err != nil {
		return fmt.Errorf("eip-6780 selfdestruct error: %w", err)
	}
	[Kroma: END] */
	if err := checkBlobTxDenial(ctx, env); err != nil {
		return fmt.Errorf("eip-4844 blob-tx denial error: %w", err)
	}
	if err := checkBeaconBlockRoot(ctx, env); err != nil {
		return fmt.Errorf("eip-4788 beacon-block-roots error: %w", err)
	}
	if err := check4788Contract(ctx, env); err != nil {
		return fmt.Errorf("eip-4788 contract check error: %w", err)
	}
	if err := checkMulticall3Contract(ctx, env); err != nil {
		return fmt.Errorf("multicall3 contract check error: %w", err)
	}
	env.log.Info("completed Cancun feature tests successfully")
	return nil
}

func checkUpgradeTxs(ctx context.Context, env *actionEnv) error {
	rollupCfg, err := env.rollupCl.RollupConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve rollup config: %w", err)
	}

	activationBlockNum := rollupCfg.Genesis.L2.Number +
		((*rollupCfg.EcotoneTime - rollupCfg.Genesis.L2Time) / rollupCfg.BlockTime)
	env.log.Info("upgrade block num", "num", activationBlockNum)
	l2RPC := client.NewBaseRPCClient(env.l2.Client())
	l2EthCl, err := sources.NewL2Client(l2RPC, env.log, nil,
		sources.L2ClientDefaultConfig(rollupCfg, false))
	if err != nil {
		return fmt.Errorf("failed to create eth client")
	}
	activBlock, txs, err := l2EthCl.InfoAndTxsByNumber(ctx, activationBlockNum)
	if err != nil {
		return fmt.Errorf("failed to get activation block: %w", err)
	}
	if len(txs) < 7 {
		return fmt.Errorf("expected at least 7 txs in Ecotone activation block, but got %d", len(txs))
	}
	for i, tx := range txs {
		if !tx.IsDepositTx() {
			return fmt.Errorf("unexpected non-deposit tx in activation block, index %d, hash %s", i, tx.Hash())
		}
	}
	_, receipts, err := l2EthCl.FetchReceipts(ctx, activBlock.Hash())
	if err != nil {
		return fmt.Errorf("failed to fetch receipts of activation block: %w", err)
	}
	for i, rec := range receipts {
		if rec.Status != types.ReceiptStatusSuccessful {
			return fmt.Errorf("failed tx receipt: %d", i)
		}
		switch i {
		case 1, 2, 6: // 2 implementations + 4788 contract deployment
			if rec.ContractAddress == (common.Address{}) {
				return fmt.Errorf("expected contract deployment, but got none")
			}
		case 3, 4, 5: // proxy upgrades and setEcotone call
			if rec.ContractAddress != (common.Address{}) {
				return fmt.Errorf("unexpected contract deployment")
			}
		}
	}
	env.log.Info("upgrade-txs receipts test: success")
	return nil
}

func checkL1Block(ctx context.Context, env *actionEnv) error {
	cl, err := bindings.NewKromaL1Block(predeploys.KromaL1BlockAddr, env.l2)
	if err != nil {
		return fmt.Errorf("failed to create bindings around L1Block contract: %w", err)
	}
	blobBaseFee, err := cl.BlobBaseFee(nil)
	if err != nil {
		return fmt.Errorf("failed to get blob basfee from L1Block contract: %w", err)
	}
	if big.NewInt(0).Cmp(blobBaseFee) == 0 {
		return errors.New("blob basefee must never be 0, EIP specifies minimum of 1")
	}
	env.log.Info("l1-block-info test: success")
	return nil
}

func checkGPO(ctx context.Context, env *actionEnv) error {
	cl, err := bindings.NewGasPriceOracle(predeploys.GasPriceOracleAddr, env.l2)
	if err != nil {
		return fmt.Errorf("failed to create bindings around L1Block contract: %w", err)
	}
	_, err = cl.Overhead(nil)
	if err == nil || !strings.Contains(err.Error(), "revert") {
		return fmt.Errorf("expected revert on legacy overhead attribute acccess, but got %w", err)
	}
	_, err = cl.Scalar(nil)
	if err == nil || !strings.Contains(err.Error(), "revert") {
		return fmt.Errorf("expected revert on legacy scalar attribute acccess, but got %w", err)
	}
	isEcotone, err := cl.IsEcotone(nil)
	if err != nil {
		return fmt.Errorf("failed to get ecotone status: %w", err)
	}
	if !isEcotone {
		return fmt.Errorf("GPO is not set to ecotone: %w", err)
	}
	blobBaseFeeScalar, err := cl.BlobBaseFeeScalar(nil)
	if err != nil {
		return fmt.Errorf("unable to get blob basefee scalar: %w", err)
	}
	if blobBaseFeeScalar == 0 {
		env.log.Warn("blob basefee scalar is set to 0. SystemConfig needs to emit scalar change to update.")
	}
	env.log.Info("GPO test: success")
	return nil
}

func checkL1Fees(ctx context.Context, env *actionEnv) error {
	rollupCfg, err := env.rollupCl.RollupConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve rollup config: %w", err)
	}
	env.log.Info("making test tx", "addr", env.addr)
	nonce, err := env.l2.PendingNonceAt(ctx, env.addr)
	if err != nil {
		return fmt.Errorf("failed to get pending nonce: %w", err)
	}
	env.log.Info("retrieved account nonce", "nonce", nonce)
	l2RPC := client.NewBaseRPCClient(env.l2.Client())
	l2EthCl, err := sources.NewL2Client(l2RPC, env.log, nil,
		sources.L2ClientDefaultConfig(rollupCfg, false))
	if err != nil {
		return fmt.Errorf("failed to create eth client: %w", err)
	}
	head, err := env.l2.HeaderByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to get head: %w", err)
	}
	gasTip := big.NewInt(2 * params.GWei)
	gasMaxFee := new(big.Int).Add(
		new(big.Int).Mul(big.NewInt(2), head.BaseFee), gasTip)
	to := common.Address{1, 2, 3, 5}
	txData := &types.DynamicFeeTx{
		ChainID:    rollupCfg.L2ChainID,
		Nonce:      nonce,
		GasTipCap:  gasTip,
		GasFeeCap:  gasMaxFee,
		Gas:        params.TxGas + 100, // some margin for the calldata
		To:         &to,
		Value:      big.NewInt(3 * params.GWei),
		Data:       []byte("hello"),
		AccessList: nil,
	}
	tx, err := types.SignNewTx(env.key, types.NewLondonSigner(txData.ChainID), txData)
	if err != nil {
		return fmt.Errorf("failed to sign test tx: %w", err)
	}
	env.log.Info("signed tx", "txhash", tx.Hash())
	if err := env.l2.SendTransaction(ctx, tx); err != nil {
		return fmt.Errorf("failed to send test tx: %w", err)
	}
	receipt, err := retry.Do(ctx, 20, retry.Fixed(time.Second*2), func() (*types.Receipt, error) {
		return env.l2.TransactionReceipt(ctx, tx.Hash())
	})
	if err != nil {
		return fmt.Errorf("failed to confirm tx %s timely: %w", tx.Hash(), err)
	}
	env.RecordGasUsed(receipt)
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("transaction failed, L1 gas used: %d", receipt.L1GasUsed)
	}
	env.log.Info("got receipt", "hash", tx.Hash(), "block", receipt.BlockHash)
	if receipt.FeeScalar != nil {
		return fmt.Errorf("expected fee scalar attribute to be deprecated, but got %v", receipt.FeeScalar)
	}
	payload, err := l2EthCl.PayloadByHash(ctx, receipt.BlockHash)
	if err != nil {
		return fmt.Errorf("failed to get head ref: %w", err)
	}
	headRef, err := derive.PayloadToBlockRef(rollupCfg, payload.ExecutionPayload)
	if err != nil {
		return fmt.Errorf("failed to convert to block-ref: %w", err)
	}
	l1Header, err := retry.Do(ctx, 5, retry.Fixed(time.Second*12), func() (*types.Header, error) {
		env.log.Info("retrieving L1 origin...", "l1", headRef.L1Origin)
		return env.l1.HeaderByHash(ctx, headRef.L1Origin.Hash)
	})
	if err != nil {
		return fmt.Errorf("failed to retrieve L1 origin %s of L2 block %s: %w", headRef.L1Origin, headRef, err)
	}
	if receipt.L1GasPrice.Cmp(l1Header.BaseFee) != 0 {
		return fmt.Errorf("L1 gas price does not include blob fee component: %d != %d", receipt.L1GasPrice, l1Header.BaseFee)
	}
	rawTx, err := tx.MarshalBinary()
	if err != nil {
		return fmt.Errorf("failed to encode tx: %w", err)
	}
	var zero, nonZero uint64
	for _, b := range rawTx {
		if b == 0 {
			zero += 1
		} else {
			nonZero += 1
		}
	}
	expectedCalldataGas := zero*4 + nonZero*16
	env.log.Info("expecting fees", "calldatagas", expectedCalldataGas)
	env.log.Info("paid fees", "l1_fee", receipt.L1Fee, "l1_basefee", receipt.L1GasPrice)
	if new(big.Int).SetUint64(expectedCalldataGas).Cmp(receipt.L1GasUsed) != 0 {
		return fmt.Errorf("expected %d L1 gas, but only spent %d", expectedCalldataGas, receipt.L1GasUsed)
	}
	if big.NewInt(0).Cmp(receipt.L1Fee) >= 0 {
		return fmt.Errorf("calculated to low L1 fee: %d", receipt.L1Fee)
	}
	env.log.Info("L1 fees test: success")
	return nil
}

func checkALL(ctx context.Context, env *actionEnv) error {
	bal, err := env.l2.BalanceAt(ctx, env.addr, nil)
	if err != nil {
		return fmt.Errorf("failed to check balance of account: %w", err)
	}
	env.log.Info("starting checks, tx account", "addr", env.addr, "balance_wei", bal)

	if err := checkAllCancun(ctx, env); err != nil {
		return fmt.Errorf("failed: Cancun error: %w", err)
	}
	if err := checkUpgradeTxs(ctx, env); err != nil {
		return fmt.Errorf("failed: Upgrade-tx error: %w", err)
	}
	if err := checkL1Block(ctx, env); err != nil {
		return fmt.Errorf("failed: L1Block contract error: %w", err)
	}
	if err := checkGPO(ctx, env); err != nil {
		return fmt.Errorf("failed: GPO contract error: %w", err)
	}
	if err := checkL1Fees(ctx, env); err != nil {
		return fmt.Errorf("failed: L1 fees error: %w", err)
	}

	finbal, err := env.l2.BalanceAt(ctx, env.addr, nil)
	if err != nil {
		return fmt.Errorf("failed to check final balance of account: %w", err)
	}
	env.log.Info("completed all tests successfully!",
		"addr", env.addr, "balance_wei", finbal,
		"spent_wei", new(big.Int).Sub(bal, finbal),
		"gas_used_total", env.gasUsed,
		"l1_gas_used_total", env.l1GasUsed,
	)

	return nil
}
