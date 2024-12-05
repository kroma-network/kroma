package main

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	kromabindings "github.com/kroma-network/kroma/kroma-bindings/bindings"
	"github.com/kroma-network/kroma/kroma-chain-ops/crossdomain"
)

var UnknownNonceVersion = errors.New("Unknown nonce version")

// checkOk checks if ok is false, and panics if so.
// Shorthand to ease go's god awful error handling
func checkOk(ok bool) {
	if !ok {
		panic(fmt.Errorf("checkOk failed"))
	}
}

// checkErr checks if err is not nil, and throws if so.
// Shorthand to ease go's god awful error handling
func checkErr(err error, failReason string) {
	if err != nil {
		panic(fmt.Errorf("%s: %w", failReason, err))
	}
}

// encodeCrossDomainMessage encodes a versioned cross domain message into a byte array.
func encodeCrossDomainMessage(nonce *big.Int, sender common.Address, target common.Address, value *big.Int, gasLimit *big.Int, data []byte) ([]byte, error) {
	_, version := crossdomain.DecodeVersionedNonce(nonce)

	var encoded []byte
	var err error
	if version.Cmp(big.NewInt(0)) == 0 {
		// Encode cross domain message V0
		// [Kroma: START]
		encoded, err = crossdomain.EncodeCrossDomainMessageV0(nonce, sender, target, value, gasLimit, data)
		// [Kroma: END]
	} else if version.Cmp(big.NewInt(1)) == 0 {
		// Encode cross domain message V1
		encoded, err = crossdomain.EncodeCrossDomainMessageV1(nonce, sender, target, value, gasLimit, data)
	} else {
		return nil, UnknownNonceVersion
	}

	return encoded, err
}

// hashWithdrawal hashes a withdrawal transaction.
func hashWithdrawal(nonce *big.Int, sender common.Address, target common.Address, value *big.Int, gasLimit *big.Int, data []byte) (common.Hash, error) {
	wd := crossdomain.Withdrawal{
		Nonce:    nonce,
		Sender:   &sender,
		Target:   &target,
		Value:    value,
		GasLimit: gasLimit,
		Data:     data,
	}
	return wd.Hash()
}

// [Kroma: START]
// hashOutputRootProof hashes an output root proof.
func hashOutputRootProof(version common.Hash, stateRoot common.Hash, messagePasserStorageRoot common.Hash, latestBlockHash common.Hash, nextBlockHash common.Hash) (common.Hash, error) {
	var hash eth.Bytes32
	var err error
	if nextBlockHash == (common.Hash{}) {
		hash, err = rollup.ComputeL2OutputRoot(&bindings.TypesOutputRootProof{
			Version:                  version,
			StateRoot:                stateRoot,
			MessagePasserStorageRoot: messagePasserStorageRoot,
			LatestBlockhash:          latestBlockHash,
		})
	} else {
		hash, err = rollup.ComputeKromaL2OutputRoot(&kromabindings.TypesOutputRootProof{
			Version:                  version,
			StateRoot:                stateRoot,
			MessagePasserStorageRoot: messagePasserStorageRoot,
			LatestBlockhash:          latestBlockHash,
			NextBlockHash:            nextBlockHash,
		})
	}
	if err != nil {
		return common.Hash{}, err
	}
	return common.Hash(hash), nil
}

// [Kroma: END]

// makeDepositTx creates a deposit transaction type.
func makeDepositTx(
	from common.Address,
	to common.Address,
	value *big.Int,
	mint *big.Int,
	gasLimit *big.Int,
	isCreate bool,
	data []byte,
	l1BlockHash common.Hash,
	logIndex *big.Int,
) types.DepositTx {
	// Create deposit transaction source
	udp := derive.UserDepositSource{
		L1BlockHash: l1BlockHash,
		LogIndex:    logIndex.Uint64(),
	}

	// Create deposit transaction
	depositTx := types.DepositTx{
		SourceHash:          udp.SourceHash(),
		From:                from,
		Value:               value,
		Gas:                 gasLimit.Uint64(),
		IsSystemTransaction: false, // This will never be a system transaction in the tests.
		Data:                data,
	}

	// Fill optional fields
	if mint.Cmp(big.NewInt(0)) == 1 {
		depositTx.Mint = mint
	}
	if !isCreate {
		depositTx.To = &to
	}

	return depositTx
}

func toKromaDepositTx(depTx types.DepositTx) types.KromaDepositTx {
	kromaDepTx := types.KromaDepositTx{
		SourceHash: depTx.SourceHash,
		From:       depTx.From,
		Value:      new(big.Int).Set(depTx.Value),
		Gas:        depTx.Gas,
		Data:       depTx.Data,
	}
	// Fill optional fields
	if depTx.Mint != nil {
		kromaDepTx.Mint = new(big.Int).Set(depTx.Mint)
	}
	if depTx.To != nil {
		to := *depTx.To
		kromaDepTx.To = &to
	}
	return kromaDepTx
}

// Custom type to write the generated proof to
type proofList [][]byte

func (n *proofList) Put(key []byte, value []byte) error {
	*n = append(*n, value)
	return nil
}

func (n *proofList) Delete(key []byte) error {
	panic("not supported")
}
