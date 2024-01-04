package crossdomain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// HashCrossDomainMessage computes cross domain messaging hashing scheme.
func HashCrossDomainMessageV0(
	nonce *big.Int,
	sender common.Address,
	target common.Address,
	value *big.Int,
	gasLimit *big.Int,
	data []byte,
) (common.Hash, error) {
	encoded, err := EncodeCrossDomainMessageV0(nonce, sender, target, value, gasLimit, data)
	if err != nil {
		return common.Hash{}, err
	}
	hash := crypto.Keccak256(encoded)
	return common.BytesToHash(hash), nil
}
