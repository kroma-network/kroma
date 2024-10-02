package withdrawals

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/poseidon"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
	zktrie "github.com/kroma-network/zktrie/trie"
	zkt "github.com/kroma-network/zktrie/types"
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

func GenerateProofDB(proof []string, isKromaMPT bool) (*proofDB, error) {
	p := proofDB{m: make(map[string][]byte)}

	if isKromaMPT {
		for _, s := range proof {
			value := common.FromHex(s)
			key := crypto.Keccak256(value)
			p.m[string(key)] = value
		}
	} else {
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
	}

	return &p, nil
}

func VerifyAccountProof(root common.Hash, address common.Address, account types.StateAccount, proof []string, isKromaMPT bool) error {
	secureKey := address[:]
	var expected []byte
	if isKromaMPT {
		var err error
		expected, err = rlp.EncodeToBytes(&account)
		if err != nil {
			return fmt.Errorf("failed to encode rlp: %w", err)
		}
		secureKey = crypto.Keccak256(secureKey)
	} else {
		expectedValue, flag := account.MarshalFields()
		for i := uint32(0); i < flag; i++ {
			expected = append(expected, expectedValue[i][:]...)
		}
	}

	db, err := GenerateProofDB(proof, isKromaMPT)
	if err != nil {
		return fmt.Errorf("failed to generate proof db: %w", err)
	}
	value, err := trie.VerifyProof(root, secureKey, db)
	if err != nil {
		return fmt.Errorf("failed to verify proof: %w", err)
	}

	if bytes.Equal(value, expected) {
		return nil
	} else {
		return errors.New("proved value is not the same as the expected value")
	}
}

func VerifyStorageProof(root common.Hash, proof gethclient.StorageResult, isKromaMPT bool) error {
	secureKey := common.FromHex(proof.Key)
	expected := proof.Value.Bytes()
	if isKromaMPT {
		secureKey = crypto.Keccak256(secureKey)
	} else {
		expected = zkt.NewByte32FromBytes(expected)[:]
	}

	db, err := GenerateProofDB(proof.Proof, isKromaMPT)
	if err != nil {
		return fmt.Errorf("failed to generate proof db: %w", err)
	}
	value, err := trie.VerifyProof(root, secureKey, db)
	if err != nil {
		return fmt.Errorf("failed to verify proof: %w", err)
	}

	if bytes.Equal(value, expected) {
		return nil
	} else {
		return errors.New("proved value is not the same as the expected value")
	}
}

func VerifyProof(stateRoot common.Hash, proof *gethclient.AccountResult, isKromaMPT bool) error {
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
		isKromaMPT,
	)
	if err != nil {
		return fmt.Errorf("failed to validate account: %w", err)
	}
	for i, storageProof := range proof.StorageProof {
		err = VerifyStorageProof(proof.StorageHash, storageProof, isKromaMPT)
		if err != nil {
			return fmt.Errorf("failed to validate storage proof %d: %w", i, err)
		}
	}
	return nil
}
