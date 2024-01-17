package crossdomain_test

import (
	"math/big"
	"testing"

	opxdm "github.com/ethereum-optimism/optimism/op-chain-ops/crossdomain"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/kroma-chain-ops/crossdomain"
)

func FuzzVersionedNonce(f *testing.F) {
	f.Fuzz(func(t *testing.T, _nonce []byte, _version uint16) {
		inputNonce := new(big.Int).SetBytes(_nonce)

		// Clamp nonce to uint240
		if inputNonce.Cmp(opxdm.NonceMask) > 0 {
			inputNonce = new(big.Int).Set(opxdm.NonceMask)
		}
		// Clamp version to 0 or 1
		_version = _version % 2

		inputVersion := new(big.Int).SetUint64(uint64(_version))
		encodedNonce := crossdomain.EncodeVersionedNonce(inputNonce, inputVersion)

		decodedNonce, decodedVersion := crossdomain.DecodeVersionedNonce(encodedNonce)

		require.Equal(t, decodedNonce.Uint64(), inputNonce.Uint64())
		require.Equal(t, decodedVersion.Uint64(), inputVersion.Uint64())
	})
}
