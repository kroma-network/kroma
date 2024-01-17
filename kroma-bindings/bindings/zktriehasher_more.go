// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const ZKTrieHasherStorageLayoutJSON = "{\"storage\":null,\"types\":{}}"

var ZKTrieHasherStorageLayout = new(solc.StorageLayout)

var ZKTrieHasherDeployedBin = "0x6080604052348015600f57600080fd5b506004361060285760003560e01c8063dc8b503814602d575b600080fd5b60537f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f3fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(ZKTrieHasherStorageLayoutJSON), ZKTrieHasherStorageLayout); err != nil {
		panic(err)
	}

	layouts["ZKTrieHasher"] = ZKTrieHasherStorageLayout
	deployedBytecodes["ZKTrieHasher"] = ZKTrieHasherDeployedBin
	immutableReferences["ZKTrieHasher"] = true
}
