// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/universal/CrossDomainMessenger.sol:CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/universal/CrossDomainMessenger.sol:CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/universal/CrossDomainMessenger.sol:CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1003,\"contract\":\"contracts/universal/CrossDomainMessenger.sol:CrossDomainMessenger\",\"label\":\"_paused\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_bool\"},{\"astId\":1004,\"contract\":\"contracts/universal/CrossDomainMessenger.sol:CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1005,\"contract\":\"contracts/universal/CrossDomainMessenger.sol:CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1006,\"contract\":\"contracts/universal/CrossDomainMessenger.sol:CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_address\"},{\"astId\":1007,\"contract\":\"contracts/universal/CrossDomainMessenger.sol:CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_uint240\"},{\"astId\":1008,\"contract\":\"contracts/universal/CrossDomainMessenger.sol:CrossDomainMessenger\",\"label\":\"failedMessages\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1009,\"contract\":\"contracts/universal/CrossDomainMessenger.sol:CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_array(t_uint256)45_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)45_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[45]\",\"numberOfBytes\":\"1440\",\"base\":\"t_uint256\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

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
