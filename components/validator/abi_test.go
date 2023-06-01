package validator

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/bindings/bindings"
	"github.com/kroma-network/kroma/components/node/testutils"
)

// setupL2OutputOracle deploys the L2OutputOracle contract to a simulated backend
func setupL2OutputOracle() (common.Address, *bind.TransactOpts, *backends.SimulatedBackend, *bindings.L2OutputOracle, error) {
	privateKey, err := crypto.GenerateKey()
	from := crypto.PubkeyToAddress(privateKey.PublicKey)
	if err != nil {
		return common.Address{}, nil, nil, nil, err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		return common.Address{}, nil, nil, nil, err
	}
	backend := backends.NewSimulatedBackend(core.GenesisAlloc{from: {Balance: big.NewInt(params.Ether)}}, 50_000_000)
	_, _, contract, err := bindings.DeployL2OutputOracle(
		opts,
		backend,
		big.NewInt(10),
		big.NewInt(2),
		big.NewInt(0),
		big.NewInt(0),
		from,
		common.Address{0xdd},
		big.NewInt(100))
	if err != nil {
		return common.Address{}, nil, nil, nil, err
	}
	return from, opts, backend, contract, nil
}

// TestManualABIPacking ensure that the manual ABI packing is the same as going through the bound contract.
// We don't use the contract to transact because it does not fit our transaction management scheme, but
// we want to make sure that we don't incorrectly create the transaction data.
func TestManualABIPacking(t *testing.T) {
	_, opts, _, contract, err := setupL2OutputOracle()
	require.NoError(t, err)

	abi, err := bindings.L2OutputOracleMetaData.GetAbi()
	require.NoError(t, err)

	rng := rand.New(rand.NewSource(1234))
	output := testutils.RandomOutputResponse(rng)
	bondAmount := uint64(1)

	txData, err := submitL2OutputTxData(abi, output, bondAmount)
	require.NoError(t, err)

	// set a gas limit to disable gas estimation. The invariants that the L2OO tries to uphold
	// are not maintained in this test.
	opts.GasLimit = 100_000
	tx, err := contract.SubmitL2Output(
		opts,
		output.OutputRoot,
		new(big.Int).SetUint64(output.BlockRef.Number),
		output.Status.CurrentL1.Hash,
		new(big.Int).SetUint64(output.Status.CurrentL1.Number),
		new(big.Int).SetUint64(bondAmount))
	require.NoError(t, err)

	require.Equal(t, txData, tx.Data())
}
