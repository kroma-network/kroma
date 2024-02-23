package derive

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum-optimism/optimism/op-service/solabi"
	"github.com/kroma-network/kroma/kroma-bindings/predeploys"
)

const (
	MintTokenFuncSignature = "mint()"
	MintTokenLen           = 4
)

var (
	MintTokenFuncBytes4    = crypto.Keccak256([]byte(MintTokenFuncSignature))[:4]
	MintTokenCallerAddress = common.HexToAddress("0xdeaddeaddeaddeaddeaddeaddeaddeaddead0070")
	MintManagerAddress     = predeploys.MintManagerAddr
)

const (
	MintTokenTxGas = 1_000_000
)

type MintToken struct{}

// Binary Format
// +---------+--------------------------+
// | Bytes   | Field                    |
// +---------+--------------------------+
// | 4       | Function signature       |
// +---------+--------------------------+

func (info *MintToken) MarshalBinary() ([]byte, error) {
	w := bytes.NewBuffer(make([]byte, 0, MintTokenLen))
	if err := solabi.WriteSignature(w, MintTokenFuncBytes4); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func (info *MintToken) UnmarshalBinary(data []byte) error {
	if len(data) != MintTokenLen {
		return fmt.Errorf("data is unexpected length: %d", len(data))
	}
	reader := bytes.NewReader(data)

	if _, err := solabi.ReadAndValidateSignature(reader, MintTokenFuncBytes4); err != nil {
		return err
	}
	if !solabi.EmptyReader(reader) {
		return errors.New("too many bytes")
	}
	return nil
}

// NewMintTokenTx creates a mint token transaction.
func NewMintTokenTx(nonce uint64) (*types.MintTokenTx, error) {
	out := &types.MintTokenTx{
		From:  MintTokenCallerAddress,
		To:    &MintManagerAddress,
		Gas:   MintTokenTxGas,
		Nonce: nonce,
		Data:  MintTokenFuncBytes4,
	}
	return out, nil
}

// MintTokenTxBytes returns a serialized mint token transaction.
func MintTokenTxBytes(nextBlockNumber uint64) ([]byte, error) {
	mintTokenTx, err := NewMintTokenTx(nextBlockNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to create mint token tx: %w", err)
	}
	tx := types.NewTx(mintTokenTx)
	opaqueTx, err := tx.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("failed to encode mint token tx: %w", err)
	}
	return opaqueTx, nil
}
