package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

var (
	logger *log.Logger
)

type StateVerifier struct {
	rawDB      ethdb.Database
	rawTransDB ethdb.Database
	mptDB      *trie.Database
	transDB    *trie.Database
}

func NewStateVerifier(mptDBOpenOption, transitionedDBOpenOption rawdb.OpenOptions) *StateVerifier {
	mptDB, err := rawdb.Open(mptDBOpenOption)
	transitionedDB, err := rawdb.Open(transitionedDBOpenOption)

	if err != nil {
		logger.Panic(fmt.Errorf("failed to db open: %w", err))
	}

	return &StateVerifier{
		rawDB:      mptDB,
		rawTransDB: transitionedDB,
		mptDB:      trie.NewDatabase(mptDB, nil),
		transDB:    trie.NewDatabase(transitionedDB, nil),
	}

}

func (v *StateVerifier) verify() (string, bool) {
	mptHeadBlock := rawdb.ReadHeadHeader(v.rawDB)
	mptHeadBlockNumber := mptHeadBlock.Number.Uint64()
	transHeadBlock := rawdb.ReadHeadHeader(v.rawTransDB)
	transHeadBlockNumber := transHeadBlock.Number.Uint64()
	var commonLatestBlockNumber uint64
	if mptHeadBlockNumber > transHeadBlockNumber {
		commonLatestBlockNumber = transHeadBlockNumber
	} else {
		commonLatestBlockNumber = mptHeadBlockNumber
	}

	var currentAccKey string

	fmt.Println("commonLatestBlockNumber :", commonLatestBlockNumber)

	mptBlockHash := rawdb.ReadCanonicalHash(v.rawDB, commonLatestBlockNumber)
	mptStateRoot := rawdb.ReadBlock(v.rawDB, mptBlockHash, commonLatestBlockNumber).Root()

	transBlockHash := rawdb.ReadCanonicalHash(v.rawTransDB, commonLatestBlockNumber)
	transStateRoot := rawdb.ReadBlock(v.rawTransDB, transBlockHash, commonLatestBlockNumber).Root()

	// Note: comment out because we want to check what is wrong
	//if mptStateRoot.Cmp(transStateRoot) != 0 {
	//	logger.Println("State Root is not equal")
	//	return false
	//}

	mptStateTrie, err := trie.NewStateTrie(trie.StateTrieID(mptStateRoot), v.mptDB)

	if err != nil {
		logger.Panicln("Failed to open mpt state trie", "root", mptStateRoot, "err", err)
	}

	transStateTrie, err := trie.NewStateTrie(trie.StateTrieID(transStateRoot), v.transDB)

	if err != nil {
		logger.Panicln("Failed to open trans state trie", "root", transStateRoot, "err", err)
	}

	acctIt, err := mptStateTrie.NodeIterator(nil)
	if err != nil {
		logger.Panicln(fmt.Errorf("Failed to open iterator: %w", err))
	}

	accIter := trie.NewIterator(acctIt)

	var (
		accounts   int
		slots      int
		codes      int
		lastReport time.Time
		start      = time.Now()
	)

	for accIter.Next() {
		accounts += 1
		currentAccKey = hex.EncodeToString(accIter.Key)
		var acc types.StateAccount
		if err := rlp.DecodeBytes(accIter.Value, &acc); err != nil {
			logger.Panicln(fmt.Errorf("Invalid account encountered during traversal: %w", err))
		}

		if err != nil {
			logger.Panicln(fmt.Errorf("decode erorr %w", err))
		}

		transAcc, err := transStateTrie.GetAccountByHash(common.BytesToHash(accIter.Key))
		//fmt.Println("hex.EncodeToString(accIter.Key) : ", hex.EncodeToString(accIter.Key))

		if err != nil {
			logger.Panicln(fmt.Errorf("Failed to get account in TRANS: %w", err))
		}

		if transAcc.Balance.Cmp(acc.Balance) != 0 {
			logger.Printf("balance mismatch. expected %s, got %s\n", transAcc.Balance, acc.Balance)
			return currentAccKey, false
		}

		if transAcc.Nonce != acc.Nonce {
			logger.Printf("nonce mismatch. expected %d, got %d\n", transAcc.Nonce, acc.Nonce)
			return currentAccKey, false
		}

		if !bytes.Equal(transAcc.CodeHash, acc.CodeHash) {
			logger.Printf("CodeHash mismatch. expected %s, got %s\n", common.BytesToHash(transAcc.CodeHash), common.BytesToHash(acc.CodeHash))
			return currentAccKey, false
		}

		if acc.Root != v.mptDB.EmptyRoot() {
			id := trie.StorageTrieID(mptStateRoot, common.BytesToHash(accIter.Key), acc.Root)
			anotherId := trie.StorageTrieID(transStateRoot, common.BytesToHash(accIter.Key), acc.Root)

			storageTrie, err := trie.NewStateTrie(id, v.mptDB)
			transLowLevelStorageTrie, err := trie.New(anotherId, v.transDB)

			if err != nil {
				logger.Panicln("Failed to open low-level trans storage trie", "root", transLowLevelStorageTrie, "err", err)
			}

			if err != nil {
				logger.Panicln(fmt.Errorf("Failed to open storage trie: %w", err))
			}
			storageIt, err := storageTrie.NodeIterator(nil)
			if err != nil {
				logger.Panicln(fmt.Errorf("Failed to open storage iterator: %w", err))
			}
			storageIter := trie.NewIterator(storageIt)
			for storageIter.Next() {
				slots += 1

				transVal, err := transLowLevelStorageTrie.Get(storageIter.Key)

				if err != nil {
					logger.Printf("failed find value for %s\n", common.BytesToHash(storageIter.Key).String())
					return currentAccKey, false
				}

				if err != nil {
					logger.Printf("failed to decode storage value for %s\n", common.BytesToHash(storageIter.Key).String())
					return currentAccKey, false
				}

				if !bytes.Equal(storageIter.Value, transVal) {
					logger.Printf("not equal storage value - mpt val : %s  VS  trans val : %s\n", common.Bytes2Hex(storageIter.Value), common.Bytes2Hex(transVal))
					return currentAccKey, false
				}

				if time.Since(lastReport) > time.Second*8 {
					logger.Println("Traversing state", "accounts", accounts, "slots", slots, "codes", codes, "elapsed", common.PrettyDuration(time.Since(start)))
					lastReport = time.Now()
				}
			}
			if storageIter.Err != nil {
				logger.Panicln(fmt.Errorf("Failed to traverse storage trie: %w", err))
			}
		}
		if !bytes.Equal(acc.CodeHash, types.EmptyCodeHash.Bytes()) {
			if !rawdb.HasCode(v.rawDB, common.BytesToHash(acc.CodeHash)) {
				logger.Panicln(fmt.Errorf("Code is missing: %w", err))
			}
			codes += 1
		}

		if time.Since(lastReport) > time.Second*8 {
			logger.Println("Traversing state", "accounts", accounts, "slots", slots, "codes", codes, "elapsed", common.PrettyDuration(time.Since(start)))
			lastReport = time.Now()
		}
	}
	if accIter.Err != nil {
		logger.Panicln(fmt.Errorf("Failed to traverse storage trie: %w", err))
	}
	logger.Println("State is complete", "accounts", accounts, "slots", slots, "codes", codes, "elapsed", common.PrettyDuration(time.Since(start)))
	logger.Println("accounts number:", accounts)

	if mptStateRoot.Cmp(transStateRoot) != 0 {
		logger.Println("State Root is not equal")
		return "", false
	}

	return currentAccKey, true
}

func main() {
	logger = log.New(os.Stdout, "INFO: ", log.LstdFlags)

	mptRawDBOpenOption := rawdb.OpenOptions{
		Type:              "pebble",
		Directory:         "/db/geth/chaindata",
		AncientsDirectory: "/db/geth/chaindata/ancient",
		Namespace:         "eth/db/chaindata/",
		Cache:             0,
		Handles:           1,
		ReadOnly:          true,
	}

	transRawDBOpenOption := rawdb.OpenOptions{
		Type:              "pebble",
		Directory:         "/transitioned-db/geth/chaindata",
		AncientsDirectory: "/transitioned-db/geth/chaindata/ancient",
		Namespace:         "eth/db/chaindata/",
		Cache:             0,
		Handles:           1,
		ReadOnly:          true,
	}

	stateVerifier := NewStateVerifier(mptRawDBOpenOption, transRawDBOpenOption)

	_ = stateVerifier

	currentAccKey, result := stateVerifier.verify()
	if !result {
		logger.Println("currentAccKey : ", currentAccKey)
	}

	fmt.Println("result : ", result)

}
