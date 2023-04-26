package eth

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/poseidon"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
	zktrie "github.com/kroma-network/zktrie/trie"
	zkt "github.com/kroma-network/zktrie/types"
)

func init() {
	zkt.InitHashScheme(poseidon.HashFixed)
}

type StorageProofEntry struct {
	Key   common.Hash     `json:"key"`
	Value hexutil.Big     `json:"value"`
	Proof []hexutil.Bytes `json:"proof"`
}

type AccountResult struct {
	AccountProof []hexutil.Bytes `json:"accountProof"`

	Address     common.Address `json:"address"`
	Balance     *hexutil.Big   `json:"balance"`
	CodeHash    common.Hash    `json:"codeHash"`
	Nonce       hexutil.Uint64 `json:"nonce"`
	StorageHash common.Hash    `json:"storageHash"`

	// Optional
	StorageProof []StorageProofEntry `json:"storageProof,omitempty"`
}

type ValueValidatorFunc func(val []byte, isZktrie bool) error

func verifyProof(proof []hexutil.Bytes, key []byte, rootHash common.Hash, validateValue ValueValidatorFunc) error {
	isZktrie := false

	if len(proof) > 0 {
		lastElem := proof[len(proof)-1]
		if bytes.Equal(zktrie.ProofMagicBytes(), lastElem) {
			isZktrie = true
		}
	}

	// load all nodes into a DB
	db := memorydb.New()

	var val []byte
	var err error

	if isZktrie {
		for i, encodedNode := range proof {
			// NOTE(chokobole): Do not check account proof for magicSMTBytes.
			if i == len(proof)-1 {
				break
			}
			node, err := zktrie.NewNodeFromBytes(encodedNode)
			if err != nil {
				return fmt.Errorf("failed to create node from encoded %s: %w", hex.EncodeToString(encodedNode), err)
			}
			nodeHash, err := node.NodeHash()
			if err != nil {
				return fmt.Errorf("failed to get node hash from node %d: %w", i, err)
			}
			if err := db.Put(nodeHash[:], node.Value()); err != nil {
				return fmt.Errorf("failed to load proof value %d into mem db: %w", i, err)
			}
		}

		val, err = trie.VerifyProofSMT(rootHash, key, db)
	} else {
		for i, encodedNode := range proof {
			nodeKey := encodedNode
			if len(encodedNode) >= 32 { // small MPT nodes are not hashed
				nodeKey = crypto.Keccak256(encodedNode)
			}
			if err := db.Put(nodeKey, encodedNode); err != nil {
				return fmt.Errorf("failed to load proof node %d into mem db: %w", i, err)
			}
		}

		path := crypto.Keccak256(key)
		val, err = trie.VerifyProof(rootHash, path, db)
	}
	if err != nil {
		return fmt.Errorf("key: %s, trie: %s error: %w", key, rootHash, err)
	}

	return validateValue(val, isZktrie)
}

func (res *AccountResult) getAccountClaimedValue(isZktrie bool) ([]byte, error) {
	if isZktrie {
		account := types.StateAccount{
			Nonce:    uint64(res.Nonce),
			Balance:  res.Balance.ToInt(),
			Root:     res.StorageHash,
			CodeHash: res.CodeHash.Bytes(),
		}
		accountClaimedValueArray, l := account.MarshalFields()
		ret := make([]byte, 0)
		for i := uint32(0); i < l; i++ {
			ret = append(ret, accountClaimedValueArray[i][:]...)
		}
		return ret, nil
	} else {
		var err error
		accountClaimed := []any{uint64(res.Nonce), (*big.Int)(res.Balance).Bytes(), res.StorageHash, res.CodeHash}
		ret, err := rlp.EncodeToBytes(accountClaimed)
		if err != nil {
			return nil, fmt.Errorf("failed to encode account from retrieved values: %w", err)
		}
		return ret, nil
	}
}

// Verify an account (and optionally storage) proof from the getProof RPC. See https://eips.ethereum.org/EIPS/eip-1186
func (res *AccountResult) Verify(stateRoot common.Hash) error {
	// verify storage proof values, if any, against the storage trie root hash of the account
	for i, entry := range res.StorageProof {
		validator := func(val []byte, isZktrie bool) error {
			_, expected, _, err := rlp.Split(val)
			if err != nil {
				return err
			}
			if !bytes.Equal(expected, entry.Value.ToInt().Bytes()) {
				return fmt.Errorf("value %d in storage proof does not match proven value at key %s", i, entry.Key)
			}
			return nil
		}

		err := verifyProof(entry.Proof, entry.Key[:], res.StorageHash, validator)
		if err != nil {
			return fmt.Errorf("failed to verify storage proof %d: %w", i, err)
		}
	}

	// now get the full value from the account proof, and check that it matches
	validator := func(val []byte, isZktrie bool) error {
		expected, err := res.getAccountClaimedValue(isZktrie)
		if err != nil {
			return err
		}
		if !bytes.Equal(expected, val) {
			return fmt.Errorf("L1 RPC is tricking us, account proof does not match provided deserialized values:\n"+
				"  claimed: %x\n"+
				"  proof:   %x", expected, val)
		}
		return nil
	}

	err := verifyProof(res.AccountProof, res.Address[:], stateRoot, validator)
	if err != nil {
		return fmt.Errorf("failed to verify account proof %w", err)
	}

	return err
}
