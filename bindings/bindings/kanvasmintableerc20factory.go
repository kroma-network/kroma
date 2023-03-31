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

// KanvasMintableERC20FactoryMetaData contains all meta data concerning the KanvasMintableERC20Factory contract.
var KanvasMintableERC20FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"deployer\",\"type\":\"address\"}],\"name\":\"KanvasMintableERC20Created\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"createKanvasMintableERC20\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61010060405234801561001157600080fd5b5060405161216338038061216383398101604081905261003091610050565b60006080819052600160a05260c0526001600160a01b031660e052610080565b60006020828403121561006257600080fd5b81516001600160a01b038116811461007957600080fd5b9392505050565b60805160a05160c05160e0516120a56100be6000396000818160b0015261022b01526000610130015260006101050152600060da01526120a56000f3fe60806040523480156200001157600080fd5b5060043610620000465760003560e01c806354fd4d50146200004b5780638c2378ba146200006d578063ee9a31a214620000aa575b600080fd5b62000055620000d2565b604051620000649190620004c5565b60405180910390f35b620000846200007e366004620005c3565b6200017d565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200162000064565b620000847f000000000000000000000000000000000000000000000000000000000000000081565b6060620000ff7f0000000000000000000000000000000000000000000000000000000000000000620002e3565b6200012a7f0000000000000000000000000000000000000000000000000000000000000000620002e3565b620001557f0000000000000000000000000000000000000000000000000000000000000000620002e3565b60405160200162000169939291906200065a565b604051602081830303815290604052905090565b600073ffffffffffffffffffffffffffffffffffffffff841662000227576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603d60248201527f4b616e7661734d696e7461626c654552433230466163746f72793a206d75737460448201527f2070726f766964652072656d6f746520746f6b656e2061646472657373000000606482015260840160405180910390fd5b60007f00000000000000000000000000000000000000000000000000000000000000008585856040516200025b9062000438565b6200026a9493929190620006d6565b604051809103906000f08015801562000287573d6000803e3d6000fd5b5060405133815290915073ffffffffffffffffffffffffffffffffffffffff80871691908316907f70fdfb981585df6100628e4095786837f68389acbfb12f54ce7f44acc30327969060200160405180910390a3949350505050565b6060816000036200032757505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156200035757806200033e816200075f565b91506200034f9050600a83620007c9565b91506200032b565b60008167ffffffffffffffff811115620003755762000375620004e1565b6040519080825280601f01601f191660200182016040528015620003a0576020820181803683370190505b5090505b84156200043057620003b8600183620007e0565b9150620003c7600a86620007fa565b620003d490603062000811565b60f81b818381518110620003ec57620003ec6200082c565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535062000428600a86620007c9565b9450620003a4565b949350505050565b61183d806200085c83390190565b60005b838110156200046357818101518382015260200162000449565b8381111562000473576000848401525b50505050565b600081518084526200049381602086016020860162000446565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000620004da602083018462000479565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126200052257600080fd5b813567ffffffffffffffff80821115620005405762000540620004e1565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715620005895762000589620004e1565b81604052838152866020858801011115620005a357600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600060608486031215620005d957600080fd5b833573ffffffffffffffffffffffffffffffffffffffff81168114620005fe57600080fd5b9250602084013567ffffffffffffffff808211156200061c57600080fd5b6200062a8783880162000510565b935060408601359150808211156200064157600080fd5b50620006508682870162000510565b9150509250925092565b600084516200066e81846020890162000446565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551620006ac816001850160208a0162000446565b60019201918201528351620006c981600284016020880162000446565b0160020195945050505050565b600073ffffffffffffffffffffffffffffffffffffffff80871683528086166020840152506080604083015262000711608083018562000479565b828103606084015262000725818562000479565b979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820362000793576200079362000730565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082620007db57620007db6200079a565b500490565b600082821015620007f557620007f562000730565b500390565b6000826200080c576200080c6200079a565b500690565b6000821982111562000827576200082762000730565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfe6101206040523480156200001257600080fd5b506040516200183d3803806200183d83398101604081905262000035916200016d565b6000600181848460036200004a83826200028c565b5060046200005982826200028c565b50505060809290925260a05260c05250506001600160a01b0390811660e052166101005262000358565b80516001600160a01b03811681146200009b57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620000c857600080fd5b81516001600160401b0380821115620000e557620000e5620000a0565b604051601f8301601f19908116603f01168101908282118183101715620001105762000110620000a0565b816040528381526020925086838588010111156200012d57600080fd5b600091505b8382101562000151578582018301518183018401529082019062000132565b83821115620001635760008385830101525b9695505050505050565b600080600080608085870312156200018457600080fd5b6200018f8562000083565b93506200019f6020860162000083565b60408601519093506001600160401b0380821115620001bd57600080fd5b620001cb88838901620000b6565b93506060870151915080821115620001e257600080fd5b50620001f187828801620000b6565b91505092959194509250565b600181811c908216806200021257607f821691505b6020821081036200023357634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200028757600081815260208120601f850160051c81016020861015620002625750805b601f850160051c820191505b8181101562000283578281556001016200026e565b5050505b505050565b81516001600160401b03811115620002a857620002a8620000a0565b620002c081620002b98454620001fd565b8462000239565b602080601f831160018114620002f85760008415620002df5750858301515b600019600386901b1c1916600185901b17855562000283565b600085815260208120601f198616915b82811015620003295788860151825594840194600190910190840162000308565b5085821015620003485787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a05160c05160e05161010051611492620003ab600039600081816102e2015281816104d601526106960152600061014d01526000610625015260006105fc015260006105d301526114926000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c806340c10f19116100b25780639dc29fac11610081578063a9059cbb11610066578063a9059cbb14610284578063dd62ed3e14610297578063ee9a31a2146102dd57600080fd5b80639dc29fac1461025e578063a457c2d71461027157600080fd5b806340c10f191461020357806354fd4d501461021857806370a082311461022057806395d89b411461025657600080fd5b806318160ddd116100ee57806318160ddd146101bc57806323b872dd146101ce578063313ce567146101e157806339509351146101f057600080fd5b806301ffc9a714610120578063033964be1461014857806306fdde0314610194578063095ea7b3146101a9575b600080fd5b61013361012e3660046110ce565b610304565b60405190151581526020015b60405180910390f35b61016f7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161013f565b61019c6103a4565b60405161013f9190611143565b6101336101b73660046111bd565b610436565b6002545b60405190815260200161013f565b6101336101dc3660046111e7565b61044e565b6040516012815260200161013f565b6101336101fe3660046111bd565b610472565b6102166102113660046111bd565b6104be565b005b61019c6105cc565b6101c061022e366004611223565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b61019c61066f565b61021661026c3660046111bd565b61067e565b61013361027f3660046111bd565b61077b565b6101336102923660046111bd565b610832565b6101c06102a536600461123e565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b61016f7f000000000000000000000000000000000000000000000000000000000000000081565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007f30a0c5a9000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000841682148061039c57507fffffffff00000000000000000000000000000000000000000000000000000000848116908216145b949350505050565b6060600380546103b390611271565b80601f01602080910402602001604051908101604052809291908181526020018280546103df90611271565b801561042c5780601f106104015761010080835404028352916020019161042c565b820191906000526020600020905b81548152906001019060200180831161040f57829003601f168201915b5050505050905090565b600033610444818585610840565b5060019392505050565b60003361045c8582856109c0565b610467858585610a7d565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061044490829086906104b99087906112f3565b610840565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461056e5760405162461bcd60e51b815260206004820152603260248201527f4b616e7661734d696e7461626c6545524332303a206f6e6c792062726964676560448201527f2063616e206d696e7420616e64206275726e000000000000000000000000000060648201526084015b60405180910390fd5b6105788282610ce2565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885826040516105c091815260200190565b60405180910390a25050565b60606105f77f0000000000000000000000000000000000000000000000000000000000000000610de8565b6106207f0000000000000000000000000000000000000000000000000000000000000000610de8565b6106497f0000000000000000000000000000000000000000000000000000000000000000610de8565b60405160200161065b9392919061130b565b604051602081830303815290604052905090565b6060600480546103b390611271565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146107295760405162461bcd60e51b815260206004820152603260248201527f4b616e7661734d696e7461626c6545524332303a206f6e6c792062726964676560448201527f2063616e206d696e7420616e64206275726e00000000000000000000000000006064820152608401610565565b6107338282610f1d565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5826040516105c091815260200190565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152812054909190838110156108255760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152608401610565565b6104678286868403610840565b600033610444818585610a7d565b73ffffffffffffffffffffffffffffffffffffffff83166108c85760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610565565b73ffffffffffffffffffffffffffffffffffffffff82166109515760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610565565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610a775781811015610a6a5760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610565565b610a778484848403610840565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610b065760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610565565b73ffffffffffffffffffffffffffffffffffffffff8216610b8f5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610565565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610c2b5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610565565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260208190526040808220858503905591851681529081208054849290610c6f9084906112f3565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610cd591815260200190565b60405180910390a3610a77565b73ffffffffffffffffffffffffffffffffffffffff8216610d455760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610565565b8060026000828254610d5791906112f3565b909155505073ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604081208054839290610d919084906112f3565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b606081600003610e2b57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115610e555780610e3f81611381565b9150610e4e9050600a836113e8565b9150610e2f565b60008167ffffffffffffffff811115610e7057610e706113fc565b6040519080825280601f01601f191660200182016040528015610e9a576020820181803683370190505b5090505b841561039c57610eaf60018361142b565b9150610ebc600a86611442565b610ec79060306112f3565b60f81b818381518110610edc57610edc611456565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610f16600a866113e8565b9450610e9e565b73ffffffffffffffffffffffffffffffffffffffff8216610fa65760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610565565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260208190526040902054818110156110425760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152608401610565565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260208190526040812083830390556002805484929061107e90849061142b565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016109b3565b6000602082840312156110e057600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461111057600080fd5b9392505050565b60005b8381101561113257818101518382015260200161111a565b83811115610a775750506000910152565b6020815260008251806020840152611162816040850160208701611117565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146111b857600080fd5b919050565b600080604083850312156111d057600080fd5b6111d983611194565b946020939093013593505050565b6000806000606084860312156111fc57600080fd5b61120584611194565b925061121360208501611194565b9150604084013590509250925092565b60006020828403121561123557600080fd5b61111082611194565b6000806040838503121561125157600080fd5b61125a83611194565b915061126860208401611194565b90509250929050565b600181811c9082168061128557607f821691505b6020821081036112be577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115611306576113066112c4565b500190565b6000845161131d818460208901611117565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611359816001850160208a01611117565b60019201918201528351611374816002840160208801611117565b0160020195945050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036113b2576113b26112c4565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826113f7576113f76113b9565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008282101561143d5761143d6112c4565b500390565b600082611451576114516113b9565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000aa164736f6c634300080f000a",
}

// KanvasMintableERC20FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use KanvasMintableERC20FactoryMetaData.ABI instead.
var KanvasMintableERC20FactoryABI = KanvasMintableERC20FactoryMetaData.ABI

// KanvasMintableERC20FactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KanvasMintableERC20FactoryMetaData.Bin instead.
var KanvasMintableERC20FactoryBin = KanvasMintableERC20FactoryMetaData.Bin

// DeployKanvasMintableERC20Factory deploys a new Ethereum contract, binding an instance of KanvasMintableERC20Factory to it.
func DeployKanvasMintableERC20Factory(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address) (common.Address, *types.Transaction, *KanvasMintableERC20Factory, error) {
	parsed, err := KanvasMintableERC20FactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KanvasMintableERC20FactoryBin), backend, _bridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KanvasMintableERC20Factory{KanvasMintableERC20FactoryCaller: KanvasMintableERC20FactoryCaller{contract: contract}, KanvasMintableERC20FactoryTransactor: KanvasMintableERC20FactoryTransactor{contract: contract}, KanvasMintableERC20FactoryFilterer: KanvasMintableERC20FactoryFilterer{contract: contract}}, nil
}

// KanvasMintableERC20Factory is an auto generated Go binding around an Ethereum contract.
type KanvasMintableERC20Factory struct {
	KanvasMintableERC20FactoryCaller     // Read-only binding to the contract
	KanvasMintableERC20FactoryTransactor // Write-only binding to the contract
	KanvasMintableERC20FactoryFilterer   // Log filterer for contract events
}

// KanvasMintableERC20FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type KanvasMintableERC20FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KanvasMintableERC20FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KanvasMintableERC20FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KanvasMintableERC20FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KanvasMintableERC20FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KanvasMintableERC20FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KanvasMintableERC20FactorySession struct {
	Contract     *KanvasMintableERC20Factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// KanvasMintableERC20FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KanvasMintableERC20FactoryCallerSession struct {
	Contract *KanvasMintableERC20FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// KanvasMintableERC20FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KanvasMintableERC20FactoryTransactorSession struct {
	Contract     *KanvasMintableERC20FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// KanvasMintableERC20FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type KanvasMintableERC20FactoryRaw struct {
	Contract *KanvasMintableERC20Factory // Generic contract binding to access the raw methods on
}

// KanvasMintableERC20FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KanvasMintableERC20FactoryCallerRaw struct {
	Contract *KanvasMintableERC20FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// KanvasMintableERC20FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KanvasMintableERC20FactoryTransactorRaw struct {
	Contract *KanvasMintableERC20FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKanvasMintableERC20Factory creates a new instance of KanvasMintableERC20Factory, bound to a specific deployed contract.
func NewKanvasMintableERC20Factory(address common.Address, backend bind.ContractBackend) (*KanvasMintableERC20Factory, error) {
	contract, err := bindKanvasMintableERC20Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KanvasMintableERC20Factory{KanvasMintableERC20FactoryCaller: KanvasMintableERC20FactoryCaller{contract: contract}, KanvasMintableERC20FactoryTransactor: KanvasMintableERC20FactoryTransactor{contract: contract}, KanvasMintableERC20FactoryFilterer: KanvasMintableERC20FactoryFilterer{contract: contract}}, nil
}

// NewKanvasMintableERC20FactoryCaller creates a new read-only instance of KanvasMintableERC20Factory, bound to a specific deployed contract.
func NewKanvasMintableERC20FactoryCaller(address common.Address, caller bind.ContractCaller) (*KanvasMintableERC20FactoryCaller, error) {
	contract, err := bindKanvasMintableERC20Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KanvasMintableERC20FactoryCaller{contract: contract}, nil
}

// NewKanvasMintableERC20FactoryTransactor creates a new write-only instance of KanvasMintableERC20Factory, bound to a specific deployed contract.
func NewKanvasMintableERC20FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*KanvasMintableERC20FactoryTransactor, error) {
	contract, err := bindKanvasMintableERC20Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KanvasMintableERC20FactoryTransactor{contract: contract}, nil
}

// NewKanvasMintableERC20FactoryFilterer creates a new log filterer instance of KanvasMintableERC20Factory, bound to a specific deployed contract.
func NewKanvasMintableERC20FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*KanvasMintableERC20FactoryFilterer, error) {
	contract, err := bindKanvasMintableERC20Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KanvasMintableERC20FactoryFilterer{contract: contract}, nil
}

// bindKanvasMintableERC20Factory binds a generic wrapper to an already deployed contract.
func bindKanvasMintableERC20Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KanvasMintableERC20FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KanvasMintableERC20Factory *KanvasMintableERC20FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KanvasMintableERC20Factory.Contract.KanvasMintableERC20FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KanvasMintableERC20Factory *KanvasMintableERC20FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KanvasMintableERC20Factory.Contract.KanvasMintableERC20FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KanvasMintableERC20Factory *KanvasMintableERC20FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KanvasMintableERC20Factory.Contract.KanvasMintableERC20FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KanvasMintableERC20Factory *KanvasMintableERC20FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KanvasMintableERC20Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KanvasMintableERC20Factory *KanvasMintableERC20FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KanvasMintableERC20Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KanvasMintableERC20Factory *KanvasMintableERC20FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KanvasMintableERC20Factory.Contract.contract.Transact(opts, method, params...)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_KanvasMintableERC20Factory *KanvasMintableERC20FactoryCaller) BRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KanvasMintableERC20Factory.contract.Call(opts, &out, "BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_KanvasMintableERC20Factory *KanvasMintableERC20FactorySession) BRIDGE() (common.Address, error) {
	return _KanvasMintableERC20Factory.Contract.BRIDGE(&_KanvasMintableERC20Factory.CallOpts)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_KanvasMintableERC20Factory *KanvasMintableERC20FactoryCallerSession) BRIDGE() (common.Address, error) {
	return _KanvasMintableERC20Factory.Contract.BRIDGE(&_KanvasMintableERC20Factory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KanvasMintableERC20Factory *KanvasMintableERC20FactoryCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KanvasMintableERC20Factory.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KanvasMintableERC20Factory *KanvasMintableERC20FactorySession) Version() (string, error) {
	return _KanvasMintableERC20Factory.Contract.Version(&_KanvasMintableERC20Factory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KanvasMintableERC20Factory *KanvasMintableERC20FactoryCallerSession) Version() (string, error) {
	return _KanvasMintableERC20Factory.Contract.Version(&_KanvasMintableERC20Factory.CallOpts)
}

// CreateKanvasMintableERC20 is a paid mutator transaction binding the contract method 0x8c2378ba.
//
// Solidity: function createKanvasMintableERC20(address _remoteToken, string _name, string _symbol) returns(address)
func (_KanvasMintableERC20Factory *KanvasMintableERC20FactoryTransactor) CreateKanvasMintableERC20(opts *bind.TransactOpts, _remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _KanvasMintableERC20Factory.contract.Transact(opts, "createKanvasMintableERC20", _remoteToken, _name, _symbol)
}

// CreateKanvasMintableERC20 is a paid mutator transaction binding the contract method 0x8c2378ba.
//
// Solidity: function createKanvasMintableERC20(address _remoteToken, string _name, string _symbol) returns(address)
func (_KanvasMintableERC20Factory *KanvasMintableERC20FactorySession) CreateKanvasMintableERC20(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _KanvasMintableERC20Factory.Contract.CreateKanvasMintableERC20(&_KanvasMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// CreateKanvasMintableERC20 is a paid mutator transaction binding the contract method 0x8c2378ba.
//
// Solidity: function createKanvasMintableERC20(address _remoteToken, string _name, string _symbol) returns(address)
func (_KanvasMintableERC20Factory *KanvasMintableERC20FactoryTransactorSession) CreateKanvasMintableERC20(_remoteToken common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _KanvasMintableERC20Factory.Contract.CreateKanvasMintableERC20(&_KanvasMintableERC20Factory.TransactOpts, _remoteToken, _name, _symbol)
}

// KanvasMintableERC20FactoryKanvasMintableERC20CreatedIterator is returned from FilterKanvasMintableERC20Created and is used to iterate over the raw logs and unpacked data for KanvasMintableERC20Created events raised by the KanvasMintableERC20Factory contract.
type KanvasMintableERC20FactoryKanvasMintableERC20CreatedIterator struct {
	Event *KanvasMintableERC20FactoryKanvasMintableERC20Created // Event containing the contract specifics and raw log

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
func (it *KanvasMintableERC20FactoryKanvasMintableERC20CreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KanvasMintableERC20FactoryKanvasMintableERC20Created)
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
		it.Event = new(KanvasMintableERC20FactoryKanvasMintableERC20Created)
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
func (it *KanvasMintableERC20FactoryKanvasMintableERC20CreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KanvasMintableERC20FactoryKanvasMintableERC20CreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KanvasMintableERC20FactoryKanvasMintableERC20Created represents a KanvasMintableERC20Created event raised by the KanvasMintableERC20Factory contract.
type KanvasMintableERC20FactoryKanvasMintableERC20Created struct {
	LocalToken  common.Address
	RemoteToken common.Address
	Deployer    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterKanvasMintableERC20Created is a free log retrieval operation binding the contract event 0x70fdfb981585df6100628e4095786837f68389acbfb12f54ce7f44acc3032796.
//
// Solidity: event KanvasMintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer)
func (_KanvasMintableERC20Factory *KanvasMintableERC20FactoryFilterer) FilterKanvasMintableERC20Created(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address) (*KanvasMintableERC20FactoryKanvasMintableERC20CreatedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _KanvasMintableERC20Factory.contract.FilterLogs(opts, "KanvasMintableERC20Created", localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &KanvasMintableERC20FactoryKanvasMintableERC20CreatedIterator{contract: _KanvasMintableERC20Factory.contract, event: "KanvasMintableERC20Created", logs: logs, sub: sub}, nil
}

// WatchKanvasMintableERC20Created is a free log subscription operation binding the contract event 0x70fdfb981585df6100628e4095786837f68389acbfb12f54ce7f44acc3032796.
//
// Solidity: event KanvasMintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer)
func (_KanvasMintableERC20Factory *KanvasMintableERC20FactoryFilterer) WatchKanvasMintableERC20Created(opts *bind.WatchOpts, sink chan<- *KanvasMintableERC20FactoryKanvasMintableERC20Created, localToken []common.Address, remoteToken []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}

	logs, sub, err := _KanvasMintableERC20Factory.contract.WatchLogs(opts, "KanvasMintableERC20Created", localTokenRule, remoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KanvasMintableERC20FactoryKanvasMintableERC20Created)
				if err := _KanvasMintableERC20Factory.contract.UnpackLog(event, "KanvasMintableERC20Created", log); err != nil {
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

// ParseKanvasMintableERC20Created is a log parse operation binding the contract event 0x70fdfb981585df6100628e4095786837f68389acbfb12f54ce7f44acc3032796.
//
// Solidity: event KanvasMintableERC20Created(address indexed localToken, address indexed remoteToken, address deployer)
func (_KanvasMintableERC20Factory *KanvasMintableERC20FactoryFilterer) ParseKanvasMintableERC20Created(log types.Log) (*KanvasMintableERC20FactoryKanvasMintableERC20Created, error) {
	event := new(KanvasMintableERC20FactoryKanvasMintableERC20Created)
	if err := _KanvasMintableERC20Factory.contract.UnpackLog(event, "KanvasMintableERC20Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
