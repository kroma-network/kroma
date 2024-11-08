package eth

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestOutputV0Codec(t *testing.T) {
	output := OutputV0{
		StateRoot:                Bytes32{1, 2, 3},
		MessagePasserStorageRoot: Bytes32{4, 5, 6},
		BlockHash:                common.Hash{7, 8, 9},
	}
	marshaled := output.Marshal()
	unmarshaled, err := UnmarshalOutput(marshaled)
	require.NoError(t, err)
	unmarshaledV0 := unmarshaled.(*OutputV0)
	require.Equal(t, output, *unmarshaledV0)

	_, err = UnmarshalOutput([]byte{0: 0xA, 32: 0xA})
	require.ErrorIs(t, err, ErrInvalidOutputVersion)
	_, err = UnmarshalOutput([]byte{64: 0xA})
	require.ErrorIs(t, err, ErrInvalidOutput)
}

// [Kroma: START]
func TestKromaOutputV0Codec(t *testing.T) {
	output := KromaOutputV0{
		OutputV0: OutputV0{
			StateRoot:                Bytes32{1, 2, 3},
			MessagePasserStorageRoot: Bytes32{4, 5, 6},
			BlockHash:                common.Hash{7, 8, 9},
		},
		NextBlockHash: common.Hash{10, 11, 12},
	}
	marshaled := output.Marshal()
	unmarshaled, err := UnmarshalOutput(marshaled)
	require.NoError(t, err)
	unmarshaledV0 := unmarshaled.(*KromaOutputV0)
	require.Equal(t, output, *unmarshaledV0)

	_, err = UnmarshalOutput([]byte{0: 0xA, 32: 0xA})
	require.ErrorIs(t, err, ErrInvalidOutputVersion)
	_, err = UnmarshalOutput([]byte{128: 0xA})
	require.ErrorIs(t, err, ErrInvalidOutput)
}

// [Kroma: END]
