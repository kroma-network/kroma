package crossdomain_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/utils/chain-ops/crossdomain"
)

// TestEncode tests the encoding of a CrossDomainMessage. The assertion was
// created using solidity.
func TestEncode(t *testing.T) {
	t.Parallel()

	t.Run("V0", func(t *testing.T) {
		expectNonce := common.Big1
		expectVersion := common.Big0

		msg := crossdomain.NewCrossDomainMessage(
			crossdomain.EncodeVersionedNonce(expectNonce, expectVersion),
			common.Address{19: 0x01},
			common.Address{19: 0x02},
			big.NewInt(100),
			big.NewInt(555),
			[]byte{},
		)

		require.Equal(t, uint64(0), msg.Version())

		encoded, err := msg.Encode()
		require.Nil(t, err)

		expect := hexutil.MustDecode("0xd764ad0b0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000064000000000000000000000000000000000000000000000000000000000000022b00000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000000")

		require.Equal(t, expect, encoded)
	})
}

// TestEncode tests the hash of a CrossDomainMessage. The assertion was
// created using solidity.
func TestHash(t *testing.T) {
	t.Parallel()

	t.Run("V0", func(t *testing.T) {
		expectNonce := common.Big0
		expectVersion := common.Big0

		msg := crossdomain.NewCrossDomainMessage(
			crossdomain.EncodeVersionedNonce(expectNonce, expectVersion),
			common.Address{},
			common.Address{19: 0x01},
			big.NewInt(0),
			big.NewInt(5),
			[]byte{},
		)

		require.Equal(t, expectVersion.Uint64(), msg.Version())

		hash, err := msg.Hash()
		require.Nil(t, err)

		expect := common.HexToHash("0xabe2ab138bea877c082a26a761f9c999ef57748d2d0ab05a24c6e8bdd1c5fb41")
		require.Equal(t, expect, hash)
	})
}
