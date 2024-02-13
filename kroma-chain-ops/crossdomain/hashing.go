package crossdomain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// HashCrossDomainMessageV0 is the same as V1 in Kroma.
func HashCrossDomainMessageV0(
	nonce *big.Int,
	sender common.Address,
	target common.Address,
	value *big.Int,
	gasLimit *big.Int,
	data []byte,
) (common.Hash, error) {
	return HashCrossDomainMessageV1(nonce, sender, target, value, gasLimit, data)
}

// HashCrossDomainMessageV1 computes the first post bedrock cross domain
// messaging hashing scheme.
func HashCrossDomainMessageV1(
	nonce *big.Int,
	sender common.Address,
	target common.Address,
	value *big.Int,
	gasLimit *big.Int,
	data []byte,
) (common.Hash, error) {
	encoded, err := EncodeCrossDomainMessageV1(nonce, sender, target, value, gasLimit, data)
	if err != nil {
		return common.Hash{}, err
	}
	hash := crypto.Keccak256(encoded)
	return common.BytesToHash(hash), nil
}
