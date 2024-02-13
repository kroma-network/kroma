package crossdomain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	opxdm "github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain"
)

// EncodeCrossDomainMessageV0 is the same as V1 in Kroma.
func EncodeCrossDomainMessageV0(
	nonce *big.Int,
	sender common.Address,
	target common.Address,
	value *big.Int,
	gasLimit *big.Int,
	data []byte,
) ([]byte, error) {
	return opxdm.EncodeCrossDomainMessageV1(nonce, sender, target, value, gasLimit, data)
}

func EncodeCrossDomainMessageV1(
	nonce *big.Int,
	sender common.Address,
	target common.Address,
	value *big.Int,
	gasLimit *big.Int,
	data []byte,
) ([]byte, error) {
	return opxdm.EncodeCrossDomainMessageV1(nonce, sender, target, value, gasLimit, data)
}

// DecodeVersionedNonce will decode the version that is encoded in the nonce
func DecodeVersionedNonce(versioned *big.Int) (*big.Int, *big.Int) {
	return opxdm.DecodeVersionedNonce(versioned)
}

// EncodeVersionedNonce will encode the version into the nonce
func EncodeVersionedNonce(nonce, version *big.Int) *big.Int {
	return opxdm.EncodeVersionedNonce(nonce, version)
}
