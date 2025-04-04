// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const CrossDomainMessengerStorageLayoutJSON = "{\"storage\":null,\"types\":{}}"

var CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var CrossDomainMessengerDeployedBin = "0x"


func init() {
	if err := json.Unmarshal([]byte(CrossDomainMessengerStorageLayoutJSON), CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["CrossDomainMessenger"] = CrossDomainMessengerStorageLayout
	deployedBytecodes["CrossDomainMessenger"] = CrossDomainMessengerDeployedBin
	immutableReferences["CrossDomainMessenger"] = false
}
