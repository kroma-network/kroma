// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const MintManagerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/governance/MintManager.sol:MintManager\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/governance/MintManager.sol:MintManager\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/governance/MintManager.sol:MintManager\",\"label\":\"recipients\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_address)dyn_storage\"},{\"astId\":1003,\"contract\":\"contracts/governance/MintManager.sol:MintManager\",\"label\":\"shareOf\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1004,\"contract\":\"contracts/governance/MintManager.sol:MintManager\",\"label\":\"lastMintedBlock\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_uint256\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_address)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"address[]\",\"numberOfBytes\":\"32\",\"base\":\"t_address\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var MintManagerStorageLayout = new(solc.StorageLayout)

var MintManagerDeployedBin = "0x608060405234801561001057600080fd5b50600436106100df5760003560e01c806354fd4d501161008c5780637fbbe46f116100665780637fbbe46f1461022d578063894bee6114610223578063a103a2dd14610240578063e2da80901461025b57600080fd5b806354fd4d50146101b35780637b6a4cda146101fc5780637eb118451461022357600080fd5b80632efd46d6116100bd5780632efd46d61461013457806349ed1e6914610180578063544a54e3146101a757600080fd5b8063062459e6146100e45780631249c58b1461010a57806321e5e2c414610114575b600080fd5b6100f76100f2366004610c94565b610282565b6040519081526020015b60405180910390f35b61011261032e565b005b6100f7610122366004610cad565b60026020526000908152604090205481565b61015b7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610101565b6100f77f000000000000000000000000000000000000000000000000000000000000000081565b6100f76402540be40081565b6101ef6040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101019190610cea565b6100f77f000000000000000000000000000000000000000000000000000000000000000081565b6100f7620186a081565b61011261023b366004610da9565b61065c565b61015b73deaddeaddeaddeaddeaddeaddeaddeaddead007081565b6100f77f000000000000000000000000000000000000000000000000000000000000000081565b60007f0000000000000000000000000000000000000000000000000000000000000000816102af84610ae3565b50905060015b8181101561032557620186a06102eb7f000000000000000000000000000000000000000000000000000000000000000085610e44565b6102f59190610eb0565b92506402540be4006103078185610eb0565b6103119190610e44565b92508061031d81610ec4565b9150506102b5565b50909392505050565b3373deaddeaddeaddeaddeaddeaddeaddeaddead0070146103d6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f4d696e744d616e616765723a206f6e6c7920746865206d696e742063616c6c6560448201527f722069732061636365707465640000000000000000000000000000000000000060648201526084015b60405180910390fd5b4360035403610467576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f4d696e744d616e616765723a20746f6b656e73206861766520616c726561647960448201527f206265656e206d696e74656420696e207468697320626c6f636b00000000000060648201526084016103cd565b6000807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104f99190610efc565b111561050f5761050843610282565b905061051b565b61051843610b8f565b90505b80156106595760005b6001548110156106535760006001828154811061054357610543610f15565b600091825260208083209091015473ffffffffffffffffffffffffffffffffffffffff168083526002909152604082205490925090620186a06105868387610e44565b6105909190610eb0565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018390529192507f0000000000000000000000000000000000000000000000000000000000000000909116906340c10f1990604401600060405180830381600087803b15801561062557600080fd5b505af1158015610639573d6000803e3d6000fd5b50505050505050808061064b90610ec4565b915050610524565b50436003555b50565b600054610100900460ff161580801561067c5750600054600160ff909116105b806106965750303b158015610696575060005460ff166001145b610722576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016103cd565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561078057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b83821461080e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f4d696e744d616e616765723a20696e76616c6964206c656e677468206f66206160448201527f727261790000000000000000000000000000000000000000000000000000000060648201526084016103cd565b6000805b85811015610a0b57600087878381811061082e5761082e610f15565b90506020020160208101906108439190610cad565b905073ffffffffffffffffffffffffffffffffffffffff81166108e8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4d696e744d616e616765723a20726563697069656e742061646472657373206360448201527f616e6e6f7420626520300000000000000000000000000000000000000000000060648201526084016103cd565b60008686848181106108fc576108fc610f15565b9050602002013590508060000361096f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f4d696e744d616e616765723a2073686172652063616e6e6f742062652030000060448201526064016103cd565b6109798185610f44565b600180548082019091557fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601805473ffffffffffffffffffffffffffffffffffffffff9094167fffffffffffffffffffffffff00000000000000000000000000000000000000009094168417905560009283526002602052604090922055915080610a0381610ec4565b915050610812565b50620186a08114610a78576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f4d696e744d616e616765723a20696e76616c696420736861726573000000000060448201526064016103cd565b508015610adc57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b60008080610b117f000000000000000000000000000000000000000000000000000000000000000085610eb0565b610b1c906001610f44565b90506000610b4a7f000000000000000000000000000000000000000000000000000000000000000086610f5c565b905080600003610b8557610b5f600183610f70565b91507f000000000000000000000000000000000000000000000000000000000000000090505b9094909350915050565b6000807f00000000000000000000000000000000000000000000000000000000000000008180610bbe86610ae3565b909250905060015b82811015610c6c57610bf87f000000000000000000000000000000000000000000000000000000000000000085610e44565b610c029086610f44565b9450620186a0610c327f000000000000000000000000000000000000000000000000000000000000000086610e44565b610c3c9190610eb0565b93506402540be400610c4e8186610eb0565b610c589190610e44565b935080610c6481610ec4565b915050610bc6565b508015610c8a57610c7d8184610e44565b610c879085610f44565b93505b5091949350505050565b600060208284031215610ca657600080fd5b5035919050565b600060208284031215610cbf57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610ce357600080fd5b9392505050565b600060208083528351808285015260005b81811015610d1757858101830151858201604001528201610cfb565b81811115610d29576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60008083601f840112610d6f57600080fd5b50813567ffffffffffffffff811115610d8757600080fd5b6020830191508360208260051b8501011115610da257600080fd5b9250929050565b60008060008060408587031215610dbf57600080fd5b843567ffffffffffffffff80821115610dd757600080fd5b610de388838901610d5d565b90965094506020870135915080821115610dfc57600080fd5b50610e0987828801610d5d565b95989497509550505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610e7c57610e7c610e15565b500290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082610ebf57610ebf610e81565b500490565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610ef557610ef5610e15565b5060010190565b600060208284031215610f0e57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008219821115610f5757610f57610e15565b500190565b600082610f6b57610f6b610e81565b500690565b600082821015610f8257610f82610e15565b50039056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(MintManagerStorageLayoutJSON), MintManagerStorageLayout); err != nil {
		panic(err)
	}

	layouts["MintManager"] = MintManagerStorageLayout
	deployedBytecodes["MintManager"] = MintManagerDeployedBin
	immutableReferences["MintManager"] = true
}
