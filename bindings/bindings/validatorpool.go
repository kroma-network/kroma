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

// ValidatorPoolMetaData contains all meta data concerning the ValidatorPool contract.
var ValidatorPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractL2OutputOracle\",\"name\":\"_l2OutputOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedValidator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minBondAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"outputIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"BondIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"outputIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"expiresAt\",\"type\":\"uint128\"}],\"name\":\"Bonded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"outputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"Unbonded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"L2_ORACLE\",\"outputs\":[{\"internalType\":\"contractL2OutputOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BOND_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRUSTED_VALIDATOR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_outputIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_expiresAt\",\"type\":\"uint128\"}],\"name\":\"createBond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_outputIndex\",\"type\":\"uint256\"}],\"name\":\"getBond\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"expiresAt\",\"type\":\"uint128\"}],\"internalType\":\"structTypes.Bond\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_outputIndex\",\"type\":\"uint256\"}],\"name\":\"increaseBond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextValidator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unbond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101406040523480156200001257600080fd5b5060405162001e3138038062001e318339810160408190526200003591620001bd565b60006080819052600160a05260c0526001600160a01b0380841660e0528216610100526101208190526200006862000071565b50505062000205565b600054610100900460ff1615808015620000925750600054600160ff909116105b80620000c25750620000af306200019860201b620010751760201c565b158015620000c2575060005460ff166001145b6200012a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b6000805460ff1916600117905580156200014e576000805461ff0019166101001790555b801562000195576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b6001600160a01b03163b151590565b6001600160a01b03811681146200019557600080fd5b600080600060608486031215620001d357600080fd5b8351620001e081620001a7565b6020850151909350620001f381620001a7565b80925050604084015190509250925092565b60805160a05160c05160e0516101005161012051611b9e620002936000396000818161029501528181610a2801528181611154015261168e0152600081816101b2015261067401526000818160fe01528181610441015281816104ef0152818161097901528181610be601526114ae015260006106f2015260006106c9015260006106a00152611b9e6000f3fe6080604052600436106100e75760003560e01c806370a082311161008a578063d0e30db011610059578063d0e30db0146102b7578063d8fe7642146102bf578063da3893f01461030f578063facd743b1461032f57600080fd5b806370a082311461020b5780638129fc1c1461024e57806396946f75146102635780639fbc4a5f1461028357600080fd5b80633a549046116100c65780633a5490461461018b5780633ee4d4a3146101a057806354fd4d50146101d45780635df6a6bc146101f657600080fd5b80621c2ff6146100ec5780630f43a6771461014a5780632e1a7d4d14610169575b600080fd5b3480156100f857600080fd5b506101207f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561015657600080fd5b506036545b604051908152602001610141565b34801561017557600080fd5b5061018961018436600461175c565b61035f565b005b34801561019757600080fd5b5061012061043c565b3480156101ac57600080fd5b506101207f000000000000000000000000000000000000000000000000000000000000000081565b3480156101e057600080fd5b506101e9610699565b60405161014191906117a5565b34801561020257600080fd5b5061018961073c565b34801561021757600080fd5b5061015b610226366004611818565b73ffffffffffffffffffffffffffffffffffffffff1660009081526033602052604090205490565b34801561025a57600080fd5b506101896107d8565b34801561026f57600080fd5b5061018961027e36600461185a565b610961565b34801561028f57600080fd5b5061015b7f000000000000000000000000000000000000000000000000000000000000000081565b610189610d0f565b3480156102cb57600080fd5b506102df6102da36600461175c565b610d1b565b6040805182516fffffffffffffffffffffffffffffffff9081168252602093840151169281019290925201610141565b34801561031b57600080fd5b5061018961032a36600461189c565b610e54565b34801561033b57600080fd5b5061034f61034a366004611818565b610fd7565b6040519015158152602001610141565b6002600154036103d0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064015b60405180910390fd5b60026001556103df3382611091565b604051600090339083908381818185875af1925050503d8060008114610421576040519150601f19603f3d011682016040523d82523d6000602084013e610426565b606091505b505090508061043457600080fd5b505060018055565b6000807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636abcf5636040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ce91906118c8565b9050801561067257600073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663a25ae55761051f600185611910565b6040518263ffffffff1660e01b815260040161053d91815260200190565b608060405180830381865afa15801561055a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057e9190611956565b603654602082810151835160408086015160608088015183519687018890529286019490945291831b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001692840192909252608090811b7fffffffffffffffffffffffffffffffff00000000000000000000000000000000908116607485015291901b1660848201529192506000916094016040516020818303038152906040528051906020012060001c6106339190611a28565b90506036818154811061064857610648611a3c565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16949350505050565b7f000000000000000000000000000000000000000000000000000000000000000091505090565b60606106c47f0000000000000000000000000000000000000000000000000000000000000000611339565b6106ed7f0000000000000000000000000000000000000000000000000000000000000000611339565b6107167f0000000000000000000000000000000000000000000000000000000000000000611339565b60405160200161072893929190611a6b565b604051602081830303815290604052905090565b6000610746611476565b9050806107d5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f56616c696461746f72506f6f6c3a206e6f20626f6e6420746861742063616e2060448201527f626520756e626f6e64000000000000000000000000000000000000000000000060648201526084016103c7565b50565b600054610100900460ff16158080156107f85750600054600160ff909116105b806108125750303b158015610812575060005460ff166001145b61089e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016103c7565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156108fc57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b80156107d557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610a26576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f56616c696461746f72506f6f6c3a2073656e646572206973206e6f74204c324f60448201527f75747075744f7261636c6500000000000000000000000000000000000000000060648201526084016103c7565b7f0000000000000000000000000000000000000000000000000000000000000000826fffffffffffffffffffffffffffffffff161015610ae8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f56616c696461746f72506f6f6c3a2074686520626f6e6420616d6f756e74206960448201527f7320746f6f20736d616c6c00000000000000000000000000000000000000000060648201526084016103c7565b6000838152603460205260409020805470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff1615610bab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603c60248201527f56616c696461746f72506f6f6c3a20626f6e64206f662074686520676976656e60448201527f206f757470757420696e64657820616c7265616479206578697374730000000060648201526084016103c7565b610bb3611476565b506040517fb0ea09a8000000000000000000000000000000000000000000000000000000008152600481018590526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063b0ea09a890602401602060405180830381865afa158015610c42573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c669190611ae1565b9050610c8481856fffffffffffffffffffffffffffffffff16611091565b6fffffffffffffffffffffffffffffffff84811670010000000000000000000000000000000091851691820281178455604080519182526020820192909252869173ffffffffffffffffffffffffffffffffffffffff8416917f5ca130257b8f76f72ad2965efcbe166f3918d820e4a49956e70081ea311f97c4910160405180910390a35050505050565b610d193334611659565b565b6040805180820190915260008082526020820152600082815260346020526040902080546fffffffffffffffffffffffffffffffff1615801590610d855750805470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff1615155b610e11576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f56616c696461746f72506f6f6c3a2074686520626f6e6420646f6573206e6f7460448201527f206578697374000000000000000000000000000000000000000000000000000060648201526084016103c7565b6040805180820190915290546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000090910416602082015292915050565b6000818152603460205260409020805470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16610f16576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f56616c696461746f72506f6f6c3a2074686520626f6e6420646f6573206e6f7460448201527f206578697374000000000000000000000000000000000000000000000000000060648201526084016103c7565b80546fffffffffffffffffffffffffffffffff16610f348482611091565b81547fffffffffffffffffffffffffffffffff0000000000000000000000000000000016600182901b6ffffffffffffffffffffffffffffffffe161782556040516fffffffffffffffffffffffffffffffff82168152839073ffffffffffffffffffffffffffffffffffffffff8616907f0d0a53301770c0275802b487151539531ef1f7f94d361e97a561ebe8233ab80c9060200160405180910390a350505050565b6036546000908103610feb57506000919050565b73ffffffffffffffffffffffffffffffffffffffff821661100e57506000919050565b73ffffffffffffffffffffffffffffffffffffffff8216600081815260376020526040902054603680549192918390811061104b5761104b611a3c565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16149392505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b73ffffffffffffffffffffffffffffffffffffffff821660009081526033602052604090205481811015611146576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f56616c696461746f72506f6f6c3a20696e73756666696369656e742062616c6160448201527f6e6365730000000000000000000000000000000000000000000000000000000060648201526084016103c7565b6111508282611910565b90507f000000000000000000000000000000000000000000000000000000000000000081108015611185575061118583610fd7565b1561130c5760365460009061119c90600190611910565b9050801561127b5773ffffffffffffffffffffffffffffffffffffffff841660009081526037602052604081205460368054919291849081106111e1576111e1611a3c565b6000918252602090912001546036805473ffffffffffffffffffffffffffffffffffffffff909216925082918490811061121d5761121d611a3c565b600091825260208083209190910180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff948516179055929091168152603790915260409020555b73ffffffffffffffffffffffffffffffffffffffff841660009081526037602052604081205560368054806112b2576112b2611afe565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055019055505b73ffffffffffffffffffffffffffffffffffffffff90921660009081526033602052604090209190915550565b60608160000361137c57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156113a6578061139081611b2d565b915061139f9050600a83611b65565b9150611380565b60008167ffffffffffffffff8111156113c1576113c1611927565b6040519080825280601f01601f1916602001820160405280156113eb576020820181803683370190505b5090505b841561146e57611400600183611910565b915061140d600a86611a28565b611418906030611b79565b60f81b81838151811061142d5761142d611a3c565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350611467600a86611b65565b94506113ef565b949350505050565b6035546040517fb0ea09a8000000000000000000000000000000000000000000000000000000008152600481018290526000919082907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063b0ea09a890602401602060405180830381865afa15801561150a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061152e9190611ae1565b600083815260346020526040902080549192509070010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff164210801590611589575080546fffffffffffffffffffffffffffffffff1615155b1561164f5780546115ad9083906fffffffffffffffffffffffffffffffff16611659565b6035546115bb906001611b79565b60355580546040516fffffffffffffffffffffffffffffffff909116815273ffffffffffffffffffffffffffffffffffffffff83169084907f7047a0fb8bfae78c0ebbd4117437945bb85240453235ac4fd2e55712eb5bf0c39060200160405180910390a380547fffffffffffffffffffffffffffffffff0000000000000000000000000000000016905550600192915050565b6000935050505090565b73ffffffffffffffffffffffffffffffffffffffff821660009081526033602052604081205461168a908390611b79565b90507f000000000000000000000000000000000000000000000000000000000000000081101580156116c257506116c083610fd7565b155b1561130c576036805473ffffffffffffffffffffffffffffffffffffffff949094166000818152603760209081526040808320889055600188019094557f4a11f94e20a93c79f6ec743a1954ec4fc2c08429ae2122118bf234b2185c81b890960180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690921790915560339094529092209190915550565b60006020828403121561176e57600080fd5b5035919050565b60005b83811015611790578181015183820152602001611778565b8381111561179f576000848401525b50505050565b60208152600082518060208401526117c4816040850160208701611775565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b73ffffffffffffffffffffffffffffffffffffffff811681146107d557600080fd5b60006020828403121561182a57600080fd5b8135611835816117f6565b9392505050565b6fffffffffffffffffffffffffffffffff811681146107d557600080fd5b60008060006060848603121561186f57600080fd5b8335925060208401356118818161183c565b915060408401356118918161183c565b809150509250925092565b600080604083850312156118af57600080fd5b82356118ba816117f6565b946020939093013593505050565b6000602082840312156118da57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015611922576119226118e1565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006080828403121561196857600080fd5b6040516080810181811067ffffffffffffffff821117156119b2577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405282516119c0816117f6565b81526020838101519082015260408301516119da8161183c565b604082015260608301516119ed8161183c565b60608201529392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082611a3757611a376119f9565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008451611a7d818460208901611775565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611ab9816001850160208a01611775565b60019201918201528351611ad4816002840160208801611775565b0160020195945050505050565b600060208284031215611af357600080fd5b8151611835816117f6565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611b5e57611b5e6118e1565b5060010190565b600082611b7457611b746119f9565b500490565b60008219821115611b8c57611b8c6118e1565b50019056fea164736f6c634300080f000a",
}

// ValidatorPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorPoolMetaData.ABI instead.
var ValidatorPoolABI = ValidatorPoolMetaData.ABI

// ValidatorPoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorPoolMetaData.Bin instead.
var ValidatorPoolBin = ValidatorPoolMetaData.Bin

// DeployValidatorPool deploys a new Ethereum contract, binding an instance of ValidatorPool to it.
func DeployValidatorPool(auth *bind.TransactOpts, backend bind.ContractBackend, _l2OutputOracle common.Address, _trustedValidator common.Address, _minBondAmount *big.Int) (common.Address, *types.Transaction, *ValidatorPool, error) {
	parsed, err := ValidatorPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorPoolBin), backend, _l2OutputOracle, _trustedValidator, _minBondAmount)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorPool{ValidatorPoolCaller: ValidatorPoolCaller{contract: contract}, ValidatorPoolTransactor: ValidatorPoolTransactor{contract: contract}, ValidatorPoolFilterer: ValidatorPoolFilterer{contract: contract}}, nil
}

// ValidatorPool is an auto generated Go binding around an Ethereum contract.
type ValidatorPool struct {
	ValidatorPoolCaller     // Read-only binding to the contract
	ValidatorPoolTransactor // Write-only binding to the contract
	ValidatorPoolFilterer   // Log filterer for contract events
}

// ValidatorPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorPoolSession struct {
	Contract     *ValidatorPool    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorPoolCallerSession struct {
	Contract *ValidatorPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ValidatorPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorPoolTransactorSession struct {
	Contract     *ValidatorPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ValidatorPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorPoolRaw struct {
	Contract *ValidatorPool // Generic contract binding to access the raw methods on
}

// ValidatorPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorPoolCallerRaw struct {
	Contract *ValidatorPoolCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorPoolTransactorRaw struct {
	Contract *ValidatorPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorPool creates a new instance of ValidatorPool, bound to a specific deployed contract.
func NewValidatorPool(address common.Address, backend bind.ContractBackend) (*ValidatorPool, error) {
	contract, err := bindValidatorPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorPool{ValidatorPoolCaller: ValidatorPoolCaller{contract: contract}, ValidatorPoolTransactor: ValidatorPoolTransactor{contract: contract}, ValidatorPoolFilterer: ValidatorPoolFilterer{contract: contract}}, nil
}

// NewValidatorPoolCaller creates a new read-only instance of ValidatorPool, bound to a specific deployed contract.
func NewValidatorPoolCaller(address common.Address, caller bind.ContractCaller) (*ValidatorPoolCaller, error) {
	contract, err := bindValidatorPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorPoolCaller{contract: contract}, nil
}

// NewValidatorPoolTransactor creates a new write-only instance of ValidatorPool, bound to a specific deployed contract.
func NewValidatorPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorPoolTransactor, error) {
	contract, err := bindValidatorPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorPoolTransactor{contract: contract}, nil
}

// NewValidatorPoolFilterer creates a new log filterer instance of ValidatorPool, bound to a specific deployed contract.
func NewValidatorPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorPoolFilterer, error) {
	contract, err := bindValidatorPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorPoolFilterer{contract: contract}, nil
}

// bindValidatorPool binds a generic wrapper to an already deployed contract.
func bindValidatorPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorPool *ValidatorPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorPool.Contract.ValidatorPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorPool *ValidatorPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPool.Contract.ValidatorPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorPool *ValidatorPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorPool.Contract.ValidatorPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorPool *ValidatorPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorPool *ValidatorPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorPool *ValidatorPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorPool.Contract.contract.Transact(opts, method, params...)
}

// L2ORACLE is a free data retrieval call binding the contract method 0x001c2ff6.
//
// Solidity: function L2_ORACLE() view returns(address)
func (_ValidatorPool *ValidatorPoolCaller) L2ORACLE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "L2_ORACLE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2ORACLE is a free data retrieval call binding the contract method 0x001c2ff6.
//
// Solidity: function L2_ORACLE() view returns(address)
func (_ValidatorPool *ValidatorPoolSession) L2ORACLE() (common.Address, error) {
	return _ValidatorPool.Contract.L2ORACLE(&_ValidatorPool.CallOpts)
}

// L2ORACLE is a free data retrieval call binding the contract method 0x001c2ff6.
//
// Solidity: function L2_ORACLE() view returns(address)
func (_ValidatorPool *ValidatorPoolCallerSession) L2ORACLE() (common.Address, error) {
	return _ValidatorPool.Contract.L2ORACLE(&_ValidatorPool.CallOpts)
}

// MINBONDAMOUNT is a free data retrieval call binding the contract method 0x9fbc4a5f.
//
// Solidity: function MIN_BOND_AMOUNT() view returns(uint256)
func (_ValidatorPool *ValidatorPoolCaller) MINBONDAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "MIN_BOND_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBONDAMOUNT is a free data retrieval call binding the contract method 0x9fbc4a5f.
//
// Solidity: function MIN_BOND_AMOUNT() view returns(uint256)
func (_ValidatorPool *ValidatorPoolSession) MINBONDAMOUNT() (*big.Int, error) {
	return _ValidatorPool.Contract.MINBONDAMOUNT(&_ValidatorPool.CallOpts)
}

// MINBONDAMOUNT is a free data retrieval call binding the contract method 0x9fbc4a5f.
//
// Solidity: function MIN_BOND_AMOUNT() view returns(uint256)
func (_ValidatorPool *ValidatorPoolCallerSession) MINBONDAMOUNT() (*big.Int, error) {
	return _ValidatorPool.Contract.MINBONDAMOUNT(&_ValidatorPool.CallOpts)
}

// TRUSTEDVALIDATOR is a free data retrieval call binding the contract method 0x3ee4d4a3.
//
// Solidity: function TRUSTED_VALIDATOR() view returns(address)
func (_ValidatorPool *ValidatorPoolCaller) TRUSTEDVALIDATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "TRUSTED_VALIDATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TRUSTEDVALIDATOR is a free data retrieval call binding the contract method 0x3ee4d4a3.
//
// Solidity: function TRUSTED_VALIDATOR() view returns(address)
func (_ValidatorPool *ValidatorPoolSession) TRUSTEDVALIDATOR() (common.Address, error) {
	return _ValidatorPool.Contract.TRUSTEDVALIDATOR(&_ValidatorPool.CallOpts)
}

// TRUSTEDVALIDATOR is a free data retrieval call binding the contract method 0x3ee4d4a3.
//
// Solidity: function TRUSTED_VALIDATOR() view returns(address)
func (_ValidatorPool *ValidatorPoolCallerSession) TRUSTEDVALIDATOR() (common.Address, error) {
	return _ValidatorPool.Contract.TRUSTEDVALIDATOR(&_ValidatorPool.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _addr) view returns(uint256)
func (_ValidatorPool *ValidatorPoolCaller) BalanceOf(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "balanceOf", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _addr) view returns(uint256)
func (_ValidatorPool *ValidatorPoolSession) BalanceOf(_addr common.Address) (*big.Int, error) {
	return _ValidatorPool.Contract.BalanceOf(&_ValidatorPool.CallOpts, _addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _addr) view returns(uint256)
func (_ValidatorPool *ValidatorPoolCallerSession) BalanceOf(_addr common.Address) (*big.Int, error) {
	return _ValidatorPool.Contract.BalanceOf(&_ValidatorPool.CallOpts, _addr)
}

// GetBond is a free data retrieval call binding the contract method 0xd8fe7642.
//
// Solidity: function getBond(uint256 _outputIndex) view returns((uint128,uint128))
func (_ValidatorPool *ValidatorPoolCaller) GetBond(opts *bind.CallOpts, _outputIndex *big.Int) (TypesBond, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "getBond", _outputIndex)

	if err != nil {
		return *new(TypesBond), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesBond)).(*TypesBond)

	return out0, err

}

// GetBond is a free data retrieval call binding the contract method 0xd8fe7642.
//
// Solidity: function getBond(uint256 _outputIndex) view returns((uint128,uint128))
func (_ValidatorPool *ValidatorPoolSession) GetBond(_outputIndex *big.Int) (TypesBond, error) {
	return _ValidatorPool.Contract.GetBond(&_ValidatorPool.CallOpts, _outputIndex)
}

// GetBond is a free data retrieval call binding the contract method 0xd8fe7642.
//
// Solidity: function getBond(uint256 _outputIndex) view returns((uint128,uint128))
func (_ValidatorPool *ValidatorPoolCallerSession) GetBond(_outputIndex *big.Int) (TypesBond, error) {
	return _ValidatorPool.Contract.GetBond(&_ValidatorPool.CallOpts, _outputIndex)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _addr) view returns(bool)
func (_ValidatorPool *ValidatorPoolCaller) IsValidator(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "isValidator", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _addr) view returns(bool)
func (_ValidatorPool *ValidatorPoolSession) IsValidator(_addr common.Address) (bool, error) {
	return _ValidatorPool.Contract.IsValidator(&_ValidatorPool.CallOpts, _addr)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _addr) view returns(bool)
func (_ValidatorPool *ValidatorPoolCallerSession) IsValidator(_addr common.Address) (bool, error) {
	return _ValidatorPool.Contract.IsValidator(&_ValidatorPool.CallOpts, _addr)
}

// NextValidator is a free data retrieval call binding the contract method 0x3a549046.
//
// Solidity: function nextValidator() view returns(address)
func (_ValidatorPool *ValidatorPoolCaller) NextValidator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "nextValidator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NextValidator is a free data retrieval call binding the contract method 0x3a549046.
//
// Solidity: function nextValidator() view returns(address)
func (_ValidatorPool *ValidatorPoolSession) NextValidator() (common.Address, error) {
	return _ValidatorPool.Contract.NextValidator(&_ValidatorPool.CallOpts)
}

// NextValidator is a free data retrieval call binding the contract method 0x3a549046.
//
// Solidity: function nextValidator() view returns(address)
func (_ValidatorPool *ValidatorPoolCallerSession) NextValidator() (common.Address, error) {
	return _ValidatorPool.Contract.NextValidator(&_ValidatorPool.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_ValidatorPool *ValidatorPoolCaller) ValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "validatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_ValidatorPool *ValidatorPoolSession) ValidatorCount() (*big.Int, error) {
	return _ValidatorPool.Contract.ValidatorCount(&_ValidatorPool.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_ValidatorPool *ValidatorPoolCallerSession) ValidatorCount() (*big.Int, error) {
	return _ValidatorPool.Contract.ValidatorCount(&_ValidatorPool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ValidatorPool *ValidatorPoolCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ValidatorPool *ValidatorPoolSession) Version() (string, error) {
	return _ValidatorPool.Contract.Version(&_ValidatorPool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ValidatorPool *ValidatorPoolCallerSession) Version() (string, error) {
	return _ValidatorPool.Contract.Version(&_ValidatorPool.CallOpts)
}

// CreateBond is a paid mutator transaction binding the contract method 0x96946f75.
//
// Solidity: function createBond(uint256 _outputIndex, uint128 _amount, uint128 _expiresAt) returns()
func (_ValidatorPool *ValidatorPoolTransactor) CreateBond(opts *bind.TransactOpts, _outputIndex *big.Int, _amount *big.Int, _expiresAt *big.Int) (*types.Transaction, error) {
	return _ValidatorPool.contract.Transact(opts, "createBond", _outputIndex, _amount, _expiresAt)
}

// CreateBond is a paid mutator transaction binding the contract method 0x96946f75.
//
// Solidity: function createBond(uint256 _outputIndex, uint128 _amount, uint128 _expiresAt) returns()
func (_ValidatorPool *ValidatorPoolSession) CreateBond(_outputIndex *big.Int, _amount *big.Int, _expiresAt *big.Int) (*types.Transaction, error) {
	return _ValidatorPool.Contract.CreateBond(&_ValidatorPool.TransactOpts, _outputIndex, _amount, _expiresAt)
}

// CreateBond is a paid mutator transaction binding the contract method 0x96946f75.
//
// Solidity: function createBond(uint256 _outputIndex, uint128 _amount, uint128 _expiresAt) returns()
func (_ValidatorPool *ValidatorPoolTransactorSession) CreateBond(_outputIndex *big.Int, _amount *big.Int, _expiresAt *big.Int) (*types.Transaction, error) {
	return _ValidatorPool.Contract.CreateBond(&_ValidatorPool.TransactOpts, _outputIndex, _amount, _expiresAt)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ValidatorPool *ValidatorPoolTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPool.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ValidatorPool *ValidatorPoolSession) Deposit() (*types.Transaction, error) {
	return _ValidatorPool.Contract.Deposit(&_ValidatorPool.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_ValidatorPool *ValidatorPoolTransactorSession) Deposit() (*types.Transaction, error) {
	return _ValidatorPool.Contract.Deposit(&_ValidatorPool.TransactOpts)
}

// IncreaseBond is a paid mutator transaction binding the contract method 0xda3893f0.
//
// Solidity: function increaseBond(address _challenger, uint256 _outputIndex) returns()
func (_ValidatorPool *ValidatorPoolTransactor) IncreaseBond(opts *bind.TransactOpts, _challenger common.Address, _outputIndex *big.Int) (*types.Transaction, error) {
	return _ValidatorPool.contract.Transact(opts, "increaseBond", _challenger, _outputIndex)
}

// IncreaseBond is a paid mutator transaction binding the contract method 0xda3893f0.
//
// Solidity: function increaseBond(address _challenger, uint256 _outputIndex) returns()
func (_ValidatorPool *ValidatorPoolSession) IncreaseBond(_challenger common.Address, _outputIndex *big.Int) (*types.Transaction, error) {
	return _ValidatorPool.Contract.IncreaseBond(&_ValidatorPool.TransactOpts, _challenger, _outputIndex)
}

// IncreaseBond is a paid mutator transaction binding the contract method 0xda3893f0.
//
// Solidity: function increaseBond(address _challenger, uint256 _outputIndex) returns()
func (_ValidatorPool *ValidatorPoolTransactorSession) IncreaseBond(_challenger common.Address, _outputIndex *big.Int) (*types.Transaction, error) {
	return _ValidatorPool.Contract.IncreaseBond(&_ValidatorPool.TransactOpts, _challenger, _outputIndex)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ValidatorPool *ValidatorPoolTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPool.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ValidatorPool *ValidatorPoolSession) Initialize() (*types.Transaction, error) {
	return _ValidatorPool.Contract.Initialize(&_ValidatorPool.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ValidatorPool *ValidatorPoolTransactorSession) Initialize() (*types.Transaction, error) {
	return _ValidatorPool.Contract.Initialize(&_ValidatorPool.TransactOpts)
}

// Unbond is a paid mutator transaction binding the contract method 0x5df6a6bc.
//
// Solidity: function unbond() returns()
func (_ValidatorPool *ValidatorPoolTransactor) Unbond(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorPool.contract.Transact(opts, "unbond")
}

// Unbond is a paid mutator transaction binding the contract method 0x5df6a6bc.
//
// Solidity: function unbond() returns()
func (_ValidatorPool *ValidatorPoolSession) Unbond() (*types.Transaction, error) {
	return _ValidatorPool.Contract.Unbond(&_ValidatorPool.TransactOpts)
}

// Unbond is a paid mutator transaction binding the contract method 0x5df6a6bc.
//
// Solidity: function unbond() returns()
func (_ValidatorPool *ValidatorPoolTransactorSession) Unbond() (*types.Transaction, error) {
	return _ValidatorPool.Contract.Unbond(&_ValidatorPool.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_ValidatorPool *ValidatorPoolTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ValidatorPool.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_ValidatorPool *ValidatorPoolSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _ValidatorPool.Contract.Withdraw(&_ValidatorPool.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_ValidatorPool *ValidatorPoolTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _ValidatorPool.Contract.Withdraw(&_ValidatorPool.TransactOpts, _amount)
}

// ValidatorPoolBondIncreasedIterator is returned from FilterBondIncreased and is used to iterate over the raw logs and unpacked data for BondIncreased events raised by the ValidatorPool contract.
type ValidatorPoolBondIncreasedIterator struct {
	Event *ValidatorPoolBondIncreased // Event containing the contract specifics and raw log

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
func (it *ValidatorPoolBondIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorPoolBondIncreased)
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
		it.Event = new(ValidatorPoolBondIncreased)
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
func (it *ValidatorPoolBondIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorPoolBondIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorPoolBondIncreased represents a BondIncreased event raised by the ValidatorPool contract.
type ValidatorPoolBondIncreased struct {
	Challenger  common.Address
	OutputIndex *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBondIncreased is a free log retrieval operation binding the contract event 0x0d0a53301770c0275802b487151539531ef1f7f94d361e97a561ebe8233ab80c.
//
// Solidity: event BondIncreased(address indexed challenger, uint256 indexed outputIndex, uint128 amount)
func (_ValidatorPool *ValidatorPoolFilterer) FilterBondIncreased(opts *bind.FilterOpts, challenger []common.Address, outputIndex []*big.Int) (*ValidatorPoolBondIncreasedIterator, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}
	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}

	logs, sub, err := _ValidatorPool.contract.FilterLogs(opts, "BondIncreased", challengerRule, outputIndexRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorPoolBondIncreasedIterator{contract: _ValidatorPool.contract, event: "BondIncreased", logs: logs, sub: sub}, nil
}

// WatchBondIncreased is a free log subscription operation binding the contract event 0x0d0a53301770c0275802b487151539531ef1f7f94d361e97a561ebe8233ab80c.
//
// Solidity: event BondIncreased(address indexed challenger, uint256 indexed outputIndex, uint128 amount)
func (_ValidatorPool *ValidatorPoolFilterer) WatchBondIncreased(opts *bind.WatchOpts, sink chan<- *ValidatorPoolBondIncreased, challenger []common.Address, outputIndex []*big.Int) (event.Subscription, error) {

	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}
	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}

	logs, sub, err := _ValidatorPool.contract.WatchLogs(opts, "BondIncreased", challengerRule, outputIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorPoolBondIncreased)
				if err := _ValidatorPool.contract.UnpackLog(event, "BondIncreased", log); err != nil {
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

// ParseBondIncreased is a log parse operation binding the contract event 0x0d0a53301770c0275802b487151539531ef1f7f94d361e97a561ebe8233ab80c.
//
// Solidity: event BondIncreased(address indexed challenger, uint256 indexed outputIndex, uint128 amount)
func (_ValidatorPool *ValidatorPoolFilterer) ParseBondIncreased(log types.Log) (*ValidatorPoolBondIncreased, error) {
	event := new(ValidatorPoolBondIncreased)
	if err := _ValidatorPool.contract.UnpackLog(event, "BondIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorPoolBondedIterator is returned from FilterBonded and is used to iterate over the raw logs and unpacked data for Bonded events raised by the ValidatorPool contract.
type ValidatorPoolBondedIterator struct {
	Event *ValidatorPoolBonded // Event containing the contract specifics and raw log

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
func (it *ValidatorPoolBondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorPoolBonded)
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
		it.Event = new(ValidatorPoolBonded)
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
func (it *ValidatorPoolBondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorPoolBondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorPoolBonded represents a Bonded event raised by the ValidatorPool contract.
type ValidatorPoolBonded struct {
	Submitter   common.Address
	OutputIndex *big.Int
	Amount      *big.Int
	ExpiresAt   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBonded is a free log retrieval operation binding the contract event 0x5ca130257b8f76f72ad2965efcbe166f3918d820e4a49956e70081ea311f97c4.
//
// Solidity: event Bonded(address indexed submitter, uint256 indexed outputIndex, uint128 amount, uint128 expiresAt)
func (_ValidatorPool *ValidatorPoolFilterer) FilterBonded(opts *bind.FilterOpts, submitter []common.Address, outputIndex []*big.Int) (*ValidatorPoolBondedIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}

	logs, sub, err := _ValidatorPool.contract.FilterLogs(opts, "Bonded", submitterRule, outputIndexRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorPoolBondedIterator{contract: _ValidatorPool.contract, event: "Bonded", logs: logs, sub: sub}, nil
}

// WatchBonded is a free log subscription operation binding the contract event 0x5ca130257b8f76f72ad2965efcbe166f3918d820e4a49956e70081ea311f97c4.
//
// Solidity: event Bonded(address indexed submitter, uint256 indexed outputIndex, uint128 amount, uint128 expiresAt)
func (_ValidatorPool *ValidatorPoolFilterer) WatchBonded(opts *bind.WatchOpts, sink chan<- *ValidatorPoolBonded, submitter []common.Address, outputIndex []*big.Int) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}

	logs, sub, err := _ValidatorPool.contract.WatchLogs(opts, "Bonded", submitterRule, outputIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorPoolBonded)
				if err := _ValidatorPool.contract.UnpackLog(event, "Bonded", log); err != nil {
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

// ParseBonded is a log parse operation binding the contract event 0x5ca130257b8f76f72ad2965efcbe166f3918d820e4a49956e70081ea311f97c4.
//
// Solidity: event Bonded(address indexed submitter, uint256 indexed outputIndex, uint128 amount, uint128 expiresAt)
func (_ValidatorPool *ValidatorPoolFilterer) ParseBonded(log types.Log) (*ValidatorPoolBonded, error) {
	event := new(ValidatorPoolBonded)
	if err := _ValidatorPool.contract.UnpackLog(event, "Bonded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorPoolInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ValidatorPool contract.
type ValidatorPoolInitializedIterator struct {
	Event *ValidatorPoolInitialized // Event containing the contract specifics and raw log

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
func (it *ValidatorPoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorPoolInitialized)
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
		it.Event = new(ValidatorPoolInitialized)
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
func (it *ValidatorPoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorPoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorPoolInitialized represents a Initialized event raised by the ValidatorPool contract.
type ValidatorPoolInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ValidatorPool *ValidatorPoolFilterer) FilterInitialized(opts *bind.FilterOpts) (*ValidatorPoolInitializedIterator, error) {

	logs, sub, err := _ValidatorPool.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ValidatorPoolInitializedIterator{contract: _ValidatorPool.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ValidatorPool *ValidatorPoolFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ValidatorPoolInitialized) (event.Subscription, error) {

	logs, sub, err := _ValidatorPool.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorPoolInitialized)
				if err := _ValidatorPool.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ValidatorPool *ValidatorPoolFilterer) ParseInitialized(log types.Log) (*ValidatorPoolInitialized, error) {
	event := new(ValidatorPoolInitialized)
	if err := _ValidatorPool.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorPoolUnbondedIterator is returned from FilterUnbonded and is used to iterate over the raw logs and unpacked data for Unbonded events raised by the ValidatorPool contract.
type ValidatorPoolUnbondedIterator struct {
	Event *ValidatorPoolUnbonded // Event containing the contract specifics and raw log

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
func (it *ValidatorPoolUnbondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorPoolUnbonded)
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
		it.Event = new(ValidatorPoolUnbonded)
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
func (it *ValidatorPoolUnbondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorPoolUnbondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorPoolUnbonded represents a Unbonded event raised by the ValidatorPool contract.
type ValidatorPoolUnbonded struct {
	OutputIndex *big.Int
	Recipient   common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnbonded is a free log retrieval operation binding the contract event 0x7047a0fb8bfae78c0ebbd4117437945bb85240453235ac4fd2e55712eb5bf0c3.
//
// Solidity: event Unbonded(uint256 indexed outputIndex, address indexed recipient, uint128 amount)
func (_ValidatorPool *ValidatorPoolFilterer) FilterUnbonded(opts *bind.FilterOpts, outputIndex []*big.Int, recipient []common.Address) (*ValidatorPoolUnbondedIterator, error) {

	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ValidatorPool.contract.FilterLogs(opts, "Unbonded", outputIndexRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorPoolUnbondedIterator{contract: _ValidatorPool.contract, event: "Unbonded", logs: logs, sub: sub}, nil
}

// WatchUnbonded is a free log subscription operation binding the contract event 0x7047a0fb8bfae78c0ebbd4117437945bb85240453235ac4fd2e55712eb5bf0c3.
//
// Solidity: event Unbonded(uint256 indexed outputIndex, address indexed recipient, uint128 amount)
func (_ValidatorPool *ValidatorPoolFilterer) WatchUnbonded(opts *bind.WatchOpts, sink chan<- *ValidatorPoolUnbonded, outputIndex []*big.Int, recipient []common.Address) (event.Subscription, error) {

	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ValidatorPool.contract.WatchLogs(opts, "Unbonded", outputIndexRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorPoolUnbonded)
				if err := _ValidatorPool.contract.UnpackLog(event, "Unbonded", log); err != nil {
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

// ParseUnbonded is a log parse operation binding the contract event 0x7047a0fb8bfae78c0ebbd4117437945bb85240453235ac4fd2e55712eb5bf0c3.
//
// Solidity: event Unbonded(uint256 indexed outputIndex, address indexed recipient, uint128 amount)
func (_ValidatorPool *ValidatorPoolFilterer) ParseUnbonded(log types.Log) (*ValidatorPoolUnbonded, error) {
	event := new(ValidatorPoolUnbonded)
	if err := _ValidatorPool.contract.UnpackLog(event, "Unbonded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
