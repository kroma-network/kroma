// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const StandardBridgeStorageLayoutJSON = "{\"storage\":null,\"types\":{}}"

var StandardBridgeStorageLayout = new(solc.StorageLayout)

var StandardBridgeDeployedBin = "0x"


func init() {
	if err := json.Unmarshal([]byte(StandardBridgeStorageLayoutJSON), StandardBridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["StandardBridge"] = StandardBridgeStorageLayout
	deployedBytecodes["StandardBridge"] = StandardBridgeDeployedBin
	immutableReferences["StandardBridge"] = false
}
