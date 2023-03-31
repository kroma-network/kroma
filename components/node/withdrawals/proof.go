package withdrawals

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/poseidon"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/trie"

	zktrie "github.com/wemixkanvas/zktrie/trie"
	zkt "github.com/wemixkanvas/zktrie/types"
)

var magicHash = []byte("THIS IS THE MAGIC INDEX FOR ZKTRIE")

func init() {
	zkt.InitHashScheme(poseidon.HashFixed)
}

type proofDB struct {
	m map[string][]byte
}

func (p *proofDB) Has(key []byte) (bool, error) {
	_, ok := p.m[string(key)]
	return ok, nil
}

func (p *proofDB) Get(key []byte) ([]byte, error) {
	v, ok := p.m[string(key)]
	if !ok {
		return nil, errors.New("not found")
	}
	return v, nil
}

func GenerateProofDB(proof []string) (*proofDB, error) {
	p := proofDB{m: make(map[string][]byte)}
	p.m[string(magicHash)] = zktrie.ProofMagicBytes()

	for i, s := range proof {
		if i == len(proof)-1 {
			break
		}

		value := common.FromHex(s)
		node, err := zktrie.NewNodeFromBytes(value)
		if err != nil {
			return nil, err
		}
		h, err := node.NodeHash()
		if err != nil {
			return nil, err
		}

		p.m[string(h[:])] = node.Value()
	}
	return &p, nil
}

func VerifyAccountProof(root common.Hash, address common.Address, account types.StateAccount, proof []string) error {
	expected, flag := account.MarshalFields()

	db, err := GenerateProofDB(proof)
	if err != nil {
		return fmt.Errorf("failed to generate proof db: %w", err)
	}
	value, err := trie.VerifyProof(root, address[:], db)
	if err != nil {
		return fmt.Errorf("failed to verify proof: %w", err)
	}

	expectedValue := make([]byte, 0)
	for i := uint32(0); i < flag; i++ {
		expectedValue = append(expectedValue, expected[i][:]...)
	}

	if bytes.Equal(value, expectedValue) {
		return nil
	} else {
		return errors.New("proved value is not the same as the expected value")
	}
}

func VerifyStorageProof(root common.Hash, proof gethclient.StorageResult) error {
	db, err := GenerateProofDB(proof.Proof)
	if err != nil {
		return fmt.Errorf("failed to generate proof db: %w", err)
	}
	value, err := trie.VerifyProof(root, common.FromHex(proof.Key), db)
	if err != nil {
		return fmt.Errorf("failed to verify proof: %w", err)
	}

	expected := proof.Value.Bytes()
	if bytes.Equal(value, zkt.NewByte32FromBytes(expected)[:]) {
		return nil
	} else {
		return errors.New("proved value is not the same as the expected value")
	}
}

func VerifyProof(stateRoot common.Hash, proof *gethclient.AccountResult) error {
	err := VerifyAccountProof(
		stateRoot,
		proof.Address,
		types.StateAccount{
			Nonce:    proof.Nonce,
			Balance:  proof.Balance,
			Root:     proof.StorageHash,
			CodeHash: proof.CodeHash[:],
		},
		proof.AccountProof,
	)
	if err != nil {
		return fmt.Errorf("failed to validate account: %w", err)
	}
	for i, storageProof := range proof.StorageProof {
		err = VerifyStorageProof(proof.StorageHash, storageProof)
		if err != nil {
			return fmt.Errorf("failed to validate storage proof %d: %w", i, err)
		}
	}
	return nil
}
