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

// KromaMintableERC20MetaData contains all meta data concerning the KromaMintableERC20 contract.
var KromaMintableERC20MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_remoteToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BRIDGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"REMOTE_TOKEN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Burn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60c06040523480156200001157600080fd5b50604051620014413803806200144183398101604081905262000034916200015a565b8181600362000044838262000279565b50600462000053828262000279565b5050506001600160a01b0392831660805250501660a05262000345565b80516001600160a01b03811681146200008857600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620000b557600080fd5b81516001600160401b0380821115620000d257620000d26200008d565b604051601f8301601f19908116603f01168101908282118183101715620000fd57620000fd6200008d565b816040528381526020925086838588010111156200011a57600080fd5b600091505b838210156200013e57858201830151818301840152908201906200011f565b83821115620001505760008385830101525b9695505050505050565b600080600080608085870312156200017157600080fd5b6200017c8562000070565b93506200018c6020860162000070565b60408601519093506001600160401b0380821115620001aa57600080fd5b620001b888838901620000a3565b93506060870151915080821115620001cf57600080fd5b50620001de87828801620000a3565b91505092959194509250565b600181811c90821680620001ff57607f821691505b6020821081036200022057634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200027457600081815260208120601f850160051c810160208610156200024f5750805b601f850160051c820191505b8181101562000270578281556001016200025b565b5050505b505050565b81516001600160401b038111156200029557620002956200008d565b620002ad81620002a68454620001ea565b8462000226565b602080601f831160018114620002e55760008415620002cc5750858301515b600019600386901b1c1916600185901b17855562000270565b600085815260208120601f198616915b828110156200031657888601518255948401946001909101908401620002f5565b5085821015620003355787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a0516110c862000379600039600081816103160152818161050a01526106270152600061014d01526110c86000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c806340c10f19116100b25780639dc29fac11610081578063a9059cbb11610066578063a9059cbb146102b8578063dd62ed3e146102cb578063ee9a31a21461031157600080fd5b80639dc29fac14610292578063a457c2d7146102a557600080fd5b806340c10f191461020357806354fd4d501461021857806370a082311461025457806395d89b411461028a57600080fd5b806318160ddd116100ee57806318160ddd146101bc57806323b872dd146101ce578063313ce567146101e157806339509351146101f057600080fd5b806301ffc9a714610120578063033964be1461014857806306fdde0314610194578063095ea7b3146101a9575b600080fd5b61013361012e366004610e90565b610338565b60405190151581526020015b60405180910390f35b61016f7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161013f565b61019c6103d8565b60405161013f9190610ed9565b6101336101b7366004610f75565b61046a565b6002545b60405190815260200161013f565b6101336101dc366004610f9f565b610482565b6040516012815260200161013f565b6101336101fe366004610f75565b6104a6565b610216610211366004610f75565b6104f2565b005b61019c6040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b6101c0610262366004610fdb565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b61019c610600565b6102166102a0366004610f75565b61060f565b6101336102b3366004610f75565b61070c565b6101336102c6366004610f75565b6107c3565b6101c06102d9366004610ff6565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b61016f7f000000000000000000000000000000000000000000000000000000000000000081565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007f30a0c5a9000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000084168214806103d057507fffffffff00000000000000000000000000000000000000000000000000000000848116908216145b949350505050565b6060600380546103e790611029565b80601f016020809104026020016040519081016040528092919081815260200182805461041390611029565b80156104605780601f1061043557610100808354040283529160200191610460565b820191906000526020600020905b81548152906001019060200180831161044357829003601f168201915b5050505050905090565b6000336104788185856107d1565b5060019392505050565b600033610490858285610951565b61049b858585610a0e565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061047890829086906104ed90879061107c565b6107d1565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146105a25760405162461bcd60e51b815260206004820152603160248201527f4b726f6d614d696e7461626c6545524332303a206f6e6c79206272696467652060448201527f63616e206d696e7420616e64206275726e00000000000000000000000000000060648201526084015b60405180910390fd5b6105ac8282610c2f565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885826040516105f491815260200190565b60405180910390a25050565b6060600480546103e790611029565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146106ba5760405162461bcd60e51b815260206004820152603160248201527f4b726f6d614d696e7461626c6545524332303a206f6e6c79206272696467652060448201527f63616e206d696e7420616e64206275726e0000000000000000000000000000006064820152608401610599565b6106c48282610d08565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5826040516105f491815260200190565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152812054909190838110156107b65760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152608401610599565b61049b82868684036107d1565b600033610478818585610a0e565b73ffffffffffffffffffffffffffffffffffffffff83166108595760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610599565b73ffffffffffffffffffffffffffffffffffffffff82166108e25760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610599565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610a0857818110156109fb5760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610599565b610a0884848484036107d1565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610a975760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610599565b73ffffffffffffffffffffffffffffffffffffffff8216610b205760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610599565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610bbc5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610599565b73ffffffffffffffffffffffffffffffffffffffff848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610a08565b73ffffffffffffffffffffffffffffffffffffffff8216610c925760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610599565b8060026000828254610ca4919061107c565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff8216610d915760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610599565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090205481811015610e2d5760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152608401610599565b73ffffffffffffffffffffffffffffffffffffffff83166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101610944565b600060208284031215610ea257600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610ed257600080fd5b9392505050565b600060208083528351808285015260005b81811015610f0657858101830151858201604001528201610eea565b81811115610f18576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610f7057600080fd5b919050565b60008060408385031215610f8857600080fd5b610f9183610f4c565b946020939093013593505050565b600080600060608486031215610fb457600080fd5b610fbd84610f4c565b9250610fcb60208501610f4c565b9150604084013590509250925092565b600060208284031215610fed57600080fd5b610ed282610f4c565b6000806040838503121561100957600080fd5b61101283610f4c565b915061102060208401610f4c565b90509250929050565b600181811c9082168061103d57607f821691505b602082108103611076577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b600082198211156110b6577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b50019056fea164736f6c634300080f000a",
}

// KromaMintableERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use KromaMintableERC20MetaData.ABI instead.
var KromaMintableERC20ABI = KromaMintableERC20MetaData.ABI

// KromaMintableERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KromaMintableERC20MetaData.Bin instead.
var KromaMintableERC20Bin = KromaMintableERC20MetaData.Bin

// DeployKromaMintableERC20 deploys a new Ethereum contract, binding an instance of KromaMintableERC20 to it.
func DeployKromaMintableERC20(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address, _remoteToken common.Address, _name string, _symbol string) (common.Address, *types.Transaction, *KromaMintableERC20, error) {
	parsed, err := KromaMintableERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KromaMintableERC20Bin), backend, _bridge, _remoteToken, _name, _symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KromaMintableERC20{KromaMintableERC20Caller: KromaMintableERC20Caller{contract: contract}, KromaMintableERC20Transactor: KromaMintableERC20Transactor{contract: contract}, KromaMintableERC20Filterer: KromaMintableERC20Filterer{contract: contract}}, nil
}

// KromaMintableERC20 is an auto generated Go binding around an Ethereum contract.
type KromaMintableERC20 struct {
	KromaMintableERC20Caller     // Read-only binding to the contract
	KromaMintableERC20Transactor // Write-only binding to the contract
	KromaMintableERC20Filterer   // Log filterer for contract events
}

// KromaMintableERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type KromaMintableERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KromaMintableERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type KromaMintableERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KromaMintableERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KromaMintableERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KromaMintableERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KromaMintableERC20Session struct {
	Contract     *KromaMintableERC20 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// KromaMintableERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KromaMintableERC20CallerSession struct {
	Contract *KromaMintableERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// KromaMintableERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KromaMintableERC20TransactorSession struct {
	Contract     *KromaMintableERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// KromaMintableERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type KromaMintableERC20Raw struct {
	Contract *KromaMintableERC20 // Generic contract binding to access the raw methods on
}

// KromaMintableERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KromaMintableERC20CallerRaw struct {
	Contract *KromaMintableERC20Caller // Generic read-only contract binding to access the raw methods on
}

// KromaMintableERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KromaMintableERC20TransactorRaw struct {
	Contract *KromaMintableERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewKromaMintableERC20 creates a new instance of KromaMintableERC20, bound to a specific deployed contract.
func NewKromaMintableERC20(address common.Address, backend bind.ContractBackend) (*KromaMintableERC20, error) {
	contract, err := bindKromaMintableERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KromaMintableERC20{KromaMintableERC20Caller: KromaMintableERC20Caller{contract: contract}, KromaMintableERC20Transactor: KromaMintableERC20Transactor{contract: contract}, KromaMintableERC20Filterer: KromaMintableERC20Filterer{contract: contract}}, nil
}

// NewKromaMintableERC20Caller creates a new read-only instance of KromaMintableERC20, bound to a specific deployed contract.
func NewKromaMintableERC20Caller(address common.Address, caller bind.ContractCaller) (*KromaMintableERC20Caller, error) {
	contract, err := bindKromaMintableERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KromaMintableERC20Caller{contract: contract}, nil
}

// NewKromaMintableERC20Transactor creates a new write-only instance of KromaMintableERC20, bound to a specific deployed contract.
func NewKromaMintableERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*KromaMintableERC20Transactor, error) {
	contract, err := bindKromaMintableERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KromaMintableERC20Transactor{contract: contract}, nil
}

// NewKromaMintableERC20Filterer creates a new log filterer instance of KromaMintableERC20, bound to a specific deployed contract.
func NewKromaMintableERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*KromaMintableERC20Filterer, error) {
	contract, err := bindKromaMintableERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KromaMintableERC20Filterer{contract: contract}, nil
}

// bindKromaMintableERC20 binds a generic wrapper to an already deployed contract.
func bindKromaMintableERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KromaMintableERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KromaMintableERC20 *KromaMintableERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KromaMintableERC20.Contract.KromaMintableERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KromaMintableERC20 *KromaMintableERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KromaMintableERC20.Contract.KromaMintableERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KromaMintableERC20 *KromaMintableERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KromaMintableERC20.Contract.KromaMintableERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KromaMintableERC20 *KromaMintableERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KromaMintableERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KromaMintableERC20 *KromaMintableERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KromaMintableERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KromaMintableERC20 *KromaMintableERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KromaMintableERC20.Contract.contract.Transact(opts, method, params...)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_KromaMintableERC20 *KromaMintableERC20Caller) BRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KromaMintableERC20.contract.Call(opts, &out, "BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_KromaMintableERC20 *KromaMintableERC20Session) BRIDGE() (common.Address, error) {
	return _KromaMintableERC20.Contract.BRIDGE(&_KromaMintableERC20.CallOpts)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_KromaMintableERC20 *KromaMintableERC20CallerSession) BRIDGE() (common.Address, error) {
	return _KromaMintableERC20.Contract.BRIDGE(&_KromaMintableERC20.CallOpts)
}

// REMOTETOKEN is a free data retrieval call binding the contract method 0x033964be.
//
// Solidity: function REMOTE_TOKEN() view returns(address)
func (_KromaMintableERC20 *KromaMintableERC20Caller) REMOTETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KromaMintableERC20.contract.Call(opts, &out, "REMOTE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REMOTETOKEN is a free data retrieval call binding the contract method 0x033964be.
//
// Solidity: function REMOTE_TOKEN() view returns(address)
func (_KromaMintableERC20 *KromaMintableERC20Session) REMOTETOKEN() (common.Address, error) {
	return _KromaMintableERC20.Contract.REMOTETOKEN(&_KromaMintableERC20.CallOpts)
}

// REMOTETOKEN is a free data retrieval call binding the contract method 0x033964be.
//
// Solidity: function REMOTE_TOKEN() view returns(address)
func (_KromaMintableERC20 *KromaMintableERC20CallerSession) REMOTETOKEN() (common.Address, error) {
	return _KromaMintableERC20.Contract.REMOTETOKEN(&_KromaMintableERC20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KromaMintableERC20 *KromaMintableERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KromaMintableERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KromaMintableERC20 *KromaMintableERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KromaMintableERC20.Contract.Allowance(&_KromaMintableERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KromaMintableERC20 *KromaMintableERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KromaMintableERC20.Contract.Allowance(&_KromaMintableERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KromaMintableERC20 *KromaMintableERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KromaMintableERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KromaMintableERC20 *KromaMintableERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _KromaMintableERC20.Contract.BalanceOf(&_KromaMintableERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KromaMintableERC20 *KromaMintableERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _KromaMintableERC20.Contract.BalanceOf(&_KromaMintableERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KromaMintableERC20 *KromaMintableERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _KromaMintableERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KromaMintableERC20 *KromaMintableERC20Session) Decimals() (uint8, error) {
	return _KromaMintableERC20.Contract.Decimals(&_KromaMintableERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KromaMintableERC20 *KromaMintableERC20CallerSession) Decimals() (uint8, error) {
	return _KromaMintableERC20.Contract.Decimals(&_KromaMintableERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KromaMintableERC20 *KromaMintableERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KromaMintableERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KromaMintableERC20 *KromaMintableERC20Session) Name() (string, error) {
	return _KromaMintableERC20.Contract.Name(&_KromaMintableERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KromaMintableERC20 *KromaMintableERC20CallerSession) Name() (string, error) {
	return _KromaMintableERC20.Contract.Name(&_KromaMintableERC20.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_KromaMintableERC20 *KromaMintableERC20Caller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _KromaMintableERC20.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_KromaMintableERC20 *KromaMintableERC20Session) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _KromaMintableERC20.Contract.SupportsInterface(&_KromaMintableERC20.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_KromaMintableERC20 *KromaMintableERC20CallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _KromaMintableERC20.Contract.SupportsInterface(&_KromaMintableERC20.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KromaMintableERC20 *KromaMintableERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KromaMintableERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KromaMintableERC20 *KromaMintableERC20Session) Symbol() (string, error) {
	return _KromaMintableERC20.Contract.Symbol(&_KromaMintableERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KromaMintableERC20 *KromaMintableERC20CallerSession) Symbol() (string, error) {
	return _KromaMintableERC20.Contract.Symbol(&_KromaMintableERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KromaMintableERC20 *KromaMintableERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KromaMintableERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KromaMintableERC20 *KromaMintableERC20Session) TotalSupply() (*big.Int, error) {
	return _KromaMintableERC20.Contract.TotalSupply(&_KromaMintableERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KromaMintableERC20 *KromaMintableERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _KromaMintableERC20.Contract.TotalSupply(&_KromaMintableERC20.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KromaMintableERC20 *KromaMintableERC20Caller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KromaMintableERC20.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KromaMintableERC20 *KromaMintableERC20Session) Version() (string, error) {
	return _KromaMintableERC20.Contract.Version(&_KromaMintableERC20.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KromaMintableERC20 *KromaMintableERC20CallerSession) Version() (string, error) {
	return _KromaMintableERC20.Contract.Version(&_KromaMintableERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KromaMintableERC20 *KromaMintableERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KromaMintableERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KromaMintableERC20 *KromaMintableERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KromaMintableERC20.Contract.Approve(&_KromaMintableERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KromaMintableERC20 *KromaMintableERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KromaMintableERC20.Contract.Approve(&_KromaMintableERC20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_KromaMintableERC20 *KromaMintableERC20Transactor) Burn(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _KromaMintableERC20.contract.Transact(opts, "burn", _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_KromaMintableERC20 *KromaMintableERC20Session) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _KromaMintableERC20.Contract.Burn(&_KromaMintableERC20.TransactOpts, _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_KromaMintableERC20 *KromaMintableERC20TransactorSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _KromaMintableERC20.Contract.Burn(&_KromaMintableERC20.TransactOpts, _from, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KromaMintableERC20 *KromaMintableERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KromaMintableERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KromaMintableERC20 *KromaMintableERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KromaMintableERC20.Contract.DecreaseAllowance(&_KromaMintableERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KromaMintableERC20 *KromaMintableERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KromaMintableERC20.Contract.DecreaseAllowance(&_KromaMintableERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KromaMintableERC20 *KromaMintableERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KromaMintableERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KromaMintableERC20 *KromaMintableERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KromaMintableERC20.Contract.IncreaseAllowance(&_KromaMintableERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KromaMintableERC20 *KromaMintableERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KromaMintableERC20.Contract.IncreaseAllowance(&_KromaMintableERC20.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_KromaMintableERC20 *KromaMintableERC20Transactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _KromaMintableERC20.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_KromaMintableERC20 *KromaMintableERC20Session) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _KromaMintableERC20.Contract.Mint(&_KromaMintableERC20.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_KromaMintableERC20 *KromaMintableERC20TransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _KromaMintableERC20.Contract.Mint(&_KromaMintableERC20.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_KromaMintableERC20 *KromaMintableERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KromaMintableERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_KromaMintableERC20 *KromaMintableERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KromaMintableERC20.Contract.Transfer(&_KromaMintableERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_KromaMintableERC20 *KromaMintableERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KromaMintableERC20.Contract.Transfer(&_KromaMintableERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_KromaMintableERC20 *KromaMintableERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KromaMintableERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_KromaMintableERC20 *KromaMintableERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KromaMintableERC20.Contract.TransferFrom(&_KromaMintableERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_KromaMintableERC20 *KromaMintableERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KromaMintableERC20.Contract.TransferFrom(&_KromaMintableERC20.TransactOpts, from, to, amount)
}

// KromaMintableERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KromaMintableERC20 contract.
type KromaMintableERC20ApprovalIterator struct {
	Event *KromaMintableERC20Approval // Event containing the contract specifics and raw log

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
func (it *KromaMintableERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KromaMintableERC20Approval)
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
		it.Event = new(KromaMintableERC20Approval)
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
func (it *KromaMintableERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KromaMintableERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KromaMintableERC20Approval represents a Approval event raised by the KromaMintableERC20 contract.
type KromaMintableERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KromaMintableERC20 *KromaMintableERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*KromaMintableERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KromaMintableERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &KromaMintableERC20ApprovalIterator{contract: _KromaMintableERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KromaMintableERC20 *KromaMintableERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KromaMintableERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KromaMintableERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KromaMintableERC20Approval)
				if err := _KromaMintableERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KromaMintableERC20 *KromaMintableERC20Filterer) ParseApproval(log types.Log) (*KromaMintableERC20Approval, error) {
	event := new(KromaMintableERC20Approval)
	if err := _KromaMintableERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KromaMintableERC20BurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the KromaMintableERC20 contract.
type KromaMintableERC20BurnIterator struct {
	Event *KromaMintableERC20Burn // Event containing the contract specifics and raw log

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
func (it *KromaMintableERC20BurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KromaMintableERC20Burn)
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
		it.Event = new(KromaMintableERC20Burn)
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
func (it *KromaMintableERC20BurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KromaMintableERC20BurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KromaMintableERC20Burn represents a Burn event raised by the KromaMintableERC20 contract.
type KromaMintableERC20Burn struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 amount)
func (_KromaMintableERC20 *KromaMintableERC20Filterer) FilterBurn(opts *bind.FilterOpts, account []common.Address) (*KromaMintableERC20BurnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _KromaMintableERC20.contract.FilterLogs(opts, "Burn", accountRule)
	if err != nil {
		return nil, err
	}
	return &KromaMintableERC20BurnIterator{contract: _KromaMintableERC20.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 amount)
func (_KromaMintableERC20 *KromaMintableERC20Filterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *KromaMintableERC20Burn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _KromaMintableERC20.contract.WatchLogs(opts, "Burn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KromaMintableERC20Burn)
				if err := _KromaMintableERC20.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 amount)
func (_KromaMintableERC20 *KromaMintableERC20Filterer) ParseBurn(log types.Log) (*KromaMintableERC20Burn, error) {
	event := new(KromaMintableERC20Burn)
	if err := _KromaMintableERC20.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KromaMintableERC20MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the KromaMintableERC20 contract.
type KromaMintableERC20MintIterator struct {
	Event *KromaMintableERC20Mint // Event containing the contract specifics and raw log

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
func (it *KromaMintableERC20MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KromaMintableERC20Mint)
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
		it.Event = new(KromaMintableERC20Mint)
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
func (it *KromaMintableERC20MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KromaMintableERC20MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KromaMintableERC20Mint represents a Mint event raised by the KromaMintableERC20 contract.
type KromaMintableERC20Mint struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_KromaMintableERC20 *KromaMintableERC20Filterer) FilterMint(opts *bind.FilterOpts, account []common.Address) (*KromaMintableERC20MintIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _KromaMintableERC20.contract.FilterLogs(opts, "Mint", accountRule)
	if err != nil {
		return nil, err
	}
	return &KromaMintableERC20MintIterator{contract: _KromaMintableERC20.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_KromaMintableERC20 *KromaMintableERC20Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *KromaMintableERC20Mint, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _KromaMintableERC20.contract.WatchLogs(opts, "Mint", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KromaMintableERC20Mint)
				if err := _KromaMintableERC20.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_KromaMintableERC20 *KromaMintableERC20Filterer) ParseMint(log types.Log) (*KromaMintableERC20Mint, error) {
	event := new(KromaMintableERC20Mint)
	if err := _KromaMintableERC20.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KromaMintableERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KromaMintableERC20 contract.
type KromaMintableERC20TransferIterator struct {
	Event *KromaMintableERC20Transfer // Event containing the contract specifics and raw log

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
func (it *KromaMintableERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KromaMintableERC20Transfer)
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
		it.Event = new(KromaMintableERC20Transfer)
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
func (it *KromaMintableERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KromaMintableERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KromaMintableERC20Transfer represents a Transfer event raised by the KromaMintableERC20 contract.
type KromaMintableERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KromaMintableERC20 *KromaMintableERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*KromaMintableERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KromaMintableERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &KromaMintableERC20TransferIterator{contract: _KromaMintableERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KromaMintableERC20 *KromaMintableERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KromaMintableERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KromaMintableERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KromaMintableERC20Transfer)
				if err := _KromaMintableERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KromaMintableERC20 *KromaMintableERC20Filterer) ParseTransfer(log types.Log) (*KromaMintableERC20Transfer, error) {
	event := new(KromaMintableERC20Transfer)
	if err := _KromaMintableERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
