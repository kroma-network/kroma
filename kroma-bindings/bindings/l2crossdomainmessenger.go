// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// L2CrossDomainMessengerMetaData contains all meta data concerning the L2CrossDomainMessenger contract.
var L2CrossDomainMessengerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_l1CrossDomainMessenger\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MESSAGE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_GAS_CALLDATA_OVERHEAD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OTHER_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RELAY_CALL_OVERHEAD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RELAY_CONSTANT_OVERHEAD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RELAY_GAS_CHECK_BUFFER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RELAY_RESERVED_GAS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"baseGas\",\"inputs\":[{\"name\":\"_message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_minGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"failedMessages\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"messageNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"relayMessage\",\"inputs\":[{\"name\":\"_nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"sendMessage\",\"inputs\":[{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_minGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"successfulMessages\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xDomainMessageSender\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"FailedRelayedMessage\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RelayedMessage\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SentMessage\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"messageNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001af038038062001af0833981016040819052620000349162000216565b6001600160a01b0381166080526200004b62000052565b5062000248565b600054610100900460ff1615808015620000735750600054600160ff909116105b80620000a3575062000090306200018460201b620010b61760201c565b158015620000a3575060005460ff166001145b6200010c5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff19166001179055801562000130576000805461ff0019166101001790555b6200013a62000193565b801562000181576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b6001600160a01b03163b151590565b600054610100900460ff16620002005760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b606482015260840162000103565b606680546001600160a01b03191661dead179055565b6000602082840312156200022957600080fd5b81516001600160a01b03811681146200024157600080fd5b9392505050565b60805161187e6200027260003960008181610301015281816103f90152610a0d015261187e6000f3fe6080604052600436106101445760003560e01c80636e296e45116100c0578063a4e7f8bd11610074578063b28ade2511610059578063b28ade2514610383578063d764ad0b146103a3578063ecc70428146103b657600080fd5b8063a4e7f8bd14610323578063b1b1b2091461035357600080fd5b806383a74074116100a557806383a74074146102d85780638cbeeef2146101e35780639fce812c146102ef57600080fd5b80636e296e45146102895780638129fc1c146102c357600080fd5b80633f827a5a1161011757806354fd4d50116100fc57806354fd4d50146101f95780635644cfdf1461024f5780635c975abb1461026557600080fd5b80633f827a5a146101bb5780634c1d6a69146101e357600080fd5b8063028f85f7146101495780630c5684981461017c5780632828d7e8146101915780633dbb202b146101a6575b600080fd5b34801561015557600080fd5b5061015e601081565b60405167ffffffffffffffff90911681526020015b60405180910390f35b34801561018857600080fd5b5061015e603f81565b34801561019d57600080fd5b5061015e604081565b6101b96101b43660046113fa565b6103f4565b005b3480156101c757600080fd5b506101d0600081565b60405161ffff9091168152602001610173565b3480156101ef57600080fd5b5061015e619c4081565b34801561020557600080fd5b506102426040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b60405161017391906114ca565b34801561025b57600080fd5b5061015e61138881565b34801561027157600080fd5b5060335460ff165b6040519015158152602001610173565b34801561029557600080fd5b5061029e6105e2565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610173565b3480156102cf57600080fd5b506101b96106ce565b3480156102e457600080fd5b5061015e62030d4081565b3480156102fb57600080fd5b5061029e7f000000000000000000000000000000000000000000000000000000000000000081565b34801561032f57600080fd5b5061027961033e3660046114e4565b60686020526000908152604090205460ff1681565b34801561035f57600080fd5b5061027961036e3660046114e4565b60656020526000908152604090205460ff1681565b34801561038f57600080fd5b5061015e61039e3660046114fd565b610860565b6101b96103b1366004611551565b6108ce565b3480156103c257600080fd5b506067547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff16604051908152602001610173565b61050c7f0000000000000000000000000000000000000000000000000000000000000000610423858585610860565b347fd764ad0b0000000000000000000000000000000000000000000000000000000061046e6067547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b338a34898c8c60405160240161048a979695949392919061161c565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526110d2565b3373ffffffffffffffffffffffffffffffffffffffff85167fdd28cef75ff18fb538e43317144469f339702f973eace2bc808f2acc37db310e3486866105716067547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b8760405161058395949392919061167b565b60405180910390a35050606780547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b60665460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff2153016106b1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f7420736574000000000000000000000060648201526084015b60405180910390fd5b5060665473ffffffffffffffffffffffffffffffffffffffff1690565b600054610100900460ff16158080156106ee5750600054600160ff909116105b806107085750303b158015610708575060005460ff166001145b610794576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016106a8565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156107f257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6107fa611160565b801561085d57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b6000611388619c4080603f61087c604063ffffffff88166116e2565b6108869190611712565b6108916010886116e2565b61089e9062030d40611760565b6108a89190611760565b6108b29190611760565b6108bc9190611760565b6108c69190611760565b949350505050565b60f087901c60018110610989576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604760248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2030206d6573736167657320697320737570706f7274656420617420746860648201527f69732074696d6500000000000000000000000000000000000000000000000000608482015260a4016106a8565b60006109cf898989898989898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061122592505050565b905073ffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffeeeeffffffffffffffffffffffffffffffffeeef330181167f000000000000000000000000000000000000000000000000000000000000000090911603610a6757853414610a4357610a4361178c565b60008181526068602052604090205460ff1615610a6257610a6261178c565b610bb9565b3415610b1b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f20756e6c657373206d6573736167652069732066726f6d2060648201527f612073797374656d206164647265737300000000000000000000000000000000608482015260a4016106a8565b60008181526068602052604090205460ff16610bb9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c617965640000000000000000000000000000000060648201526084016106a8565b610bc287611248565b15610c75576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a4016106a8565b60008181526065602052604090205460ff1615610d14576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c617965640000000000000000000060648201526084016106a8565b610d3585610d26611388619c40611760565b67ffffffffffffffff1661129d565b1580610d5b575060665473ffffffffffffffffffffffffffffffffffffffff1661dead14155b15610e745760008181526068602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555182917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff3201610e6d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d6573736167650000000000000000000000000000000000000060648201526084016106a8565b50506110ad565b606680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a161790556000610f0588619c405a610ec891906117bb565b8988888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506112bb92505050565b606680547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead17905590508015610f9c5760008281526065602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a26110a9565b60008281526068602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff32016110a9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d6573736167650000000000000000000000000000000000000060648201526084016106a8565b5050505b50505050505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b6040517fc2b3e5ac0000000000000000000000000000000000000000000000000000000081527342000000000000000000000000000000000000039063c2b3e5ac908490611128908890889087906004016117d2565b6000604051808303818588803b15801561114157600080fd5b505af1158015611155573d6000803e3d6000fd5b505050505050505050565b600054610100900460ff166111f7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106a8565b606680547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055565b60006112358787878787876112d5565b8051906020012090509695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff8216301480611297575073ffffffffffffffffffffffffffffffffffffffff8216734200000000000000000000000000000000000003145b92915050565b600080603f83619c4001026040850201603f5a021015949350505050565b600080600080845160208601878a8af19695505050505050565b60608686868686866040516024016112f29695949392919061181a565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461139857600080fd5b919050565b60008083601f8401126113af57600080fd5b50813567ffffffffffffffff8111156113c757600080fd5b6020830191508360208285010111156113df57600080fd5b9250929050565b803563ffffffff8116811461139857600080fd5b6000806000806060858703121561141057600080fd5b61141985611374565b9350602085013567ffffffffffffffff81111561143557600080fd5b6114418782880161139d565b90945092506114549050604086016113e6565b905092959194509250565b6000815180845260005b8181101561148557602081850181015186830182015201611469565b81811115611497576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006114dd602083018461145f565b9392505050565b6000602082840312156114f657600080fd5b5035919050565b60008060006040848603121561151257600080fd5b833567ffffffffffffffff81111561152957600080fd5b6115358682870161139d565b90945092506115489050602085016113e6565b90509250925092565b600080600080600080600060c0888a03121561156c57600080fd5b8735965061157c60208901611374565b955061158a60408901611374565b9450606088013593506080880135925060a088013567ffffffffffffffff8111156115b457600080fd5b6115c08a828b0161139d565b989b979a50959850939692959293505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a083015261166e60c0830184866115d3565b9998505050505050505050565b8581526080602082015260006116956080830186886115d3565b905083604083015263ffffffff831660608301529695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff80831681851681830481118215151615611709576117096116b3565b02949350505050565b600067ffffffffffffffff80841680611754577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600067ffffffffffffffff808316818516808303821115611783576117836116b3565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b6000828210156117cd576117cd6116b3565b500390565b73ffffffffffffffffffffffffffffffffffffffff8416815267ffffffffffffffff83166020820152606060408201526000611811606083018461145f565b95945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a083015261186560c083018461145f565b9897505050505050505056fea164736f6c634300080f000a",
}

// L2CrossDomainMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use L2CrossDomainMessengerMetaData.ABI instead.
var L2CrossDomainMessengerABI = L2CrossDomainMessengerMetaData.ABI

// L2CrossDomainMessengerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2CrossDomainMessengerMetaData.Bin instead.
var L2CrossDomainMessengerBin = L2CrossDomainMessengerMetaData.Bin

// DeployL2CrossDomainMessenger deploys a new Ethereum contract, binding an instance of L2CrossDomainMessenger to it.
func DeployL2CrossDomainMessenger(auth *bind.TransactOpts, backend bind.ContractBackend, _l1CrossDomainMessenger common.Address) (common.Address, *types.Transaction, *L2CrossDomainMessenger, error) {
	parsed, err := L2CrossDomainMessengerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2CrossDomainMessengerBin), backend, _l1CrossDomainMessenger)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2CrossDomainMessenger{L2CrossDomainMessengerCaller: L2CrossDomainMessengerCaller{contract: contract}, L2CrossDomainMessengerTransactor: L2CrossDomainMessengerTransactor{contract: contract}, L2CrossDomainMessengerFilterer: L2CrossDomainMessengerFilterer{contract: contract}}, nil
}

// L2CrossDomainMessenger is an auto generated Go binding around an Ethereum contract.
type L2CrossDomainMessenger struct {
	L2CrossDomainMessengerCaller     // Read-only binding to the contract
	L2CrossDomainMessengerTransactor // Write-only binding to the contract
	L2CrossDomainMessengerFilterer   // Log filterer for contract events
}

// L2CrossDomainMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2CrossDomainMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2CrossDomainMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2CrossDomainMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2CrossDomainMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2CrossDomainMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2CrossDomainMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2CrossDomainMessengerSession struct {
	Contract     *L2CrossDomainMessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// L2CrossDomainMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2CrossDomainMessengerCallerSession struct {
	Contract *L2CrossDomainMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// L2CrossDomainMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2CrossDomainMessengerTransactorSession struct {
	Contract     *L2CrossDomainMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// L2CrossDomainMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2CrossDomainMessengerRaw struct {
	Contract *L2CrossDomainMessenger // Generic contract binding to access the raw methods on
}

// L2CrossDomainMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2CrossDomainMessengerCallerRaw struct {
	Contract *L2CrossDomainMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// L2CrossDomainMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2CrossDomainMessengerTransactorRaw struct {
	Contract *L2CrossDomainMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2CrossDomainMessenger creates a new instance of L2CrossDomainMessenger, bound to a specific deployed contract.
func NewL2CrossDomainMessenger(address common.Address, backend bind.ContractBackend) (*L2CrossDomainMessenger, error) {
	contract, err := bindL2CrossDomainMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessenger{L2CrossDomainMessengerCaller: L2CrossDomainMessengerCaller{contract: contract}, L2CrossDomainMessengerTransactor: L2CrossDomainMessengerTransactor{contract: contract}, L2CrossDomainMessengerFilterer: L2CrossDomainMessengerFilterer{contract: contract}}, nil
}

// NewL2CrossDomainMessengerCaller creates a new read-only instance of L2CrossDomainMessenger, bound to a specific deployed contract.
func NewL2CrossDomainMessengerCaller(address common.Address, caller bind.ContractCaller) (*L2CrossDomainMessengerCaller, error) {
	contract, err := bindL2CrossDomainMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerCaller{contract: contract}, nil
}

// NewL2CrossDomainMessengerTransactor creates a new write-only instance of L2CrossDomainMessenger, bound to a specific deployed contract.
func NewL2CrossDomainMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*L2CrossDomainMessengerTransactor, error) {
	contract, err := bindL2CrossDomainMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerTransactor{contract: contract}, nil
}

// NewL2CrossDomainMessengerFilterer creates a new log filterer instance of L2CrossDomainMessenger, bound to a specific deployed contract.
func NewL2CrossDomainMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*L2CrossDomainMessengerFilterer, error) {
	contract, err := bindL2CrossDomainMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerFilterer{contract: contract}, nil
}

// bindL2CrossDomainMessenger binds a generic wrapper to an already deployed contract.
func bindL2CrossDomainMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2CrossDomainMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2CrossDomainMessenger *L2CrossDomainMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2CrossDomainMessenger.Contract.L2CrossDomainMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2CrossDomainMessenger *L2CrossDomainMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.L2CrossDomainMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2CrossDomainMessenger *L2CrossDomainMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.L2CrossDomainMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2CrossDomainMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.contract.Transact(opts, method, params...)
}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) MESSAGEVERSION(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "MESSAGE_VERSION")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) MESSAGEVERSION() (uint16, error) {
	return _L2CrossDomainMessenger.Contract.MESSAGEVERSION(&_L2CrossDomainMessenger.CallOpts)
}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) MESSAGEVERSION() (uint16, error) {
	return _L2CrossDomainMessenger.Contract.MESSAGEVERSION(&_L2CrossDomainMessenger.CallOpts)
}

// MINGASCALLDATAOVERHEAD is a free data retrieval call binding the contract method 0x028f85f7.
//
// Solidity: function MIN_GAS_CALLDATA_OVERHEAD() view returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) MINGASCALLDATAOVERHEAD(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "MIN_GAS_CALLDATA_OVERHEAD")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MINGASCALLDATAOVERHEAD is a free data retrieval call binding the contract method 0x028f85f7.
//
// Solidity: function MIN_GAS_CALLDATA_OVERHEAD() view returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) MINGASCALLDATAOVERHEAD() (uint64, error) {
	return _L2CrossDomainMessenger.Contract.MINGASCALLDATAOVERHEAD(&_L2CrossDomainMessenger.CallOpts)
}

// MINGASCALLDATAOVERHEAD is a free data retrieval call binding the contract method 0x028f85f7.
//
// Solidity: function MIN_GAS_CALLDATA_OVERHEAD() view returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) MINGASCALLDATAOVERHEAD() (uint64, error) {
	return _L2CrossDomainMessenger.Contract.MINGASCALLDATAOVERHEAD(&_L2CrossDomainMessenger.CallOpts)
}

// MINGASDYNAMICOVERHEADDENOMINATOR is a free data retrieval call binding the contract method 0x0c568498.
//
// Solidity: function MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR() view returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) MINGASDYNAMICOVERHEADDENOMINATOR(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MINGASDYNAMICOVERHEADDENOMINATOR is a free data retrieval call binding the contract method 0x0c568498.
//
// Solidity: function MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR() view returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) MINGASDYNAMICOVERHEADDENOMINATOR() (uint64, error) {
	return _L2CrossDomainMessenger.Contract.MINGASDYNAMICOVERHEADDENOMINATOR(&_L2CrossDomainMessenger.CallOpts)
}

// MINGASDYNAMICOVERHEADDENOMINATOR is a free data retrieval call binding the contract method 0x0c568498.
//
// Solidity: function MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR() view returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) MINGASDYNAMICOVERHEADDENOMINATOR() (uint64, error) {
	return _L2CrossDomainMessenger.Contract.MINGASDYNAMICOVERHEADDENOMINATOR(&_L2CrossDomainMessenger.CallOpts)
}

// MINGASDYNAMICOVERHEADNUMERATOR is a free data retrieval call binding the contract method 0x2828d7e8.
//
// Solidity: function MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR() view returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) MINGASDYNAMICOVERHEADNUMERATOR(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MINGASDYNAMICOVERHEADNUMERATOR is a free data retrieval call binding the contract method 0x2828d7e8.
//
// Solidity: function MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR() view returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) MINGASDYNAMICOVERHEADNUMERATOR() (uint64, error) {
	return _L2CrossDomainMessenger.Contract.MINGASDYNAMICOVERHEADNUMERATOR(&_L2CrossDomainMessenger.CallOpts)
}

// MINGASDYNAMICOVERHEADNUMERATOR is a free data retrieval call binding the contract method 0x2828d7e8.
//
// Solidity: function MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR() view returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) MINGASDYNAMICOVERHEADNUMERATOR() (uint64, error) {
	return _L2CrossDomainMessenger.Contract.MINGASDYNAMICOVERHEADNUMERATOR(&_L2CrossDomainMessenger.CallOpts)
}

// OTHERMESSENGER is a free data retrieval call binding the contract method 0x9fce812c.
//
// Solidity: function OTHER_MESSENGER() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) OTHERMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "OTHER_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OTHERMESSENGER is a free data retrieval call binding the contract method 0x9fce812c.
//
// Solidity: function OTHER_MESSENGER() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) OTHERMESSENGER() (common.Address, error) {
	return _L2CrossDomainMessenger.Contract.OTHERMESSENGER(&_L2CrossDomainMessenger.CallOpts)
}

// OTHERMESSENGER is a free data retrieval call binding the contract method 0x9fce812c.
//
// Solidity: function OTHER_MESSENGER() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) OTHERMESSENGER() (common.Address, error) {
	return _L2CrossDomainMessenger.Contract.OTHERMESSENGER(&_L2CrossDomainMessenger.CallOpts)
}

// RELAYCALLOVERHEAD is a free data retrieval call binding the contract method 0x4c1d6a69.
//
// Solidity: function RELAY_CALL_OVERHEAD() view returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) RELAYCALLOVERHEAD(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "RELAY_CALL_OVERHEAD")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RELAYCALLOVERHEAD is a free data retrieval call binding the contract method 0x4c1d6a69.
//
// Solidity: function RELAY_CALL_OVERHEAD() view returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) RELAYCALLOVERHEAD() (uint64, error) {
	return _L2CrossDomainMessenger.Contract.RELAYCALLOVERHEAD(&_L2CrossDomainMessenger.CallOpts)
}

// RELAYCALLOVERHEAD is a free data retrieval call binding the contract method 0x4c1d6a69.
//
// Solidity: function RELAY_CALL_OVERHEAD() view returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) RELAYCALLOVERHEAD() (uint64, error) {
	return _L2CrossDomainMessenger.Contract.RELAYCALLOVERHEAD(&_L2CrossDomainMessenger.CallOpts)
}

// RELAYCONSTANTOVERHEAD is a free data retrieval call binding the contract method 0x83a74074.
//
// Solidity: function RELAY_CONSTANT_OVERHEAD() view returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) RELAYCONSTANTOVERHEAD(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "RELAY_CONSTANT_OVERHEAD")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RELAYCONSTANTOVERHEAD is a free data retrieval call binding the contract method 0x83a74074.
//
// Solidity: function RELAY_CONSTANT_OVERHEAD() view returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) RELAYCONSTANTOVERHEAD() (uint64, error) {
	return _L2CrossDomainMessenger.Contract.RELAYCONSTANTOVERHEAD(&_L2CrossDomainMessenger.CallOpts)
}

// RELAYCONSTANTOVERHEAD is a free data retrieval call binding the contract method 0x83a74074.
//
// Solidity: function RELAY_CONSTANT_OVERHEAD() view returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) RELAYCONSTANTOVERHEAD() (uint64, error) {
	return _L2CrossDomainMessenger.Contract.RELAYCONSTANTOVERHEAD(&_L2CrossDomainMessenger.CallOpts)
}

// RELAYGASCHECKBUFFER is a free data retrieval call binding the contract method 0x5644cfdf.
//
// Solidity: function RELAY_GAS_CHECK_BUFFER() view returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) RELAYGASCHECKBUFFER(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "RELAY_GAS_CHECK_BUFFER")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RELAYGASCHECKBUFFER is a free data retrieval call binding the contract method 0x5644cfdf.
//
// Solidity: function RELAY_GAS_CHECK_BUFFER() view returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) RELAYGASCHECKBUFFER() (uint64, error) {
	return _L2CrossDomainMessenger.Contract.RELAYGASCHECKBUFFER(&_L2CrossDomainMessenger.CallOpts)
}

// RELAYGASCHECKBUFFER is a free data retrieval call binding the contract method 0x5644cfdf.
//
// Solidity: function RELAY_GAS_CHECK_BUFFER() view returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) RELAYGASCHECKBUFFER() (uint64, error) {
	return _L2CrossDomainMessenger.Contract.RELAYGASCHECKBUFFER(&_L2CrossDomainMessenger.CallOpts)
}

// RELAYRESERVEDGAS is a free data retrieval call binding the contract method 0x8cbeeef2.
//
// Solidity: function RELAY_RESERVED_GAS() view returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) RELAYRESERVEDGAS(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "RELAY_RESERVED_GAS")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RELAYRESERVEDGAS is a free data retrieval call binding the contract method 0x8cbeeef2.
//
// Solidity: function RELAY_RESERVED_GAS() view returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) RELAYRESERVEDGAS() (uint64, error) {
	return _L2CrossDomainMessenger.Contract.RELAYRESERVEDGAS(&_L2CrossDomainMessenger.CallOpts)
}

// RELAYRESERVEDGAS is a free data retrieval call binding the contract method 0x8cbeeef2.
//
// Solidity: function RELAY_RESERVED_GAS() view returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) RELAYRESERVEDGAS() (uint64, error) {
	return _L2CrossDomainMessenger.Contract.RELAYRESERVEDGAS(&_L2CrossDomainMessenger.CallOpts)
}

// BaseGas is a free data retrieval call binding the contract method 0xb28ade25.
//
// Solidity: function baseGas(bytes _message, uint32 _minGasLimit) pure returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) BaseGas(opts *bind.CallOpts, _message []byte, _minGasLimit uint32) (uint64, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "baseGas", _message, _minGasLimit)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BaseGas is a free data retrieval call binding the contract method 0xb28ade25.
//
// Solidity: function baseGas(bytes _message, uint32 _minGasLimit) pure returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) BaseGas(_message []byte, _minGasLimit uint32) (uint64, error) {
	return _L2CrossDomainMessenger.Contract.BaseGas(&_L2CrossDomainMessenger.CallOpts, _message, _minGasLimit)
}

// BaseGas is a free data retrieval call binding the contract method 0xb28ade25.
//
// Solidity: function baseGas(bytes _message, uint32 _minGasLimit) pure returns(uint64)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) BaseGas(_message []byte, _minGasLimit uint32) (uint64, error) {
	return _L2CrossDomainMessenger.Contract.BaseGas(&_L2CrossDomainMessenger.CallOpts, _message, _minGasLimit)
}

// FailedMessages is a free data retrieval call binding the contract method 0xa4e7f8bd.
//
// Solidity: function failedMessages(bytes32 ) view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) FailedMessages(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "failedMessages", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FailedMessages is a free data retrieval call binding the contract method 0xa4e7f8bd.
//
// Solidity: function failedMessages(bytes32 ) view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) FailedMessages(arg0 [32]byte) (bool, error) {
	return _L2CrossDomainMessenger.Contract.FailedMessages(&_L2CrossDomainMessenger.CallOpts, arg0)
}

// FailedMessages is a free data retrieval call binding the contract method 0xa4e7f8bd.
//
// Solidity: function failedMessages(bytes32 ) view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) FailedMessages(arg0 [32]byte) (bool, error) {
	return _L2CrossDomainMessenger.Contract.FailedMessages(&_L2CrossDomainMessenger.CallOpts, arg0)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) MessageNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "messageNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) MessageNonce() (*big.Int, error) {
	return _L2CrossDomainMessenger.Contract.MessageNonce(&_L2CrossDomainMessenger.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) MessageNonce() (*big.Int, error) {
	return _L2CrossDomainMessenger.Contract.MessageNonce(&_L2CrossDomainMessenger.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) Paused() (bool, error) {
	return _L2CrossDomainMessenger.Contract.Paused(&_L2CrossDomainMessenger.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) Paused() (bool, error) {
	return _L2CrossDomainMessenger.Contract.Paused(&_L2CrossDomainMessenger.CallOpts)
}

// SuccessfulMessages is a free data retrieval call binding the contract method 0xb1b1b209.
//
// Solidity: function successfulMessages(bytes32 ) view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) SuccessfulMessages(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "successfulMessages", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SuccessfulMessages is a free data retrieval call binding the contract method 0xb1b1b209.
//
// Solidity: function successfulMessages(bytes32 ) view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) SuccessfulMessages(arg0 [32]byte) (bool, error) {
	return _L2CrossDomainMessenger.Contract.SuccessfulMessages(&_L2CrossDomainMessenger.CallOpts, arg0)
}

// SuccessfulMessages is a free data retrieval call binding the contract method 0xb1b1b209.
//
// Solidity: function successfulMessages(bytes32 ) view returns(bool)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) SuccessfulMessages(arg0 [32]byte) (bool, error) {
	return _L2CrossDomainMessenger.Contract.SuccessfulMessages(&_L2CrossDomainMessenger.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) Version() (string, error) {
	return _L2CrossDomainMessenger.Contract.Version(&_L2CrossDomainMessenger.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) Version() (string, error) {
	return _L2CrossDomainMessenger.Contract.Version(&_L2CrossDomainMessenger.CallOpts)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCaller) XDomainMessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2CrossDomainMessenger.contract.Call(opts, &out, "xDomainMessageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) XDomainMessageSender() (common.Address, error) {
	return _L2CrossDomainMessenger.Contract.XDomainMessageSender(&_L2CrossDomainMessenger.CallOpts)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerCallerSession) XDomainMessageSender() (common.Address, error) {
	return _L2CrossDomainMessenger.Contract.XDomainMessageSender(&_L2CrossDomainMessenger.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) Initialize() (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.Initialize(&_L2CrossDomainMessenger.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactorSession) Initialize() (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.Initialize(&_L2CrossDomainMessenger.TransactOpts)
}

// RelayMessage is a paid mutator transaction binding the contract method 0xd764ad0b.
//
// Solidity: function relayMessage(uint256 _nonce, address _sender, address _target, uint256 _value, uint256 _minGasLimit, bytes _message) payable returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactor) RelayMessage(opts *bind.TransactOpts, _nonce *big.Int, _sender common.Address, _target common.Address, _value *big.Int, _minGasLimit *big.Int, _message []byte) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.contract.Transact(opts, "relayMessage", _nonce, _sender, _target, _value, _minGasLimit, _message)
}

// RelayMessage is a paid mutator transaction binding the contract method 0xd764ad0b.
//
// Solidity: function relayMessage(uint256 _nonce, address _sender, address _target, uint256 _value, uint256 _minGasLimit, bytes _message) payable returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) RelayMessage(_nonce *big.Int, _sender common.Address, _target common.Address, _value *big.Int, _minGasLimit *big.Int, _message []byte) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.RelayMessage(&_L2CrossDomainMessenger.TransactOpts, _nonce, _sender, _target, _value, _minGasLimit, _message)
}

// RelayMessage is a paid mutator transaction binding the contract method 0xd764ad0b.
//
// Solidity: function relayMessage(uint256 _nonce, address _sender, address _target, uint256 _value, uint256 _minGasLimit, bytes _message) payable returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactorSession) RelayMessage(_nonce *big.Int, _sender common.Address, _target common.Address, _value *big.Int, _minGasLimit *big.Int, _message []byte) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.RelayMessage(&_L2CrossDomainMessenger.TransactOpts, _nonce, _sender, _target, _value, _minGasLimit, _message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x3dbb202b.
//
// Solidity: function sendMessage(address _target, bytes _message, uint32 _minGasLimit) payable returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactor) SendMessage(opts *bind.TransactOpts, _target common.Address, _message []byte, _minGasLimit uint32) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.contract.Transact(opts, "sendMessage", _target, _message, _minGasLimit)
}

// SendMessage is a paid mutator transaction binding the contract method 0x3dbb202b.
//
// Solidity: function sendMessage(address _target, bytes _message, uint32 _minGasLimit) payable returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerSession) SendMessage(_target common.Address, _message []byte, _minGasLimit uint32) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.SendMessage(&_L2CrossDomainMessenger.TransactOpts, _target, _message, _minGasLimit)
}

// SendMessage is a paid mutator transaction binding the contract method 0x3dbb202b.
//
// Solidity: function sendMessage(address _target, bytes _message, uint32 _minGasLimit) payable returns()
func (_L2CrossDomainMessenger *L2CrossDomainMessengerTransactorSession) SendMessage(_target common.Address, _message []byte, _minGasLimit uint32) (*types.Transaction, error) {
	return _L2CrossDomainMessenger.Contract.SendMessage(&_L2CrossDomainMessenger.TransactOpts, _target, _message, _minGasLimit)
}

// L2CrossDomainMessengerFailedRelayedMessageIterator is returned from FilterFailedRelayedMessage and is used to iterate over the raw logs and unpacked data for FailedRelayedMessage events raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerFailedRelayedMessageIterator struct {
	Event *L2CrossDomainMessengerFailedRelayedMessage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *L2CrossDomainMessengerFailedRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CrossDomainMessengerFailedRelayedMessage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(L2CrossDomainMessengerFailedRelayedMessage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *L2CrossDomainMessengerFailedRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CrossDomainMessengerFailedRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CrossDomainMessengerFailedRelayedMessage represents a FailedRelayedMessage event raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerFailedRelayedMessage struct {
	MsgHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFailedRelayedMessage is a free log retrieval operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed msgHash)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) FilterFailedRelayedMessage(opts *bind.FilterOpts, msgHash [][32]byte) (*L2CrossDomainMessengerFailedRelayedMessageIterator, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _L2CrossDomainMessenger.contract.FilterLogs(opts, "FailedRelayedMessage", msgHashRule)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerFailedRelayedMessageIterator{contract: _L2CrossDomainMessenger.contract, event: "FailedRelayedMessage", logs: logs, sub: sub}, nil
}

// WatchFailedRelayedMessage is a free log subscription operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed msgHash)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) WatchFailedRelayedMessage(opts *bind.WatchOpts, sink chan<- *L2CrossDomainMessengerFailedRelayedMessage, msgHash [][32]byte) (event.Subscription, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _L2CrossDomainMessenger.contract.WatchLogs(opts, "FailedRelayedMessage", msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CrossDomainMessengerFailedRelayedMessage)
				if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFailedRelayedMessage is a log parse operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed msgHash)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) ParseFailedRelayedMessage(log types.Log) (*L2CrossDomainMessengerFailedRelayedMessage, error) {
	event := new(L2CrossDomainMessengerFailedRelayedMessage)
	if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2CrossDomainMessengerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerInitializedIterator struct {
	Event *L2CrossDomainMessengerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *L2CrossDomainMessengerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CrossDomainMessengerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(L2CrossDomainMessengerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *L2CrossDomainMessengerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CrossDomainMessengerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CrossDomainMessengerInitialized represents a Initialized event raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2CrossDomainMessengerInitializedIterator, error) {

	logs, sub, err := _L2CrossDomainMessenger.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerInitializedIterator{contract: _L2CrossDomainMessenger.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2CrossDomainMessengerInitialized) (event.Subscription, error) {

	logs, sub, err := _L2CrossDomainMessenger.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CrossDomainMessengerInitialized)
				if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) ParseInitialized(log types.Log) (*L2CrossDomainMessengerInitialized, error) {
	event := new(L2CrossDomainMessengerInitialized)
	if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2CrossDomainMessengerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerPausedIterator struct {
	Event *L2CrossDomainMessengerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *L2CrossDomainMessengerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CrossDomainMessengerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(L2CrossDomainMessengerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *L2CrossDomainMessengerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CrossDomainMessengerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CrossDomainMessengerPaused represents a Paused event raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) FilterPaused(opts *bind.FilterOpts) (*L2CrossDomainMessengerPausedIterator, error) {

	logs, sub, err := _L2CrossDomainMessenger.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerPausedIterator{contract: _L2CrossDomainMessenger.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *L2CrossDomainMessengerPaused) (event.Subscription, error) {

	logs, sub, err := _L2CrossDomainMessenger.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CrossDomainMessengerPaused)
				if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) ParsePaused(log types.Log) (*L2CrossDomainMessengerPaused, error) {
	event := new(L2CrossDomainMessengerPaused)
	if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2CrossDomainMessengerRelayedMessageIterator is returned from FilterRelayedMessage and is used to iterate over the raw logs and unpacked data for RelayedMessage events raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerRelayedMessageIterator struct {
	Event *L2CrossDomainMessengerRelayedMessage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *L2CrossDomainMessengerRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CrossDomainMessengerRelayedMessage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(L2CrossDomainMessengerRelayedMessage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *L2CrossDomainMessengerRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CrossDomainMessengerRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CrossDomainMessengerRelayedMessage represents a RelayedMessage event raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerRelayedMessage struct {
	MsgHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayedMessage is a free log retrieval operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed msgHash)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) FilterRelayedMessage(opts *bind.FilterOpts, msgHash [][32]byte) (*L2CrossDomainMessengerRelayedMessageIterator, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _L2CrossDomainMessenger.contract.FilterLogs(opts, "RelayedMessage", msgHashRule)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerRelayedMessageIterator{contract: _L2CrossDomainMessenger.contract, event: "RelayedMessage", logs: logs, sub: sub}, nil
}

// WatchRelayedMessage is a free log subscription operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed msgHash)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) WatchRelayedMessage(opts *bind.WatchOpts, sink chan<- *L2CrossDomainMessengerRelayedMessage, msgHash [][32]byte) (event.Subscription, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _L2CrossDomainMessenger.contract.WatchLogs(opts, "RelayedMessage", msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CrossDomainMessengerRelayedMessage)
				if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayedMessage is a log parse operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed msgHash)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) ParseRelayedMessage(log types.Log) (*L2CrossDomainMessengerRelayedMessage, error) {
	event := new(L2CrossDomainMessengerRelayedMessage)
	if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2CrossDomainMessengerSentMessageIterator is returned from FilterSentMessage and is used to iterate over the raw logs and unpacked data for SentMessage events raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerSentMessageIterator struct {
	Event *L2CrossDomainMessengerSentMessage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *L2CrossDomainMessengerSentMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CrossDomainMessengerSentMessage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(L2CrossDomainMessengerSentMessage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *L2CrossDomainMessengerSentMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CrossDomainMessengerSentMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CrossDomainMessengerSentMessage represents a SentMessage event raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerSentMessage struct {
	Target       common.Address
	Sender       common.Address
	Value        *big.Int
	Message      []byte
	MessageNonce *big.Int
	GasLimit     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSentMessage is a free log retrieval operation binding the contract event 0xdd28cef75ff18fb538e43317144469f339702f973eace2bc808f2acc37db310e.
//
// Solidity: event SentMessage(address indexed target, address indexed sender, uint256 value, bytes message, uint256 messageNonce, uint256 gasLimit)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) FilterSentMessage(opts *bind.FilterOpts, target []common.Address, sender []common.Address) (*L2CrossDomainMessengerSentMessageIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _L2CrossDomainMessenger.contract.FilterLogs(opts, "SentMessage", targetRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerSentMessageIterator{contract: _L2CrossDomainMessenger.contract, event: "SentMessage", logs: logs, sub: sub}, nil
}

// WatchSentMessage is a free log subscription operation binding the contract event 0xdd28cef75ff18fb538e43317144469f339702f973eace2bc808f2acc37db310e.
//
// Solidity: event SentMessage(address indexed target, address indexed sender, uint256 value, bytes message, uint256 messageNonce, uint256 gasLimit)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) WatchSentMessage(opts *bind.WatchOpts, sink chan<- *L2CrossDomainMessengerSentMessage, target []common.Address, sender []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _L2CrossDomainMessenger.contract.WatchLogs(opts, "SentMessage", targetRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CrossDomainMessengerSentMessage)
				if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSentMessage is a log parse operation binding the contract event 0xdd28cef75ff18fb538e43317144469f339702f973eace2bc808f2acc37db310e.
//
// Solidity: event SentMessage(address indexed target, address indexed sender, uint256 value, bytes message, uint256 messageNonce, uint256 gasLimit)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) ParseSentMessage(log types.Log) (*L2CrossDomainMessengerSentMessage, error) {
	event := new(L2CrossDomainMessengerSentMessage)
	if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2CrossDomainMessengerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerUnpausedIterator struct {
	Event *L2CrossDomainMessengerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *L2CrossDomainMessengerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2CrossDomainMessengerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(L2CrossDomainMessengerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *L2CrossDomainMessengerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2CrossDomainMessengerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2CrossDomainMessengerUnpaused represents a Unpaused event raised by the L2CrossDomainMessenger contract.
type L2CrossDomainMessengerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*L2CrossDomainMessengerUnpausedIterator, error) {

	logs, sub, err := _L2CrossDomainMessenger.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &L2CrossDomainMessengerUnpausedIterator{contract: _L2CrossDomainMessenger.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *L2CrossDomainMessengerUnpaused) (event.Subscription, error) {

	logs, sub, err := _L2CrossDomainMessenger.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2CrossDomainMessengerUnpaused)
				if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L2CrossDomainMessenger *L2CrossDomainMessengerFilterer) ParseUnpaused(log types.Log) (*L2CrossDomainMessengerUnpaused, error) {
	event := new(L2CrossDomainMessengerUnpaused)
	if err := _L2CrossDomainMessenger.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
