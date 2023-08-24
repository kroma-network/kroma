package solabi_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/kroma-network/kroma/utils/service/solabi"
)

func TestEmptyReader(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		r := new(bytes.Buffer)
		require.True(t, solabi.EmptyReader(r))
	})
	t.Run("empty after read", func(t *testing.T) {
		r := bytes.NewBufferString("not empty")
		tmp := make([]byte, 9)
		n, err := r.Read(tmp)
		require.Equal(t, 9, n)
		require.NoError(t, err)
		require.True(t, solabi.EmptyReader(r))
	})
	t.Run("extra bytes", func(t *testing.T) {
		r := bytes.NewBufferString("not empty")
		require.False(t, solabi.EmptyReader(r))
	})
}
