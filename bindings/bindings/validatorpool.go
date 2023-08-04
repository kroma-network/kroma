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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractL2OutputOracle\",\"name\":\"_l2OutputOracle\",\"type\":\"address\"},{\"internalType\":\"contractKromaPortal\",\"name\":\"_portal\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedValidator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minBondAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxUnbond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_roundDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"outputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"BondIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"outputIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"expiresAt\",\"type\":\"uint128\"}],\"name\":\"Bonded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"outputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"PendingBondAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"outputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"PendingBondRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"outputIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"Unbonded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"L2_ORACLE\",\"outputs\":[{\"internalType\":\"contractL2OutputOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_UNBOND\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BOND_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PORTAL\",\"outputs\":[{\"internalType\":\"contractKromaPortal\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUND_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRUSTED_VALIDATOR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VAULT_REWARD_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_outputIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_challenger\",\"type\":\"address\"}],\"name\":\"addPendingBond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_outputIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_expiresAt\",\"type\":\"uint128\"}],\"name\":\"createBond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_outputIndex\",\"type\":\"uint256\"}],\"name\":\"getBond\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"expiresAt\",\"type\":\"uint128\"}],\"internalType\":\"structTypes.Bond\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_outputIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_challenger\",\"type\":\"address\"}],\"name\":\"getPendingBond\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_outputIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_challenger\",\"type\":\"address\"}],\"name\":\"increaseBond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextValidator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_outputIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_challenger\",\"type\":\"address\"}],\"name\":\"refundPendingBond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unbond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101a06040523480156200001257600080fd5b5060405162002bfd38038062002bfd833981016040819052620000359162000251565b60006080819052600160a05260c0526001600160a01b0380871660e052858116610100528416610120526101408390526101608290526101808190526200007b62000087565b505050505050620002c1565b600054610100900460ff1615808015620000a85750600054600160ff909116105b80620000d85750620000c530620001b960201b6200194f1760201c565b158015620000d8575060005460ff166001145b620001415760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff19166001179055801562000165576000805461ff0019166101001790555b6200016f620001c8565b8015620001b6576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b6001600160a01b03163b151590565b600054610100900460ff16620002355760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201526a6e697469616c697a696e6760a81b606482015260840162000138565b60018055565b6001600160a01b0381168114620001b657600080fd5b60008060008060008060c087890312156200026b57600080fd5b865162000278816200023b565b60208801519096506200028b816200023b565b60408801519095506200029e816200023b565b80945050606087015192506080870151915060a087015190509295509295509295565b60805160a05160c05160e05161010051610120516101405161016051610180516128626200039b6000396000818161031e0152610a670152600081816103eb0152611e9e01526000818161043f01528181611197015281816119a00152611b310152600081816102930152610acf015260008181610208015261215501526000818161018b0152818161053b01528181610917015281816109bf01528181610b960152818161110201528181611321015281816115770152611f6801526000610b4a01526000610b2101526000610af801526128626000f3fe6080604052600436106101745760003560e01c806370a08231116100cb5780639fbc4a5f1161007f578063d8fe764211610059578063d8fe764214610499578063dd215c5d146104e9578063facd743b1461050957600080fd5b80639fbc4a5f1461042d578063ab91f19014610461578063d0e30db01461049157600080fd5b80638f09ade4116100b05780638f09ade414610398578063946765fd146103d957806396946f751461040d57600080fd5b806370a08231146103405780638129fc1c1461038357600080fd5b80633a5490461161012d5780635a544742116101075780635a544742146102d75780635df6a6bc146102f75780636641ea081461030c57600080fd5b80633a5490461461026c5780633ee4d4a31461028157806354fd4d50146102b557600080fd5b80630ff754ea1161015e5780630ff754ea146101f65780632b62052c1461022a5780632e1a7d4d1461024c57600080fd5b80621c2ff6146101795780630f43a677146101d7575b600080fd5b34801561018557600080fd5b506101ad7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156101e357600080fd5b506036545b6040519081526020016101ce565b34801561020257600080fd5b506101ad7f000000000000000000000000000000000000000000000000000000000000000081565b34801561023657600080fd5b5061024a6102453660046123f1565b610539565b005b34801561025857600080fd5b5061024a610267366004612421565b6107f7565b34801561027857600080fd5b506101ad6108f2565b34801561028d57600080fd5b506101ad7f000000000000000000000000000000000000000000000000000000000000000081565b3480156102c157600080fd5b506102ca610af1565b6040516101ce91906124b4565b3480156102e357600080fd5b5061024a6102f23660046123f1565b610b94565b34801561030357600080fd5b5061024a610e35565b34801561031857600080fd5b506101e87f000000000000000000000000000000000000000000000000000000000000000081565b34801561034c57600080fd5b506101e861035b3660046124c7565b73ffffffffffffffffffffffffffffffffffffffff1660009081526033602052604090205490565b34801561038f57600080fd5b5061024a610eb7565b3480156103a457600080fd5b506103b86103b33660046123f1565b61102e565b6040516fffffffffffffffffffffffffffffffff90911681526020016101ce565b3480156103e557600080fd5b506101e87f000000000000000000000000000000000000000000000000000000000000000081565b34801561041957600080fd5b5061024a610428366004612502565b6110ea565b34801561043957600080fd5b506101e87f000000000000000000000000000000000000000000000000000000000000000081565b34801561046d57600080fd5b50610478620186a081565b60405167ffffffffffffffff90911681526020016101ce565b61024a61144a565b3480156104a557600080fd5b506104b96104b4366004612421565b611456565b6040805182516fffffffffffffffffffffffffffffffff90811682526020938401511692810192909252016101ce565b3480156104f557600080fd5b5061024a6105043660046123f1565b611575565b34801561051557600080fd5b506105296105243660046124c7565b6118b1565b60405190151581526020016101ce565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16639e45e8f46040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c89190612544565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461066d5760405162461bcd60e51b815260206004820152602660248201527f56616c696461746f72506f6f6c3a2073656e646572206973206e6f7420436f6c60448201527f6f737365756d000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b600082815260396020908152604080832073ffffffffffffffffffffffffffffffffffffffff851684529091529020546fffffffffffffffffffffffffffffffff16806107225760405162461bcd60e51b815260206004820152602e60248201527f56616c696461746f72506f6f6c3a207468652070656e64696e6720626f6e642060448201527f646f6573206e6f742065786973740000000000000000000000000000000000006064820152608401610664565b600083815260396020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffff00000000000000000000000000000000169055610793826fffffffffffffffffffffffffffffffff831661196b565b6040516fffffffffffffffffffffffffffffffff8216815273ffffffffffffffffffffffffffffffffffffffff83169084907f7dd3b3488bb77a5de73fe533a6fe2e06d495b6498cd57eef623053754b7a3f119060200160405180910390a3505050565b6002600154036108495760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610664565b60026001556108583382611a88565b6000610875335a8460405180602001604052806000815250611d15565b9050806108ea5760405162461bcd60e51b815260206004820152602260248201527f56616c696461746f72506f6f6c3a20455448207472616e73666572206661696c60448201527f65640000000000000000000000000000000000000000000000000000000000006064820152608401610664565b505060018055565b60385460009073ffffffffffffffffffffffffffffffffffffffff1615610acc5760007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663dcec33486040518163ffffffff1660e01b8152600401602060405180830381865afa158015610980573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109a49190612561565b9050600073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663d1de856c6109ef8460016125a9565b6040518263ffffffff1660e01b8152600401610a0d91815260200190565b602060405180830381865afa158015610a2a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a4e9190612561565b9050804210610aac576000610a6382426125c1565b90507f0000000000000000000000000000000000000000000000000000000000000000811115610aaa5773ffffffffffffffffffffffffffffffffffffffff935050505090565b505b505060385473ffffffffffffffffffffffffffffffffffffffff16919050565b507f000000000000000000000000000000000000000000000000000000000000000090565b6060610b1c7f0000000000000000000000000000000000000000000000000000000000000000611d2f565b610b457f0000000000000000000000000000000000000000000000000000000000000000611d2f565b610b6e7f0000000000000000000000000000000000000000000000000000000000000000611d2f565b604051602001610b80939291906125d8565b604051602081830303815290604052905090565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16639e45e8f46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610bff573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c239190612544565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610cc35760405162461bcd60e51b815260206004820152602660248201527f56616c696461746f72506f6f6c3a2073656e646572206973206e6f7420436f6c60448201527f6f737365756d00000000000000000000000000000000000000000000000000006064820152608401610664565b60008281526034602052604090208054427001000000000000000000000000000000009091046fffffffffffffffffffffffffffffffff161015610d6f5760405162461bcd60e51b815260206004820152602e60248201527f56616c696461746f72506f6f6c3a20746865206f757470757420697320616c7260448201527f656164792066696e616c697a65640000000000000000000000000000000000006064820152608401610664565b80546fffffffffffffffffffffffffffffffff16610d8d8382611a88565b600084815260396020908152604080832073ffffffffffffffffffffffffffffffffffffffff87168085529083529281902080547fffffffffffffffffffffffffffffffff00000000000000000000000000000000166fffffffffffffffffffffffffffffffff8616908117909155905190815286917f2904258f32adf74dd8f23ad6f17ff50209896039c8ee3d4728ff55bd05c4cf2a91015b60405180910390a350505050565b6000610e3f611e6c565b905080610eb45760405162461bcd60e51b815260206004820152602960248201527f56616c696461746f72506f6f6c3a206e6f20626f6e6420746861742063616e2060448201527f626520756e626f6e6400000000000000000000000000000000000000000000006064820152608401610664565b50565b600054610100900460ff1615808015610ed75750600054600160ff909116105b80610ef15750303b158015610ef1575060005460ff166001145b610f635760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610664565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610fc157600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610fc96120b1565b8015610eb457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150565b600082815260396020908152604080832073ffffffffffffffffffffffffffffffffffffffff851684529091528120546fffffffffffffffffffffffffffffffff16806110e35760405162461bcd60e51b815260206004820152602e60248201527f56616c696461746f72506f6f6c3a207468652070656e64696e6720626f6e642060448201527f646f6573206e6f742065786973740000000000000000000000000000000000006064820152608401610664565b9392505050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146111955760405162461bcd60e51b815260206004820152602b60248201527f56616c696461746f72506f6f6c3a2073656e646572206973206e6f74204c324f60448201527f75747075744f7261636c650000000000000000000000000000000000000000006064820152608401610664565b7f0000000000000000000000000000000000000000000000000000000000000000826fffffffffffffffffffffffffffffffff16101561123d5760405162461bcd60e51b815260206004820152602b60248201527f56616c696461746f72506f6f6c3a2074686520626f6e6420616d6f756e74206960448201527f7320746f6f20736d616c6c0000000000000000000000000000000000000000006064820152608401610664565b6000838152603460205260409020805470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff16156112e65760405162461bcd60e51b815260206004820152603c60248201527f56616c696461746f72506f6f6c3a20626f6e64206f662074686520676976656e60448201527f206f757470757420696e64657820616c726561647920657869737473000000006064820152608401610664565b6112ee611e6c565b506040517fb0ea09a8000000000000000000000000000000000000000000000000000000008152600481018590526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063b0ea09a890602401602060405180830381865afa15801561137d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113a19190612544565b90506113bf81856fffffffffffffffffffffffffffffffff16611a88565b6fffffffffffffffffffffffffffffffff84811670010000000000000000000000000000000091851691820281178455604080519182526020820192909252869173ffffffffffffffffffffffffffffffffffffffff8416917f5ca130257b8f76f72ad2965efcbe166f3918d820e4a49956e70081ea311f97c4910160405180910390a35050505050565b611454333461196b565b565b6040805180820190915260008082526020820152600082815260346020526040902080546fffffffffffffffffffffffffffffffff16158015906114c05750805470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff1615155b6115325760405162461bcd60e51b815260206004820152602660248201527f56616c696461746f72506f6f6c3a2074686520626f6e6420646f6573206e6f7460448201527f20657869737400000000000000000000000000000000000000000000000000006064820152608401610664565b6040805180820190915290546fffffffffffffffffffffffffffffffff808216835270010000000000000000000000000000000090910416602082015292915050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16639e45e8f46040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116049190612544565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146116a45760405162461bcd60e51b815260206004820152602660248201527f56616c696461746f72506f6f6c3a2073656e646572206973206e6f7420436f6c60448201527f6f737365756d00000000000000000000000000000000000000000000000000006064820152608401610664565b60008281526034602052604090208054427001000000000000000000000000000000009091046fffffffffffffffffffffffffffffffff1610156117505760405162461bcd60e51b815260206004820152602e60248201527f56616c696461746f72506f6f6c3a20746865206f757470757420697320616c7260448201527f656164792066696e616c697a65640000000000000000000000000000000000006064820152608401610664565b600083815260396020908152604080832073ffffffffffffffffffffffffffffffffffffffff861684529091529020546fffffffffffffffffffffffffffffffff16806118055760405162461bcd60e51b815260206004820152602e60248201527f56616c696461746f72506f6f6c3a207468652070656e64696e6720626f6e642060448201527f646f6573206e6f742065786973740000000000000000000000000000000000006064820152608401610664565b600084815260396020908152604080832073ffffffffffffffffffffffffffffffffffffffff87168085529083529281902080547fffffffffffffffffffffffffffffffff0000000000000000000000000000000090811690915585549081166fffffffffffffffffffffffffffffffff918216860182161786559051908416815286917f383f9b8b5a1fc2ec555726eb895621a312042e18b764135fa12ef1a520ad30db9101610e27565b60365460009081036118c557506000919050565b73ffffffffffffffffffffffffffffffffffffffff82166118e857506000919050565b73ffffffffffffffffffffffffffffffffffffffff821660008181526037602052604090205460368054919291839081106119255761192561264e565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16149392505050565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b73ffffffffffffffffffffffffffffffffffffffff821660009081526033602052604081205461199c9083906125a9565b90507f000000000000000000000000000000000000000000000000000000000000000081101580156119d457506119d2836118b1565b155b15611a5b576036805473ffffffffffffffffffffffffffffffffffffffff85166000818152603760205260408120839055600183018455929092527f4a11f94e20a93c79f6ec743a1954ec4fc2c08429ae2122118bf234b2185c81b80180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690911790555b73ffffffffffffffffffffffffffffffffffffffff90921660009081526033602052604090209190915550565b73ffffffffffffffffffffffffffffffffffffffff821660009081526033602052604090205481811015611b235760405162461bcd60e51b8152602060048201526024808201527f56616c696461746f72506f6f6c3a20696e73756666696369656e742062616c6160448201527f6e636573000000000000000000000000000000000000000000000000000000006064820152608401610664565b611b2d82826125c1565b90507f000000000000000000000000000000000000000000000000000000000000000081108015611b625750611b62836118b1565b15611a5b57603654600090611b79906001906125c1565b90508015611c585773ffffffffffffffffffffffffffffffffffffffff84166000908152603760205260408120546036805491929184908110611bbe57611bbe61264e565b6000918252602090912001546036805473ffffffffffffffffffffffffffffffffffffffff9092169250829184908110611bfa57611bfa61264e565b600091825260208083209190910180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff948516179055929091168152603790915260409020555b73ffffffffffffffffffffffffffffffffffffffff84166000908152603760205260408120556036805480611c8f57611c8f61267d565b60008281526020902081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690550190555073ffffffffffffffffffffffffffffffffffffffff90921660009081526033602052604090209190915550565b600080600080845160208601878a8af19695505050505050565b606081600003611d7257505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115611d9c5780611d86816126ac565b9150611d959050600a83612713565b9150611d76565b60008167ffffffffffffffff811115611db757611db7612727565b6040519080825280601f01601f191660200182016040528015611de1576020820181803683370190505b5090505b8415611e6457611df66001836125c1565b9150611e03600a86612756565b611e0e9060306125a9565b60f81b818381518110611e2357611e2361264e565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350611e5d600a86612713565b9450611de5565b949350505050565b60355460408051608081018252600080825260208201819052918101829052606081018290529091908290819060005b7f000000000000000000000000000000000000000000000000000000000000000081101561208157600085815260346020526040902080546fffffffffffffffffffffffffffffffff80821696509194507001000000000000000000000000000000009004164210801590611f2357506000846fffffffffffffffffffffffffffffffff16115b156120815760008581526034602052604080822091909155517fa25ae557000000000000000000000000000000000000000000000000000000008152600481018690527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063a25ae55790602401608060405180830381865afa158015611fc4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fe8919061276a565b915061200a8260000151856fffffffffffffffffffffffffffffffff1661196b565b81516040516fffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff9091169086907f7047a0fb8bfae78c0ebbd4117437945bb85240453235ac4fd2e55712eb5bf0c39060200160405180910390a361207582612134565b60019485019401611e9c565b80156120a55761209482602001516122bf565b505050603591909155506001919050565b60009550505050505090565b600054610100900460ff1661212e5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610664565b60018055565b8051606082015160405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169263c30af3889273420000000000000000000000000000000000000892620186a0927f21670f2200000000000000000000000000000000000000000000000000000000926121f89260240173ffffffffffffffffffffffffffffffffffffffff9290921682526fffffffffffffffffffffffffffffffff16602082015260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e086901b909216825261228a93929160040161280d565b600060405180830381600087803b1580156122a457600080fd5b505af11580156122b8573d6000803e3d6000fd5b5050505050565b60365480156123a35760008183434160405160200161231693929190928352602083019190915260601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016604082015260540190565b6040516020818303038152906040528051906020012060001c6123399190612756565b90506036818154811061234e5761234e61264e565b600091825260209091200154603880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909216919091179055505050565b603880547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555050565b73ffffffffffffffffffffffffffffffffffffffff81168114610eb457600080fd5b6000806040838503121561240457600080fd5b823591506020830135612416816123cf565b809150509250929050565b60006020828403121561243357600080fd5b5035919050565b60005b8381101561245557818101518382015260200161243d565b83811115612464576000848401525b50505050565b6000815180845261248281602086016020860161243a565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006110e3602083018461246a565b6000602082840312156124d957600080fd5b81356110e3816123cf565b6fffffffffffffffffffffffffffffffff81168114610eb457600080fd5b60008060006060848603121561251757600080fd5b833592506020840135612529816124e4565b91506040840135612539816124e4565b809150509250925092565b60006020828403121561255657600080fd5b81516110e3816123cf565b60006020828403121561257357600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156125bc576125bc61257a565b500190565b6000828210156125d3576125d361257a565b500390565b600084516125ea81846020890161243a565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551612626816001850160208a0161243a565b6001920191820152835161264181600284016020880161243a565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036126dd576126dd61257a565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082612722576127226126e4565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082612765576127656126e4565b500690565b60006080828403121561277c57600080fd5b6040516080810181811067ffffffffffffffff821117156127c6577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405282516127d4816123cf565b81526020838101519082015260408301516127ee816124e4565b60408201526060830151612801816124e4565b60608201529392505050565b73ffffffffffffffffffffffffffffffffffffffff8416815267ffffffffffffffff8316602082015260606040820152600061284c606083018461246a565b9594505050505056fea164736f6c634300080f000a",
}

// ValidatorPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorPoolMetaData.ABI instead.
var ValidatorPoolABI = ValidatorPoolMetaData.ABI

// ValidatorPoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorPoolMetaData.Bin instead.
var ValidatorPoolBin = ValidatorPoolMetaData.Bin

// DeployValidatorPool deploys a new Ethereum contract, binding an instance of ValidatorPool to it.
func DeployValidatorPool(auth *bind.TransactOpts, backend bind.ContractBackend, _l2OutputOracle common.Address, _portal common.Address, _trustedValidator common.Address, _minBondAmount *big.Int, _maxUnbond *big.Int, _roundDuration *big.Int) (common.Address, *types.Transaction, *ValidatorPool, error) {
	parsed, err := ValidatorPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorPoolBin), backend, _l2OutputOracle, _portal, _trustedValidator, _minBondAmount, _maxUnbond, _roundDuration)
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

// MAXUNBOND is a free data retrieval call binding the contract method 0x946765fd.
//
// Solidity: function MAX_UNBOND() view returns(uint256)
func (_ValidatorPool *ValidatorPoolCaller) MAXUNBOND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "MAX_UNBOND")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXUNBOND is a free data retrieval call binding the contract method 0x946765fd.
//
// Solidity: function MAX_UNBOND() view returns(uint256)
func (_ValidatorPool *ValidatorPoolSession) MAXUNBOND() (*big.Int, error) {
	return _ValidatorPool.Contract.MAXUNBOND(&_ValidatorPool.CallOpts)
}

// MAXUNBOND is a free data retrieval call binding the contract method 0x946765fd.
//
// Solidity: function MAX_UNBOND() view returns(uint256)
func (_ValidatorPool *ValidatorPoolCallerSession) MAXUNBOND() (*big.Int, error) {
	return _ValidatorPool.Contract.MAXUNBOND(&_ValidatorPool.CallOpts)
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

// PORTAL is a free data retrieval call binding the contract method 0x0ff754ea.
//
// Solidity: function PORTAL() view returns(address)
func (_ValidatorPool *ValidatorPoolCaller) PORTAL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "PORTAL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PORTAL is a free data retrieval call binding the contract method 0x0ff754ea.
//
// Solidity: function PORTAL() view returns(address)
func (_ValidatorPool *ValidatorPoolSession) PORTAL() (common.Address, error) {
	return _ValidatorPool.Contract.PORTAL(&_ValidatorPool.CallOpts)
}

// PORTAL is a free data retrieval call binding the contract method 0x0ff754ea.
//
// Solidity: function PORTAL() view returns(address)
func (_ValidatorPool *ValidatorPoolCallerSession) PORTAL() (common.Address, error) {
	return _ValidatorPool.Contract.PORTAL(&_ValidatorPool.CallOpts)
}

// ROUNDDURATION is a free data retrieval call binding the contract method 0x6641ea08.
//
// Solidity: function ROUND_DURATION() view returns(uint256)
func (_ValidatorPool *ValidatorPoolCaller) ROUNDDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "ROUND_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ROUNDDURATION is a free data retrieval call binding the contract method 0x6641ea08.
//
// Solidity: function ROUND_DURATION() view returns(uint256)
func (_ValidatorPool *ValidatorPoolSession) ROUNDDURATION() (*big.Int, error) {
	return _ValidatorPool.Contract.ROUNDDURATION(&_ValidatorPool.CallOpts)
}

// ROUNDDURATION is a free data retrieval call binding the contract method 0x6641ea08.
//
// Solidity: function ROUND_DURATION() view returns(uint256)
func (_ValidatorPool *ValidatorPoolCallerSession) ROUNDDURATION() (*big.Int, error) {
	return _ValidatorPool.Contract.ROUNDDURATION(&_ValidatorPool.CallOpts)
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

// VAULTREWARDGASLIMIT is a free data retrieval call binding the contract method 0xab91f190.
//
// Solidity: function VAULT_REWARD_GAS_LIMIT() view returns(uint64)
func (_ValidatorPool *ValidatorPoolCaller) VAULTREWARDGASLIMIT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "VAULT_REWARD_GAS_LIMIT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VAULTREWARDGASLIMIT is a free data retrieval call binding the contract method 0xab91f190.
//
// Solidity: function VAULT_REWARD_GAS_LIMIT() view returns(uint64)
func (_ValidatorPool *ValidatorPoolSession) VAULTREWARDGASLIMIT() (uint64, error) {
	return _ValidatorPool.Contract.VAULTREWARDGASLIMIT(&_ValidatorPool.CallOpts)
}

// VAULTREWARDGASLIMIT is a free data retrieval call binding the contract method 0xab91f190.
//
// Solidity: function VAULT_REWARD_GAS_LIMIT() view returns(uint64)
func (_ValidatorPool *ValidatorPoolCallerSession) VAULTREWARDGASLIMIT() (uint64, error) {
	return _ValidatorPool.Contract.VAULTREWARDGASLIMIT(&_ValidatorPool.CallOpts)
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

// GetPendingBond is a free data retrieval call binding the contract method 0x8f09ade4.
//
// Solidity: function getPendingBond(uint256 _outputIndex, address _challenger) view returns(uint128)
func (_ValidatorPool *ValidatorPoolCaller) GetPendingBond(opts *bind.CallOpts, _outputIndex *big.Int, _challenger common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ValidatorPool.contract.Call(opts, &out, "getPendingBond", _outputIndex, _challenger)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingBond is a free data retrieval call binding the contract method 0x8f09ade4.
//
// Solidity: function getPendingBond(uint256 _outputIndex, address _challenger) view returns(uint128)
func (_ValidatorPool *ValidatorPoolSession) GetPendingBond(_outputIndex *big.Int, _challenger common.Address) (*big.Int, error) {
	return _ValidatorPool.Contract.GetPendingBond(&_ValidatorPool.CallOpts, _outputIndex, _challenger)
}

// GetPendingBond is a free data retrieval call binding the contract method 0x8f09ade4.
//
// Solidity: function getPendingBond(uint256 _outputIndex, address _challenger) view returns(uint128)
func (_ValidatorPool *ValidatorPoolCallerSession) GetPendingBond(_outputIndex *big.Int, _challenger common.Address) (*big.Int, error) {
	return _ValidatorPool.Contract.GetPendingBond(&_ValidatorPool.CallOpts, _outputIndex, _challenger)
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

// AddPendingBond is a paid mutator transaction binding the contract method 0x5a544742.
//
// Solidity: function addPendingBond(uint256 _outputIndex, address _challenger) returns()
func (_ValidatorPool *ValidatorPoolTransactor) AddPendingBond(opts *bind.TransactOpts, _outputIndex *big.Int, _challenger common.Address) (*types.Transaction, error) {
	return _ValidatorPool.contract.Transact(opts, "addPendingBond", _outputIndex, _challenger)
}

// AddPendingBond is a paid mutator transaction binding the contract method 0x5a544742.
//
// Solidity: function addPendingBond(uint256 _outputIndex, address _challenger) returns()
func (_ValidatorPool *ValidatorPoolSession) AddPendingBond(_outputIndex *big.Int, _challenger common.Address) (*types.Transaction, error) {
	return _ValidatorPool.Contract.AddPendingBond(&_ValidatorPool.TransactOpts, _outputIndex, _challenger)
}

// AddPendingBond is a paid mutator transaction binding the contract method 0x5a544742.
//
// Solidity: function addPendingBond(uint256 _outputIndex, address _challenger) returns()
func (_ValidatorPool *ValidatorPoolTransactorSession) AddPendingBond(_outputIndex *big.Int, _challenger common.Address) (*types.Transaction, error) {
	return _ValidatorPool.Contract.AddPendingBond(&_ValidatorPool.TransactOpts, _outputIndex, _challenger)
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

// IncreaseBond is a paid mutator transaction binding the contract method 0xdd215c5d.
//
// Solidity: function increaseBond(uint256 _outputIndex, address _challenger) returns()
func (_ValidatorPool *ValidatorPoolTransactor) IncreaseBond(opts *bind.TransactOpts, _outputIndex *big.Int, _challenger common.Address) (*types.Transaction, error) {
	return _ValidatorPool.contract.Transact(opts, "increaseBond", _outputIndex, _challenger)
}

// IncreaseBond is a paid mutator transaction binding the contract method 0xdd215c5d.
//
// Solidity: function increaseBond(uint256 _outputIndex, address _challenger) returns()
func (_ValidatorPool *ValidatorPoolSession) IncreaseBond(_outputIndex *big.Int, _challenger common.Address) (*types.Transaction, error) {
	return _ValidatorPool.Contract.IncreaseBond(&_ValidatorPool.TransactOpts, _outputIndex, _challenger)
}

// IncreaseBond is a paid mutator transaction binding the contract method 0xdd215c5d.
//
// Solidity: function increaseBond(uint256 _outputIndex, address _challenger) returns()
func (_ValidatorPool *ValidatorPoolTransactorSession) IncreaseBond(_outputIndex *big.Int, _challenger common.Address) (*types.Transaction, error) {
	return _ValidatorPool.Contract.IncreaseBond(&_ValidatorPool.TransactOpts, _outputIndex, _challenger)
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

// RefundPendingBond is a paid mutator transaction binding the contract method 0x2b62052c.
//
// Solidity: function refundPendingBond(uint256 _outputIndex, address _challenger) returns()
func (_ValidatorPool *ValidatorPoolTransactor) RefundPendingBond(opts *bind.TransactOpts, _outputIndex *big.Int, _challenger common.Address) (*types.Transaction, error) {
	return _ValidatorPool.contract.Transact(opts, "refundPendingBond", _outputIndex, _challenger)
}

// RefundPendingBond is a paid mutator transaction binding the contract method 0x2b62052c.
//
// Solidity: function refundPendingBond(uint256 _outputIndex, address _challenger) returns()
func (_ValidatorPool *ValidatorPoolSession) RefundPendingBond(_outputIndex *big.Int, _challenger common.Address) (*types.Transaction, error) {
	return _ValidatorPool.Contract.RefundPendingBond(&_ValidatorPool.TransactOpts, _outputIndex, _challenger)
}

// RefundPendingBond is a paid mutator transaction binding the contract method 0x2b62052c.
//
// Solidity: function refundPendingBond(uint256 _outputIndex, address _challenger) returns()
func (_ValidatorPool *ValidatorPoolTransactorSession) RefundPendingBond(_outputIndex *big.Int, _challenger common.Address) (*types.Transaction, error) {
	return _ValidatorPool.Contract.RefundPendingBond(&_ValidatorPool.TransactOpts, _outputIndex, _challenger)
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
	OutputIndex *big.Int
	Challenger  common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBondIncreased is a free log retrieval operation binding the contract event 0x383f9b8b5a1fc2ec555726eb895621a312042e18b764135fa12ef1a520ad30db.
//
// Solidity: event BondIncreased(uint256 indexed outputIndex, address indexed challenger, uint128 amount)
func (_ValidatorPool *ValidatorPoolFilterer) FilterBondIncreased(opts *bind.FilterOpts, outputIndex []*big.Int, challenger []common.Address) (*ValidatorPoolBondIncreasedIterator, error) {

	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ValidatorPool.contract.FilterLogs(opts, "BondIncreased", outputIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorPoolBondIncreasedIterator{contract: _ValidatorPool.contract, event: "BondIncreased", logs: logs, sub: sub}, nil
}

// WatchBondIncreased is a free log subscription operation binding the contract event 0x383f9b8b5a1fc2ec555726eb895621a312042e18b764135fa12ef1a520ad30db.
//
// Solidity: event BondIncreased(uint256 indexed outputIndex, address indexed challenger, uint128 amount)
func (_ValidatorPool *ValidatorPoolFilterer) WatchBondIncreased(opts *bind.WatchOpts, sink chan<- *ValidatorPoolBondIncreased, outputIndex []*big.Int, challenger []common.Address) (event.Subscription, error) {

	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ValidatorPool.contract.WatchLogs(opts, "BondIncreased", outputIndexRule, challengerRule)
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

// ParseBondIncreased is a log parse operation binding the contract event 0x383f9b8b5a1fc2ec555726eb895621a312042e18b764135fa12ef1a520ad30db.
//
// Solidity: event BondIncreased(uint256 indexed outputIndex, address indexed challenger, uint128 amount)
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

// ValidatorPoolPendingBondAddedIterator is returned from FilterPendingBondAdded and is used to iterate over the raw logs and unpacked data for PendingBondAdded events raised by the ValidatorPool contract.
type ValidatorPoolPendingBondAddedIterator struct {
	Event *ValidatorPoolPendingBondAdded // Event containing the contract specifics and raw log

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
func (it *ValidatorPoolPendingBondAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorPoolPendingBondAdded)
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
		it.Event = new(ValidatorPoolPendingBondAdded)
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
func (it *ValidatorPoolPendingBondAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorPoolPendingBondAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorPoolPendingBondAdded represents a PendingBondAdded event raised by the ValidatorPool contract.
type ValidatorPoolPendingBondAdded struct {
	OutputIndex *big.Int
	Challenger  common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPendingBondAdded is a free log retrieval operation binding the contract event 0x2904258f32adf74dd8f23ad6f17ff50209896039c8ee3d4728ff55bd05c4cf2a.
//
// Solidity: event PendingBondAdded(uint256 indexed outputIndex, address indexed challenger, uint128 amount)
func (_ValidatorPool *ValidatorPoolFilterer) FilterPendingBondAdded(opts *bind.FilterOpts, outputIndex []*big.Int, challenger []common.Address) (*ValidatorPoolPendingBondAddedIterator, error) {

	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ValidatorPool.contract.FilterLogs(opts, "PendingBondAdded", outputIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorPoolPendingBondAddedIterator{contract: _ValidatorPool.contract, event: "PendingBondAdded", logs: logs, sub: sub}, nil
}

// WatchPendingBondAdded is a free log subscription operation binding the contract event 0x2904258f32adf74dd8f23ad6f17ff50209896039c8ee3d4728ff55bd05c4cf2a.
//
// Solidity: event PendingBondAdded(uint256 indexed outputIndex, address indexed challenger, uint128 amount)
func (_ValidatorPool *ValidatorPoolFilterer) WatchPendingBondAdded(opts *bind.WatchOpts, sink chan<- *ValidatorPoolPendingBondAdded, outputIndex []*big.Int, challenger []common.Address) (event.Subscription, error) {

	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ValidatorPool.contract.WatchLogs(opts, "PendingBondAdded", outputIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorPoolPendingBondAdded)
				if err := _ValidatorPool.contract.UnpackLog(event, "PendingBondAdded", log); err != nil {
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

// ParsePendingBondAdded is a log parse operation binding the contract event 0x2904258f32adf74dd8f23ad6f17ff50209896039c8ee3d4728ff55bd05c4cf2a.
//
// Solidity: event PendingBondAdded(uint256 indexed outputIndex, address indexed challenger, uint128 amount)
func (_ValidatorPool *ValidatorPoolFilterer) ParsePendingBondAdded(log types.Log) (*ValidatorPoolPendingBondAdded, error) {
	event := new(ValidatorPoolPendingBondAdded)
	if err := _ValidatorPool.contract.UnpackLog(event, "PendingBondAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorPoolPendingBondRefundedIterator is returned from FilterPendingBondRefunded and is used to iterate over the raw logs and unpacked data for PendingBondRefunded events raised by the ValidatorPool contract.
type ValidatorPoolPendingBondRefundedIterator struct {
	Event *ValidatorPoolPendingBondRefunded // Event containing the contract specifics and raw log

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
func (it *ValidatorPoolPendingBondRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorPoolPendingBondRefunded)
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
		it.Event = new(ValidatorPoolPendingBondRefunded)
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
func (it *ValidatorPoolPendingBondRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorPoolPendingBondRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorPoolPendingBondRefunded represents a PendingBondRefunded event raised by the ValidatorPool contract.
type ValidatorPoolPendingBondRefunded struct {
	OutputIndex *big.Int
	Challenger  common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPendingBondRefunded is a free log retrieval operation binding the contract event 0x7dd3b3488bb77a5de73fe533a6fe2e06d495b6498cd57eef623053754b7a3f11.
//
// Solidity: event PendingBondRefunded(uint256 indexed outputIndex, address indexed challenger, uint128 amount)
func (_ValidatorPool *ValidatorPoolFilterer) FilterPendingBondRefunded(opts *bind.FilterOpts, outputIndex []*big.Int, challenger []common.Address) (*ValidatorPoolPendingBondRefundedIterator, error) {

	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ValidatorPool.contract.FilterLogs(opts, "PendingBondRefunded", outputIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorPoolPendingBondRefundedIterator{contract: _ValidatorPool.contract, event: "PendingBondRefunded", logs: logs, sub: sub}, nil
}

// WatchPendingBondRefunded is a free log subscription operation binding the contract event 0x7dd3b3488bb77a5de73fe533a6fe2e06d495b6498cd57eef623053754b7a3f11.
//
// Solidity: event PendingBondRefunded(uint256 indexed outputIndex, address indexed challenger, uint128 amount)
func (_ValidatorPool *ValidatorPoolFilterer) WatchPendingBondRefunded(opts *bind.WatchOpts, sink chan<- *ValidatorPoolPendingBondRefunded, outputIndex []*big.Int, challenger []common.Address) (event.Subscription, error) {

	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ValidatorPool.contract.WatchLogs(opts, "PendingBondRefunded", outputIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorPoolPendingBondRefunded)
				if err := _ValidatorPool.contract.UnpackLog(event, "PendingBondRefunded", log); err != nil {
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

// ParsePendingBondRefunded is a log parse operation binding the contract event 0x7dd3b3488bb77a5de73fe533a6fe2e06d495b6498cd57eef623053754b7a3f11.
//
// Solidity: event PendingBondRefunded(uint256 indexed outputIndex, address indexed challenger, uint128 amount)
func (_ValidatorPool *ValidatorPoolFilterer) ParsePendingBondRefunded(log types.Log) (*ValidatorPoolPendingBondRefunded, error) {
	event := new(ValidatorPoolPendingBondRefunded)
	if err := _ValidatorPool.contract.UnpackLog(event, "PendingBondRefunded", log); err != nil {
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
