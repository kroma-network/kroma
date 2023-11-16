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

// SecurityCouncilMetaData contains all meta data concerning the SecurityCouncil contract.
var SecurityCouncilMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_colosseum\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_governor\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"ConfirmationRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"outputIndex\",\"type\":\"uint256\"}],\"name\":\"DeletionRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"TransactionConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"TransactionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"TransactionSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"outputRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"l2BlockNumber\",\"type\":\"uint256\"}],\"name\":\"ValidationRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COLOSSEUM\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNOR\",\"outputs\":[{\"internalType\":\"contractUpgradeGovernor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clock\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_transactionId\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"confirmations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"confirmationCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_transactionId\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"generateTransactionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmationCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_transactionId\",\"type\":\"uint256\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_transactionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isConfirmedBy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"outputsDeleteRequested\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_outputIndex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_force\",\"type\":\"bool\"}],\"name\":\"requestDeletion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_outputRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"requestValidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_transactionId\",\"type\":\"uint256\"}],\"name\":\"revokeConfirmation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"submitTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transactionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b506040516200247538038062002475833981016040819052620000359162000076565b6001600160a01b03908116608052600160a081905260c052600060e0521661010052620000b5565b6001600160a01b03811681146200007357600080fd5b50565b600080604083850312156200008a57600080fd5b825162000097816200005d565b6020840151909250620000aa816200005d565b809150509250929050565b60805160a05160c05160e0516101005161234a6200012b600039600081816103450152818161056e015261061f01526000610daa01526000610d8101526000610d5801526000818161022701528181610726015281816107e20152818161086401528181610e1c0152610f48015261234a6000f3fe608060405234801561001057600080fd5b506004361061016c5760003560e01c80638b51d13f116100cd578063b77bf60011610081578063c01a8c8411610066578063c01a8c8414610390578063c6427474146103a3578063ee22610b146103b657600080fd5b8063b77bf60014610367578063b9774f7b1461037057600080fd5b80639ab24eb0116100b25780639ab24eb01461030a5780639ace38c21461031d5780639e45e8f41461034057600080fd5b80638b51d13f146102cb57806391ddadf4146102eb57600080fd5b806349ae963d116101245780636dc0ae22116101095780636dc0ae2214610222578063784547a71461026e5780638a8e784c1461028157600080fd5b806349ae963d146101fa57806354fd4d501461020d57600080fd5b80631703a018116101555780631703a0181461019957806320ea8d86146101b45780632a758595146101c757600080fd5b80630192337114610171578063080b91ee14610186575b600080fd5b61018461017f366004611c6c565b6103c9565b005b610184610194366004611d7b565b610607565b6101a1610703565b6040519081526020015b60405180910390f35b6101846101c2366004611dcb565b610986565b6101ea6101d5366004611dcb565b60396020526000908152604090205460ff1681565b60405190151581526020016101ab565b6101a1610208366004611e06565b610c68565b610215610d51565b6040516101ab9190611ebf565b6102497f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101ab565b6101ea61027c366004611dcb565b610df4565b6101ea61028f366004611ed9565b600082815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845260010190915290205460ff1692915050565b6101a16102d9366004611dcb565b60009081526034602052604090205490565b6102f3610e18565b60405165ffffffffffff90911681526020016101ab565b6101a1610318366004611efe565b610f44565b61033061032b366004611dcb565b61106d565b6040516101ab9493929190611f1b565b6102497f000000000000000000000000000000000000000000000000000000000000000081565b6101a160385481565b6101a161037e366004611dcb565b60346020526000908152604090205481565b61018461039e366004611dcb565b61114b565b6101a16103b1366004611e06565b6113ad565b6101846103c4366004611dcb565b611442565b3360006103d582610f44565b1161044d5760405162461bcd60e51b815260206004820152603b60248201527f546f6b656e4d756c746953696757616c6c65743a206f6e6c7920616c6c6f776560448201527f6420746f20676f7665726e616e636520746f6b656e206f776e6572000000000060648201526084015b60405180910390fd5b60008381526039602052604090205460ff1615806104685750815b6105015760405162461bcd60e51b8152602060048201526044602482018190527f5365637572697479436f756e63696c3a20746865206f75747075742068617320908201527f616c7265616479206265656e2072657175657374656420746f2062652064656c60648201527f6574656400000000000000000000000000000000000000000000000000000000608482015260a401610444565b6040805160248082018690528251808303909101815260449091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe39a219c0000000000000000000000000000000000000000000000000000000017905260006105947f000000000000000000000000000000000000000000000000000000000000000082846113ad565b905061059f8161114b565b60008581526039602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905551869183917fc63c84660a471a970585c7cab9d0601af8e717ff0822a2ea049a3542fc5aa55a9190a35050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146106b25760405162461bcd60e51b815260206004820152603c60248201527f5365637572697479436f756e63696c3a206f6e6c792074686520636f6c6f737360448201527f65756d20636f6e74726163742063616e20626520612073656e646572000000006064820152608401610444565b60006106c0336000846117b0565b604080518681526020810186905291925082917eef5106e82a682c776fd7748be042f406a9ee0feaaea86ae9029477c2b91f2a910160405180910390a250505050565b6000806001610710610e18565b61071a9190611f91565b65ffffffffffff1690507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166397c3d3346040518163ffffffff1660e01b8152600401602060405180830381865afa15801561078f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107b39190611fb8565b6040517f60c4247f000000000000000000000000000000000000000000000000000000008152600481018390527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906360c4247f90602401602060405180830381865afa15801561083e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108629190611fb8565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663fc0c546a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156108cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108f19190611fd1565b73ffffffffffffffffffffffffffffffffffffffff16638e539e8c846040518263ffffffff1660e01b815260040161092b91815260200190565b602060405180830381865afa158015610948573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061096c9190611fb8565b6109769190611fee565b610980919061202b565b91505090565b33600061099282610f44565b11610a055760405162461bcd60e51b815260206004820152603b60248201527f546f6b656e4d756c746953696757616c6c65743a206f6e6c7920616c6c6f776560448201527f6420746f20676f7665726e616e636520746f6b656e206f776e657200000000006064820152608401610444565b600082815260336020526040902054829073ffffffffffffffffffffffffffffffffffffffff16610a9e5760405162461bcd60e51b815260206004820152602f60248201527f546f6b656e4d756c746953696757616c6c65743a207472616e73616374696f6e60448201527f20646f6573206e6f7420657869737400000000000000000000000000000000006064820152608401610444565b600083815260336020526040902054839074010000000000000000000000000000000000000000900460ff1615610b3d5760405162461bcd60e51b815260206004820152602560248201527f546f6b656e4d756c746953696757616c6c65743a20616c72656164792065786560448201527f63757465640000000000000000000000000000000000000000000000000000006064820152608401610444565b600084815260346020908152604080832033845260010190915290205460ff16610bcf5760405162461bcd60e51b815260206004820152602660248201527f546f6b656e4d756c746953696757616c6c65743a206e6f7420636f6e6669726d60448201527f65642079657400000000000000000000000000000000000000000000000000006064820152608401610444565b60008481526034602090815260408083203380855260018201909352922080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055610c1c90610f44565b816000016000828254610c2f9190612066565b9091555050604051859033907f795394da21278ca39d59bb3ca00efeebdc0679acc420916c7385c2c5d942656f90600090a35050505050565b60008373ffffffffffffffffffffffffffffffffffffffff8116610cf45760405162461bcd60e51b815260206004820152602960248201527f546f6b656e4d756c746953696757616c6c65743a20616464726573732069732060448201527f6e6f742076616c696400000000000000000000000000000000000000000000006064820152608401610444565b848484610cff610e18565b604051602001610d12949392919061207d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012095945050505050565b6060610d7c7f00000000000000000000000000000000000000000000000000000000000000006119d7565b610da57f00000000000000000000000000000000000000000000000000000000000000006119d7565b610dce7f00000000000000000000000000000000000000000000000000000000000000006119d7565b604051602001610de0939291906120cb565b604051602081830303815290604052905090565b6000610dfe610703565b600092835260346020526040909220549190911015919050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663fc0c546a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e85573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ea99190611fd1565b73ffffffffffffffffffffffffffffffffffffffff166391ddadf46040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610f2d575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252610f2a91810190612141565b60015b610f3f57610f3a43611a95565b905090565b919050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663fc0c546a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fb1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fd59190611fd1565b6040517f9ab24eb000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301529190911690639ab24eb090602401602060405180830381865afa158015611043573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110679190611fb8565b92915050565b60336020526000908152604090208054600182015460028301805473ffffffffffffffffffffffffffffffffffffffff8416947401000000000000000000000000000000000000000090940460ff169391906110c890612169565b80601f01602080910402602001604051908101604052809291908181526020018280546110f490612169565b80156111415780601f1061111657610100808354040283529160200191611141565b820191906000526020600020905b81548152906001019060200180831161112457829003601f168201915b5050505050905084565b33600061115782610f44565b116111ca5760405162461bcd60e51b815260206004820152603b60248201527f546f6b656e4d756c746953696757616c6c65743a206f6e6c7920616c6c6f776560448201527f6420746f20676f7665726e616e636520746f6b656e206f776e657200000000006064820152608401610444565b600082815260336020526040902054829073ffffffffffffffffffffffffffffffffffffffff166112635760405162461bcd60e51b815260206004820152602f60248201527f546f6b656e4d756c746953696757616c6c65743a207472616e73616374696f6e60448201527f20646f6573206e6f7420657869737400000000000000000000000000000000006064820152608401610444565b6000838152603460209081526040808320338452600181019092529091205460ff16156112f85760405162461bcd60e51b815260206004820152602660248201527f546f6b656e4d756c746953696757616c6c65743a20616c726561647920636f6e60448201527f6669726d656400000000000000000000000000000000000000000000000000006064820152608401610444565b3360008181526001838101602052604090912080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016909117905561133d90610f44565b81600001600082825461135091906121bc565b9091555050604051849033907ff8a17c9136a3ae33364fac05eb088a3cbafee10c1889c88593e20ee2d8e4eb8890600090a361138a610703565b600085815260346020526040902054106113a7576113a784611442565b50505050565b60003360006113bb82610f44565b1161142e5760405162461bcd60e51b815260206004820152603b60248201527f546f6b656e4d756c746953696757616c6c65743a206f6e6c7920616c6c6f776560448201527f6420746f20676f7665726e616e636520746f6b656e206f776e657200000000006064820152608401610444565b6114398585856117b0565b95945050505050565b61144a611b17565b600081815260336020526040902054819073ffffffffffffffffffffffffffffffffffffffff166114e35760405162461bcd60e51b815260206004820152602f60248201527f546f6b656e4d756c746953696757616c6c65743a207472616e73616374696f6e60448201527f20646f6573206e6f7420657869737400000000000000000000000000000000006064820152608401610444565b600082815260336020526040902054829074010000000000000000000000000000000000000000900460ff16156115825760405162461bcd60e51b815260206004820152602560248201527f546f6b656e4d756c746953696757616c6c65743a20616c72656164792065786560448201527f63757465640000000000000000000000000000000000000000000000000000006064820152608401610444565b61158b83610df4565b6115fd5760405162461bcd60e51b815260206004820152602760248201527f546f6b656e4d756c746953696757616c6c65743a2071756f72756d206e6f742060448201527f72656163686564000000000000000000000000000000000000000000000000006064820152608401610444565b600083815260336020526040812080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff8116740100000000000000000000000000000000000000001782559091906116fe9073ffffffffffffffffffffffffffffffffffffffff165a846001015485600201805461167b90612169565b80601f01602080910402602001604051908101604052809291908181526020018280546116a790612169565b80156116f45780601f106116c9576101008083540402835291602001916116f4565b820191906000526020600020905b8154815290600101906020018083116116d757829003601f168201915b5050505050611b70565b9050806117735760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e4d756c746953696757616c6c65743a2063616c6c207472616e736160448201527f6374696f6e206661696c656400000000000000000000000000000000000000006064820152608401610444565b604051859033907f4e86ad0da28cbaaaa7e93e36c43b32696e970535225b316f1b84fbf30bdc04e890600090a3505050506117ad60018055565b50565b60008373ffffffffffffffffffffffffffffffffffffffff811661183c5760405162461bcd60e51b815260206004820152602960248201527f546f6b656e4d756c746953696757616c6c65743a20616464726573732069732060448201527f6e6f742076616c696400000000000000000000000000000000000000000000006064820152608401610444565b6000611849868686610c68565b60008181526033602052604090205490915073ffffffffffffffffffffffffffffffffffffffff16156118e45760405162461bcd60e51b815260206004820152602f60248201527f546f6b656e4d756c746953696757616c6c65743a207472616e73616374696f6e60448201527f20616c72656164792065786973747300000000000000000000000000000000006064820152608401610444565b6040805160808101825273ffffffffffffffffffffffffffffffffffffffff8089168252600060208084018281528486018b8152606086018b8152888552603390935295909220845181549351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00000000000000000000000000000000000000000090941694169390931791909117825592516001820155915190919060028201906119959082612223565b505060388054600101905550604051819033907f1f50cd00b6a6fe3928bf4a5f2f23829e9a1c9396573b828b5fa14d95aae7e77590600090a395945050505050565b606060006119e483611b8a565b600101905060008167ffffffffffffffff811115611a0457611a04611ca1565b6040519080825280601f01601f191660200182016040528015611a2e576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a8504945084611a3857509392505050565b600065ffffffffffff821115611b135760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203460448201527f38206269747300000000000000000000000000000000000000000000000000006064820152608401610444565b5090565b600260015403611b695760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610444565b6002600155565b600080600080845160208601878a8af19695505050505050565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310611bd3577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310611bff576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310611c1d57662386f26fc10000830492506010015b6305f5e1008310611c35576305f5e100830492506008015b6127108310611c4957612710830492506004015b60648310611c5b576064830492506002015b600a83106110675760010192915050565b60008060408385031215611c7f57600080fd5b8235915060208301358015158114611c9657600080fd5b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112611ce157600080fd5b813567ffffffffffffffff80821115611cfc57611cfc611ca1565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715611d4257611d42611ca1565b81604052838152866020858801011115611d5b57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600060608486031215611d9057600080fd5b8335925060208401359150604084013567ffffffffffffffff811115611db557600080fd5b611dc186828701611cd0565b9150509250925092565b600060208284031215611ddd57600080fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff811681146117ad57600080fd5b600080600060608486031215611e1b57600080fd5b8335611e2681611de4565b925060208401359150604084013567ffffffffffffffff811115611db557600080fd5b60005b83811015611e64578181015183820152602001611e4c565b838111156113a75750506000910152565b60008151808452611e8d816020860160208601611e49565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611ed26020830184611e75565b9392505050565b60008060408385031215611eec57600080fd5b823591506020830135611c9681611de4565b600060208284031215611f1057600080fd5b8135611ed281611de4565b73ffffffffffffffffffffffffffffffffffffffff851681528315156020820152826040820152608060608201526000611f586080830184611e75565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600065ffffffffffff83811690831681811015611fb057611fb0611f62565b039392505050565b600060208284031215611fca57600080fd5b5051919050565b600060208284031215611fe357600080fd5b8151611ed281611de4565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561202657612026611f62565b500290565b600082612061577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60008282101561207857612078611f62565b500390565b73ffffffffffffffffffffffffffffffffffffffff851681528360208201526080604082015260006120b26080830185611e75565b905065ffffffffffff8316606083015295945050505050565b600084516120dd818460208901611e49565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551612119816001850160208a01611e49565b60019201918201528351612134816002840160208801611e49565b0160020195945050505050565b60006020828403121561215357600080fd5b815165ffffffffffff81168114611ed257600080fd5b600181811c9082168061217d57607f821691505b6020821081036121b6577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b600082198211156121cf576121cf611f62565b500190565b601f82111561221e57600081815260208120601f850160051c810160208610156121fb5750805b601f850160051c820191505b8181101561221a57828155600101612207565b5050505b505050565b815167ffffffffffffffff81111561223d5761223d611ca1565b6122518161224b8454612169565b846121d4565b602080601f8311600181146122a4576000841561226e5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b17855561221a565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b828110156122f1578886015182559484019460019091019084016122d2565b508582101561232d57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b0190555056fea164736f6c634300080f000a",
}

// SecurityCouncilABI is the input ABI used to generate the binding from.
// Deprecated: Use SecurityCouncilMetaData.ABI instead.
var SecurityCouncilABI = SecurityCouncilMetaData.ABI

// SecurityCouncilBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SecurityCouncilMetaData.Bin instead.
var SecurityCouncilBin = SecurityCouncilMetaData.Bin

// DeploySecurityCouncil deploys a new Ethereum contract, binding an instance of SecurityCouncil to it.
func DeploySecurityCouncil(auth *bind.TransactOpts, backend bind.ContractBackend, _colosseum common.Address, _governor common.Address) (common.Address, *types.Transaction, *SecurityCouncil, error) {
	parsed, err := SecurityCouncilMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SecurityCouncilBin), backend, _colosseum, _governor)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SecurityCouncil{SecurityCouncilCaller: SecurityCouncilCaller{contract: contract}, SecurityCouncilTransactor: SecurityCouncilTransactor{contract: contract}, SecurityCouncilFilterer: SecurityCouncilFilterer{contract: contract}}, nil
}

// SecurityCouncil is an auto generated Go binding around an Ethereum contract.
type SecurityCouncil struct {
	SecurityCouncilCaller     // Read-only binding to the contract
	SecurityCouncilTransactor // Write-only binding to the contract
	SecurityCouncilFilterer   // Log filterer for contract events
}

// SecurityCouncilCaller is an auto generated read-only Go binding around an Ethereum contract.
type SecurityCouncilCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecurityCouncilTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SecurityCouncilTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecurityCouncilFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SecurityCouncilFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecurityCouncilSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SecurityCouncilSession struct {
	Contract     *SecurityCouncil  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SecurityCouncilCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SecurityCouncilCallerSession struct {
	Contract *SecurityCouncilCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SecurityCouncilTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SecurityCouncilTransactorSession struct {
	Contract     *SecurityCouncilTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SecurityCouncilRaw is an auto generated low-level Go binding around an Ethereum contract.
type SecurityCouncilRaw struct {
	Contract *SecurityCouncil // Generic contract binding to access the raw methods on
}

// SecurityCouncilCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SecurityCouncilCallerRaw struct {
	Contract *SecurityCouncilCaller // Generic read-only contract binding to access the raw methods on
}

// SecurityCouncilTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SecurityCouncilTransactorRaw struct {
	Contract *SecurityCouncilTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSecurityCouncil creates a new instance of SecurityCouncil, bound to a specific deployed contract.
func NewSecurityCouncil(address common.Address, backend bind.ContractBackend) (*SecurityCouncil, error) {
	contract, err := bindSecurityCouncil(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SecurityCouncil{SecurityCouncilCaller: SecurityCouncilCaller{contract: contract}, SecurityCouncilTransactor: SecurityCouncilTransactor{contract: contract}, SecurityCouncilFilterer: SecurityCouncilFilterer{contract: contract}}, nil
}

// NewSecurityCouncilCaller creates a new read-only instance of SecurityCouncil, bound to a specific deployed contract.
func NewSecurityCouncilCaller(address common.Address, caller bind.ContractCaller) (*SecurityCouncilCaller, error) {
	contract, err := bindSecurityCouncil(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SecurityCouncilCaller{contract: contract}, nil
}

// NewSecurityCouncilTransactor creates a new write-only instance of SecurityCouncil, bound to a specific deployed contract.
func NewSecurityCouncilTransactor(address common.Address, transactor bind.ContractTransactor) (*SecurityCouncilTransactor, error) {
	contract, err := bindSecurityCouncil(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SecurityCouncilTransactor{contract: contract}, nil
}

// NewSecurityCouncilFilterer creates a new log filterer instance of SecurityCouncil, bound to a specific deployed contract.
func NewSecurityCouncilFilterer(address common.Address, filterer bind.ContractFilterer) (*SecurityCouncilFilterer, error) {
	contract, err := bindSecurityCouncil(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SecurityCouncilFilterer{contract: contract}, nil
}

// bindSecurityCouncil binds a generic wrapper to an already deployed contract.
func bindSecurityCouncil(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SecurityCouncilMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SecurityCouncil *SecurityCouncilRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SecurityCouncil.Contract.SecurityCouncilCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SecurityCouncil *SecurityCouncilRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SecurityCouncil.Contract.SecurityCouncilTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SecurityCouncil *SecurityCouncilRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SecurityCouncil.Contract.SecurityCouncilTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SecurityCouncil *SecurityCouncilCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SecurityCouncil.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SecurityCouncil *SecurityCouncilTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SecurityCouncil.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SecurityCouncil *SecurityCouncilTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SecurityCouncil.Contract.contract.Transact(opts, method, params...)
}

// COLOSSEUM is a free data retrieval call binding the contract method 0x9e45e8f4.
//
// Solidity: function COLOSSEUM() view returns(address)
func (_SecurityCouncil *SecurityCouncilCaller) COLOSSEUM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SecurityCouncil.contract.Call(opts, &out, "COLOSSEUM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// COLOSSEUM is a free data retrieval call binding the contract method 0x9e45e8f4.
//
// Solidity: function COLOSSEUM() view returns(address)
func (_SecurityCouncil *SecurityCouncilSession) COLOSSEUM() (common.Address, error) {
	return _SecurityCouncil.Contract.COLOSSEUM(&_SecurityCouncil.CallOpts)
}

// COLOSSEUM is a free data retrieval call binding the contract method 0x9e45e8f4.
//
// Solidity: function COLOSSEUM() view returns(address)
func (_SecurityCouncil *SecurityCouncilCallerSession) COLOSSEUM() (common.Address, error) {
	return _SecurityCouncil.Contract.COLOSSEUM(&_SecurityCouncil.CallOpts)
}

// GOVERNOR is a free data retrieval call binding the contract method 0x6dc0ae22.
//
// Solidity: function GOVERNOR() view returns(address)
func (_SecurityCouncil *SecurityCouncilCaller) GOVERNOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SecurityCouncil.contract.Call(opts, &out, "GOVERNOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVERNOR is a free data retrieval call binding the contract method 0x6dc0ae22.
//
// Solidity: function GOVERNOR() view returns(address)
func (_SecurityCouncil *SecurityCouncilSession) GOVERNOR() (common.Address, error) {
	return _SecurityCouncil.Contract.GOVERNOR(&_SecurityCouncil.CallOpts)
}

// GOVERNOR is a free data retrieval call binding the contract method 0x6dc0ae22.
//
// Solidity: function GOVERNOR() view returns(address)
func (_SecurityCouncil *SecurityCouncilCallerSession) GOVERNOR() (common.Address, error) {
	return _SecurityCouncil.Contract.GOVERNOR(&_SecurityCouncil.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_SecurityCouncil *SecurityCouncilCaller) Clock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SecurityCouncil.contract.Call(opts, &out, "clock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_SecurityCouncil *SecurityCouncilSession) Clock() (*big.Int, error) {
	return _SecurityCouncil.Contract.Clock(&_SecurityCouncil.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_SecurityCouncil *SecurityCouncilCallerSession) Clock() (*big.Int, error) {
	return _SecurityCouncil.Contract.Clock(&_SecurityCouncil.CallOpts)
}

// Confirmations is a free data retrieval call binding the contract method 0xb9774f7b.
//
// Solidity: function confirmations(uint256 ) view returns(uint256 confirmationCount)
func (_SecurityCouncil *SecurityCouncilCaller) Confirmations(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SecurityCouncil.contract.Call(opts, &out, "confirmations", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Confirmations is a free data retrieval call binding the contract method 0xb9774f7b.
//
// Solidity: function confirmations(uint256 ) view returns(uint256 confirmationCount)
func (_SecurityCouncil *SecurityCouncilSession) Confirmations(arg0 *big.Int) (*big.Int, error) {
	return _SecurityCouncil.Contract.Confirmations(&_SecurityCouncil.CallOpts, arg0)
}

// Confirmations is a free data retrieval call binding the contract method 0xb9774f7b.
//
// Solidity: function confirmations(uint256 ) view returns(uint256 confirmationCount)
func (_SecurityCouncil *SecurityCouncilCallerSession) Confirmations(arg0 *big.Int) (*big.Int, error) {
	return _SecurityCouncil.Contract.Confirmations(&_SecurityCouncil.CallOpts, arg0)
}

// GenerateTransactionId is a free data retrieval call binding the contract method 0x49ae963d.
//
// Solidity: function generateTransactionId(address _target, uint256 _value, bytes _data) view returns(uint256)
func (_SecurityCouncil *SecurityCouncilCaller) GenerateTransactionId(opts *bind.CallOpts, _target common.Address, _value *big.Int, _data []byte) (*big.Int, error) {
	var out []interface{}
	err := _SecurityCouncil.contract.Call(opts, &out, "generateTransactionId", _target, _value, _data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenerateTransactionId is a free data retrieval call binding the contract method 0x49ae963d.
//
// Solidity: function generateTransactionId(address _target, uint256 _value, bytes _data) view returns(uint256)
func (_SecurityCouncil *SecurityCouncilSession) GenerateTransactionId(_target common.Address, _value *big.Int, _data []byte) (*big.Int, error) {
	return _SecurityCouncil.Contract.GenerateTransactionId(&_SecurityCouncil.CallOpts, _target, _value, _data)
}

// GenerateTransactionId is a free data retrieval call binding the contract method 0x49ae963d.
//
// Solidity: function generateTransactionId(address _target, uint256 _value, bytes _data) view returns(uint256)
func (_SecurityCouncil *SecurityCouncilCallerSession) GenerateTransactionId(_target common.Address, _value *big.Int, _data []byte) (*big.Int, error) {
	return _SecurityCouncil.Contract.GenerateTransactionId(&_SecurityCouncil.CallOpts, _target, _value, _data)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 _transactionId) view returns(uint256)
func (_SecurityCouncil *SecurityCouncilCaller) GetConfirmationCount(opts *bind.CallOpts, _transactionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SecurityCouncil.contract.Call(opts, &out, "getConfirmationCount", _transactionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 _transactionId) view returns(uint256)
func (_SecurityCouncil *SecurityCouncilSession) GetConfirmationCount(_transactionId *big.Int) (*big.Int, error) {
	return _SecurityCouncil.Contract.GetConfirmationCount(&_SecurityCouncil.CallOpts, _transactionId)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 _transactionId) view returns(uint256)
func (_SecurityCouncil *SecurityCouncilCallerSession) GetConfirmationCount(_transactionId *big.Int) (*big.Int, error) {
	return _SecurityCouncil.Contract.GetConfirmationCount(&_SecurityCouncil.CallOpts, _transactionId)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_SecurityCouncil *SecurityCouncilCaller) GetVotes(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SecurityCouncil.contract.Call(opts, &out, "getVotes", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_SecurityCouncil *SecurityCouncilSession) GetVotes(account common.Address) (*big.Int, error) {
	return _SecurityCouncil.Contract.GetVotes(&_SecurityCouncil.CallOpts, account)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_SecurityCouncil *SecurityCouncilCallerSession) GetVotes(account common.Address) (*big.Int, error) {
	return _SecurityCouncil.Contract.GetVotes(&_SecurityCouncil.CallOpts, account)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 _transactionId) view returns(bool)
func (_SecurityCouncil *SecurityCouncilCaller) IsConfirmed(opts *bind.CallOpts, _transactionId *big.Int) (bool, error) {
	var out []interface{}
	err := _SecurityCouncil.contract.Call(opts, &out, "isConfirmed", _transactionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 _transactionId) view returns(bool)
func (_SecurityCouncil *SecurityCouncilSession) IsConfirmed(_transactionId *big.Int) (bool, error) {
	return _SecurityCouncil.Contract.IsConfirmed(&_SecurityCouncil.CallOpts, _transactionId)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 _transactionId) view returns(bool)
func (_SecurityCouncil *SecurityCouncilCallerSession) IsConfirmed(_transactionId *big.Int) (bool, error) {
	return _SecurityCouncil.Contract.IsConfirmed(&_SecurityCouncil.CallOpts, _transactionId)
}

// IsConfirmedBy is a free data retrieval call binding the contract method 0x8a8e784c.
//
// Solidity: function isConfirmedBy(uint256 _transactionId, address _account) view returns(bool)
func (_SecurityCouncil *SecurityCouncilCaller) IsConfirmedBy(opts *bind.CallOpts, _transactionId *big.Int, _account common.Address) (bool, error) {
	var out []interface{}
	err := _SecurityCouncil.contract.Call(opts, &out, "isConfirmedBy", _transactionId, _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConfirmedBy is a free data retrieval call binding the contract method 0x8a8e784c.
//
// Solidity: function isConfirmedBy(uint256 _transactionId, address _account) view returns(bool)
func (_SecurityCouncil *SecurityCouncilSession) IsConfirmedBy(_transactionId *big.Int, _account common.Address) (bool, error) {
	return _SecurityCouncil.Contract.IsConfirmedBy(&_SecurityCouncil.CallOpts, _transactionId, _account)
}

// IsConfirmedBy is a free data retrieval call binding the contract method 0x8a8e784c.
//
// Solidity: function isConfirmedBy(uint256 _transactionId, address _account) view returns(bool)
func (_SecurityCouncil *SecurityCouncilCallerSession) IsConfirmedBy(_transactionId *big.Int, _account common.Address) (bool, error) {
	return _SecurityCouncil.Contract.IsConfirmedBy(&_SecurityCouncil.CallOpts, _transactionId, _account)
}

// OutputsDeleteRequested is a free data retrieval call binding the contract method 0x2a758595.
//
// Solidity: function outputsDeleteRequested(uint256 ) view returns(bool)
func (_SecurityCouncil *SecurityCouncilCaller) OutputsDeleteRequested(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _SecurityCouncil.contract.Call(opts, &out, "outputsDeleteRequested", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OutputsDeleteRequested is a free data retrieval call binding the contract method 0x2a758595.
//
// Solidity: function outputsDeleteRequested(uint256 ) view returns(bool)
func (_SecurityCouncil *SecurityCouncilSession) OutputsDeleteRequested(arg0 *big.Int) (bool, error) {
	return _SecurityCouncil.Contract.OutputsDeleteRequested(&_SecurityCouncil.CallOpts, arg0)
}

// OutputsDeleteRequested is a free data retrieval call binding the contract method 0x2a758595.
//
// Solidity: function outputsDeleteRequested(uint256 ) view returns(bool)
func (_SecurityCouncil *SecurityCouncilCallerSession) OutputsDeleteRequested(arg0 *big.Int) (bool, error) {
	return _SecurityCouncil.Contract.OutputsDeleteRequested(&_SecurityCouncil.CallOpts, arg0)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint256)
func (_SecurityCouncil *SecurityCouncilCaller) Quorum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SecurityCouncil.contract.Call(opts, &out, "quorum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint256)
func (_SecurityCouncil *SecurityCouncilSession) Quorum() (*big.Int, error) {
	return _SecurityCouncil.Contract.Quorum(&_SecurityCouncil.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint256)
func (_SecurityCouncil *SecurityCouncilCallerSession) Quorum() (*big.Int, error) {
	return _SecurityCouncil.Contract.Quorum(&_SecurityCouncil.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_SecurityCouncil *SecurityCouncilCaller) TransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SecurityCouncil.contract.Call(opts, &out, "transactionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_SecurityCouncil *SecurityCouncilSession) TransactionCount() (*big.Int, error) {
	return _SecurityCouncil.Contract.TransactionCount(&_SecurityCouncil.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_SecurityCouncil *SecurityCouncilCallerSession) TransactionCount() (*big.Int, error) {
	return _SecurityCouncil.Contract.TransactionCount(&_SecurityCouncil.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address target, bool executed, uint256 value, bytes data)
func (_SecurityCouncil *SecurityCouncilCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Target   common.Address
	Executed bool
	Value    *big.Int
	Data     []byte
}, error) {
	var out []interface{}
	err := _SecurityCouncil.contract.Call(opts, &out, "transactions", arg0)

	outstruct := new(struct {
		Target   common.Address
		Executed bool
		Value    *big.Int
		Data     []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Target = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Executed = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Value = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address target, bool executed, uint256 value, bytes data)
func (_SecurityCouncil *SecurityCouncilSession) Transactions(arg0 *big.Int) (struct {
	Target   common.Address
	Executed bool
	Value    *big.Int
	Data     []byte
}, error) {
	return _SecurityCouncil.Contract.Transactions(&_SecurityCouncil.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address target, bool executed, uint256 value, bytes data)
func (_SecurityCouncil *SecurityCouncilCallerSession) Transactions(arg0 *big.Int) (struct {
	Target   common.Address
	Executed bool
	Value    *big.Int
	Data     []byte
}, error) {
	return _SecurityCouncil.Contract.Transactions(&_SecurityCouncil.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SecurityCouncil *SecurityCouncilCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SecurityCouncil.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SecurityCouncil *SecurityCouncilSession) Version() (string, error) {
	return _SecurityCouncil.Contract.Version(&_SecurityCouncil.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SecurityCouncil *SecurityCouncilCallerSession) Version() (string, error) {
	return _SecurityCouncil.Contract.Version(&_SecurityCouncil.CallOpts)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 _transactionId) returns()
func (_SecurityCouncil *SecurityCouncilTransactor) ConfirmTransaction(opts *bind.TransactOpts, _transactionId *big.Int) (*types.Transaction, error) {
	return _SecurityCouncil.contract.Transact(opts, "confirmTransaction", _transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 _transactionId) returns()
func (_SecurityCouncil *SecurityCouncilSession) ConfirmTransaction(_transactionId *big.Int) (*types.Transaction, error) {
	return _SecurityCouncil.Contract.ConfirmTransaction(&_SecurityCouncil.TransactOpts, _transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 _transactionId) returns()
func (_SecurityCouncil *SecurityCouncilTransactorSession) ConfirmTransaction(_transactionId *big.Int) (*types.Transaction, error) {
	return _SecurityCouncil.Contract.ConfirmTransaction(&_SecurityCouncil.TransactOpts, _transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 _transactionId) returns()
func (_SecurityCouncil *SecurityCouncilTransactor) ExecuteTransaction(opts *bind.TransactOpts, _transactionId *big.Int) (*types.Transaction, error) {
	return _SecurityCouncil.contract.Transact(opts, "executeTransaction", _transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 _transactionId) returns()
func (_SecurityCouncil *SecurityCouncilSession) ExecuteTransaction(_transactionId *big.Int) (*types.Transaction, error) {
	return _SecurityCouncil.Contract.ExecuteTransaction(&_SecurityCouncil.TransactOpts, _transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 _transactionId) returns()
func (_SecurityCouncil *SecurityCouncilTransactorSession) ExecuteTransaction(_transactionId *big.Int) (*types.Transaction, error) {
	return _SecurityCouncil.Contract.ExecuteTransaction(&_SecurityCouncil.TransactOpts, _transactionId)
}

// RequestDeletion is a paid mutator transaction binding the contract method 0x01923371.
//
// Solidity: function requestDeletion(uint256 _outputIndex, bool _force) returns()
func (_SecurityCouncil *SecurityCouncilTransactor) RequestDeletion(opts *bind.TransactOpts, _outputIndex *big.Int, _force bool) (*types.Transaction, error) {
	return _SecurityCouncil.contract.Transact(opts, "requestDeletion", _outputIndex, _force)
}

// RequestDeletion is a paid mutator transaction binding the contract method 0x01923371.
//
// Solidity: function requestDeletion(uint256 _outputIndex, bool _force) returns()
func (_SecurityCouncil *SecurityCouncilSession) RequestDeletion(_outputIndex *big.Int, _force bool) (*types.Transaction, error) {
	return _SecurityCouncil.Contract.RequestDeletion(&_SecurityCouncil.TransactOpts, _outputIndex, _force)
}

// RequestDeletion is a paid mutator transaction binding the contract method 0x01923371.
//
// Solidity: function requestDeletion(uint256 _outputIndex, bool _force) returns()
func (_SecurityCouncil *SecurityCouncilTransactorSession) RequestDeletion(_outputIndex *big.Int, _force bool) (*types.Transaction, error) {
	return _SecurityCouncil.Contract.RequestDeletion(&_SecurityCouncil.TransactOpts, _outputIndex, _force)
}

// RequestValidation is a paid mutator transaction binding the contract method 0x080b91ee.
//
// Solidity: function requestValidation(bytes32 _outputRoot, uint256 _l2BlockNumber, bytes _data) returns()
func (_SecurityCouncil *SecurityCouncilTransactor) RequestValidation(opts *bind.TransactOpts, _outputRoot [32]byte, _l2BlockNumber *big.Int, _data []byte) (*types.Transaction, error) {
	return _SecurityCouncil.contract.Transact(opts, "requestValidation", _outputRoot, _l2BlockNumber, _data)
}

// RequestValidation is a paid mutator transaction binding the contract method 0x080b91ee.
//
// Solidity: function requestValidation(bytes32 _outputRoot, uint256 _l2BlockNumber, bytes _data) returns()
func (_SecurityCouncil *SecurityCouncilSession) RequestValidation(_outputRoot [32]byte, _l2BlockNumber *big.Int, _data []byte) (*types.Transaction, error) {
	return _SecurityCouncil.Contract.RequestValidation(&_SecurityCouncil.TransactOpts, _outputRoot, _l2BlockNumber, _data)
}

// RequestValidation is a paid mutator transaction binding the contract method 0x080b91ee.
//
// Solidity: function requestValidation(bytes32 _outputRoot, uint256 _l2BlockNumber, bytes _data) returns()
func (_SecurityCouncil *SecurityCouncilTransactorSession) RequestValidation(_outputRoot [32]byte, _l2BlockNumber *big.Int, _data []byte) (*types.Transaction, error) {
	return _SecurityCouncil.Contract.RequestValidation(&_SecurityCouncil.TransactOpts, _outputRoot, _l2BlockNumber, _data)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 _transactionId) returns()
func (_SecurityCouncil *SecurityCouncilTransactor) RevokeConfirmation(opts *bind.TransactOpts, _transactionId *big.Int) (*types.Transaction, error) {
	return _SecurityCouncil.contract.Transact(opts, "revokeConfirmation", _transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 _transactionId) returns()
func (_SecurityCouncil *SecurityCouncilSession) RevokeConfirmation(_transactionId *big.Int) (*types.Transaction, error) {
	return _SecurityCouncil.Contract.RevokeConfirmation(&_SecurityCouncil.TransactOpts, _transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 _transactionId) returns()
func (_SecurityCouncil *SecurityCouncilTransactorSession) RevokeConfirmation(_transactionId *big.Int) (*types.Transaction, error) {
	return _SecurityCouncil.Contract.RevokeConfirmation(&_SecurityCouncil.TransactOpts, _transactionId)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address _target, uint256 _value, bytes _data) returns(uint256)
func (_SecurityCouncil *SecurityCouncilTransactor) SubmitTransaction(opts *bind.TransactOpts, _target common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _SecurityCouncil.contract.Transact(opts, "submitTransaction", _target, _value, _data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address _target, uint256 _value, bytes _data) returns(uint256)
func (_SecurityCouncil *SecurityCouncilSession) SubmitTransaction(_target common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _SecurityCouncil.Contract.SubmitTransaction(&_SecurityCouncil.TransactOpts, _target, _value, _data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address _target, uint256 _value, bytes _data) returns(uint256)
func (_SecurityCouncil *SecurityCouncilTransactorSession) SubmitTransaction(_target common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _SecurityCouncil.Contract.SubmitTransaction(&_SecurityCouncil.TransactOpts, _target, _value, _data)
}

// SecurityCouncilConfirmationRevokedIterator is returned from FilterConfirmationRevoked and is used to iterate over the raw logs and unpacked data for ConfirmationRevoked events raised by the SecurityCouncil contract.
type SecurityCouncilConfirmationRevokedIterator struct {
	Event *SecurityCouncilConfirmationRevoked // Event containing the contract specifics and raw log

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
func (it *SecurityCouncilConfirmationRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityCouncilConfirmationRevoked)
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
		it.Event = new(SecurityCouncilConfirmationRevoked)
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
func (it *SecurityCouncilConfirmationRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SecurityCouncilConfirmationRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SecurityCouncilConfirmationRevoked represents a ConfirmationRevoked event raised by the SecurityCouncil contract.
type SecurityCouncilConfirmationRevoked struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConfirmationRevoked is a free log retrieval operation binding the contract event 0x795394da21278ca39d59bb3ca00efeebdc0679acc420916c7385c2c5d942656f.
//
// Solidity: event ConfirmationRevoked(address indexed sender, uint256 indexed transactionId)
func (_SecurityCouncil *SecurityCouncilFilterer) FilterConfirmationRevoked(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*SecurityCouncilConfirmationRevokedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SecurityCouncil.contract.FilterLogs(opts, "ConfirmationRevoked", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SecurityCouncilConfirmationRevokedIterator{contract: _SecurityCouncil.contract, event: "ConfirmationRevoked", logs: logs, sub: sub}, nil
}

// WatchConfirmationRevoked is a free log subscription operation binding the contract event 0x795394da21278ca39d59bb3ca00efeebdc0679acc420916c7385c2c5d942656f.
//
// Solidity: event ConfirmationRevoked(address indexed sender, uint256 indexed transactionId)
func (_SecurityCouncil *SecurityCouncilFilterer) WatchConfirmationRevoked(opts *bind.WatchOpts, sink chan<- *SecurityCouncilConfirmationRevoked, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SecurityCouncil.contract.WatchLogs(opts, "ConfirmationRevoked", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SecurityCouncilConfirmationRevoked)
				if err := _SecurityCouncil.contract.UnpackLog(event, "ConfirmationRevoked", log); err != nil {
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

// ParseConfirmationRevoked is a log parse operation binding the contract event 0x795394da21278ca39d59bb3ca00efeebdc0679acc420916c7385c2c5d942656f.
//
// Solidity: event ConfirmationRevoked(address indexed sender, uint256 indexed transactionId)
func (_SecurityCouncil *SecurityCouncilFilterer) ParseConfirmationRevoked(log types.Log) (*SecurityCouncilConfirmationRevoked, error) {
	event := new(SecurityCouncilConfirmationRevoked)
	if err := _SecurityCouncil.contract.UnpackLog(event, "ConfirmationRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SecurityCouncilDeletionRequestedIterator is returned from FilterDeletionRequested and is used to iterate over the raw logs and unpacked data for DeletionRequested events raised by the SecurityCouncil contract.
type SecurityCouncilDeletionRequestedIterator struct {
	Event *SecurityCouncilDeletionRequested // Event containing the contract specifics and raw log

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
func (it *SecurityCouncilDeletionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityCouncilDeletionRequested)
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
		it.Event = new(SecurityCouncilDeletionRequested)
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
func (it *SecurityCouncilDeletionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SecurityCouncilDeletionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SecurityCouncilDeletionRequested represents a DeletionRequested event raised by the SecurityCouncil contract.
type SecurityCouncilDeletionRequested struct {
	TransactionId *big.Int
	OutputIndex   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeletionRequested is a free log retrieval operation binding the contract event 0xc63c84660a471a970585c7cab9d0601af8e717ff0822a2ea049a3542fc5aa55a.
//
// Solidity: event DeletionRequested(uint256 indexed transactionId, uint256 indexed outputIndex)
func (_SecurityCouncil *SecurityCouncilFilterer) FilterDeletionRequested(opts *bind.FilterOpts, transactionId []*big.Int, outputIndex []*big.Int) (*SecurityCouncilDeletionRequestedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}

	logs, sub, err := _SecurityCouncil.contract.FilterLogs(opts, "DeletionRequested", transactionIdRule, outputIndexRule)
	if err != nil {
		return nil, err
	}
	return &SecurityCouncilDeletionRequestedIterator{contract: _SecurityCouncil.contract, event: "DeletionRequested", logs: logs, sub: sub}, nil
}

// WatchDeletionRequested is a free log subscription operation binding the contract event 0xc63c84660a471a970585c7cab9d0601af8e717ff0822a2ea049a3542fc5aa55a.
//
// Solidity: event DeletionRequested(uint256 indexed transactionId, uint256 indexed outputIndex)
func (_SecurityCouncil *SecurityCouncilFilterer) WatchDeletionRequested(opts *bind.WatchOpts, sink chan<- *SecurityCouncilDeletionRequested, transactionId []*big.Int, outputIndex []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var outputIndexRule []interface{}
	for _, outputIndexItem := range outputIndex {
		outputIndexRule = append(outputIndexRule, outputIndexItem)
	}

	logs, sub, err := _SecurityCouncil.contract.WatchLogs(opts, "DeletionRequested", transactionIdRule, outputIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SecurityCouncilDeletionRequested)
				if err := _SecurityCouncil.contract.UnpackLog(event, "DeletionRequested", log); err != nil {
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

// ParseDeletionRequested is a log parse operation binding the contract event 0xc63c84660a471a970585c7cab9d0601af8e717ff0822a2ea049a3542fc5aa55a.
//
// Solidity: event DeletionRequested(uint256 indexed transactionId, uint256 indexed outputIndex)
func (_SecurityCouncil *SecurityCouncilFilterer) ParseDeletionRequested(log types.Log) (*SecurityCouncilDeletionRequested, error) {
	event := new(SecurityCouncilDeletionRequested)
	if err := _SecurityCouncil.contract.UnpackLog(event, "DeletionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SecurityCouncilInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SecurityCouncil contract.
type SecurityCouncilInitializedIterator struct {
	Event *SecurityCouncilInitialized // Event containing the contract specifics and raw log

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
func (it *SecurityCouncilInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityCouncilInitialized)
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
		it.Event = new(SecurityCouncilInitialized)
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
func (it *SecurityCouncilInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SecurityCouncilInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SecurityCouncilInitialized represents a Initialized event raised by the SecurityCouncil contract.
type SecurityCouncilInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SecurityCouncil *SecurityCouncilFilterer) FilterInitialized(opts *bind.FilterOpts) (*SecurityCouncilInitializedIterator, error) {

	logs, sub, err := _SecurityCouncil.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SecurityCouncilInitializedIterator{contract: _SecurityCouncil.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SecurityCouncil *SecurityCouncilFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SecurityCouncilInitialized) (event.Subscription, error) {

	logs, sub, err := _SecurityCouncil.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SecurityCouncilInitialized)
				if err := _SecurityCouncil.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SecurityCouncil *SecurityCouncilFilterer) ParseInitialized(log types.Log) (*SecurityCouncilInitialized, error) {
	event := new(SecurityCouncilInitialized)
	if err := _SecurityCouncil.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SecurityCouncilTransactionConfirmedIterator is returned from FilterTransactionConfirmed and is used to iterate over the raw logs and unpacked data for TransactionConfirmed events raised by the SecurityCouncil contract.
type SecurityCouncilTransactionConfirmedIterator struct {
	Event *SecurityCouncilTransactionConfirmed // Event containing the contract specifics and raw log

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
func (it *SecurityCouncilTransactionConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityCouncilTransactionConfirmed)
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
		it.Event = new(SecurityCouncilTransactionConfirmed)
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
func (it *SecurityCouncilTransactionConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SecurityCouncilTransactionConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SecurityCouncilTransactionConfirmed represents a TransactionConfirmed event raised by the SecurityCouncil contract.
type SecurityCouncilTransactionConfirmed struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTransactionConfirmed is a free log retrieval operation binding the contract event 0xf8a17c9136a3ae33364fac05eb088a3cbafee10c1889c88593e20ee2d8e4eb88.
//
// Solidity: event TransactionConfirmed(address indexed sender, uint256 indexed transactionId)
func (_SecurityCouncil *SecurityCouncilFilterer) FilterTransactionConfirmed(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*SecurityCouncilTransactionConfirmedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SecurityCouncil.contract.FilterLogs(opts, "TransactionConfirmed", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SecurityCouncilTransactionConfirmedIterator{contract: _SecurityCouncil.contract, event: "TransactionConfirmed", logs: logs, sub: sub}, nil
}

// WatchTransactionConfirmed is a free log subscription operation binding the contract event 0xf8a17c9136a3ae33364fac05eb088a3cbafee10c1889c88593e20ee2d8e4eb88.
//
// Solidity: event TransactionConfirmed(address indexed sender, uint256 indexed transactionId)
func (_SecurityCouncil *SecurityCouncilFilterer) WatchTransactionConfirmed(opts *bind.WatchOpts, sink chan<- *SecurityCouncilTransactionConfirmed, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SecurityCouncil.contract.WatchLogs(opts, "TransactionConfirmed", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SecurityCouncilTransactionConfirmed)
				if err := _SecurityCouncil.contract.UnpackLog(event, "TransactionConfirmed", log); err != nil {
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

// ParseTransactionConfirmed is a log parse operation binding the contract event 0xf8a17c9136a3ae33364fac05eb088a3cbafee10c1889c88593e20ee2d8e4eb88.
//
// Solidity: event TransactionConfirmed(address indexed sender, uint256 indexed transactionId)
func (_SecurityCouncil *SecurityCouncilFilterer) ParseTransactionConfirmed(log types.Log) (*SecurityCouncilTransactionConfirmed, error) {
	event := new(SecurityCouncilTransactionConfirmed)
	if err := _SecurityCouncil.contract.UnpackLog(event, "TransactionConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SecurityCouncilTransactionExecutedIterator is returned from FilterTransactionExecuted and is used to iterate over the raw logs and unpacked data for TransactionExecuted events raised by the SecurityCouncil contract.
type SecurityCouncilTransactionExecutedIterator struct {
	Event *SecurityCouncilTransactionExecuted // Event containing the contract specifics and raw log

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
func (it *SecurityCouncilTransactionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityCouncilTransactionExecuted)
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
		it.Event = new(SecurityCouncilTransactionExecuted)
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
func (it *SecurityCouncilTransactionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SecurityCouncilTransactionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SecurityCouncilTransactionExecuted represents a TransactionExecuted event raised by the SecurityCouncil contract.
type SecurityCouncilTransactionExecuted struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTransactionExecuted is a free log retrieval operation binding the contract event 0x4e86ad0da28cbaaaa7e93e36c43b32696e970535225b316f1b84fbf30bdc04e8.
//
// Solidity: event TransactionExecuted(address indexed sender, uint256 indexed transactionId)
func (_SecurityCouncil *SecurityCouncilFilterer) FilterTransactionExecuted(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*SecurityCouncilTransactionExecutedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SecurityCouncil.contract.FilterLogs(opts, "TransactionExecuted", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SecurityCouncilTransactionExecutedIterator{contract: _SecurityCouncil.contract, event: "TransactionExecuted", logs: logs, sub: sub}, nil
}

// WatchTransactionExecuted is a free log subscription operation binding the contract event 0x4e86ad0da28cbaaaa7e93e36c43b32696e970535225b316f1b84fbf30bdc04e8.
//
// Solidity: event TransactionExecuted(address indexed sender, uint256 indexed transactionId)
func (_SecurityCouncil *SecurityCouncilFilterer) WatchTransactionExecuted(opts *bind.WatchOpts, sink chan<- *SecurityCouncilTransactionExecuted, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SecurityCouncil.contract.WatchLogs(opts, "TransactionExecuted", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SecurityCouncilTransactionExecuted)
				if err := _SecurityCouncil.contract.UnpackLog(event, "TransactionExecuted", log); err != nil {
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

// ParseTransactionExecuted is a log parse operation binding the contract event 0x4e86ad0da28cbaaaa7e93e36c43b32696e970535225b316f1b84fbf30bdc04e8.
//
// Solidity: event TransactionExecuted(address indexed sender, uint256 indexed transactionId)
func (_SecurityCouncil *SecurityCouncilFilterer) ParseTransactionExecuted(log types.Log) (*SecurityCouncilTransactionExecuted, error) {
	event := new(SecurityCouncilTransactionExecuted)
	if err := _SecurityCouncil.contract.UnpackLog(event, "TransactionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SecurityCouncilTransactionSubmittedIterator is returned from FilterTransactionSubmitted and is used to iterate over the raw logs and unpacked data for TransactionSubmitted events raised by the SecurityCouncil contract.
type SecurityCouncilTransactionSubmittedIterator struct {
	Event *SecurityCouncilTransactionSubmitted // Event containing the contract specifics and raw log

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
func (it *SecurityCouncilTransactionSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityCouncilTransactionSubmitted)
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
		it.Event = new(SecurityCouncilTransactionSubmitted)
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
func (it *SecurityCouncilTransactionSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SecurityCouncilTransactionSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SecurityCouncilTransactionSubmitted represents a TransactionSubmitted event raised by the SecurityCouncil contract.
type SecurityCouncilTransactionSubmitted struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTransactionSubmitted is a free log retrieval operation binding the contract event 0x1f50cd00b6a6fe3928bf4a5f2f23829e9a1c9396573b828b5fa14d95aae7e775.
//
// Solidity: event TransactionSubmitted(address indexed sender, uint256 indexed transactionId)
func (_SecurityCouncil *SecurityCouncilFilterer) FilterTransactionSubmitted(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*SecurityCouncilTransactionSubmittedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SecurityCouncil.contract.FilterLogs(opts, "TransactionSubmitted", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SecurityCouncilTransactionSubmittedIterator{contract: _SecurityCouncil.contract, event: "TransactionSubmitted", logs: logs, sub: sub}, nil
}

// WatchTransactionSubmitted is a free log subscription operation binding the contract event 0x1f50cd00b6a6fe3928bf4a5f2f23829e9a1c9396573b828b5fa14d95aae7e775.
//
// Solidity: event TransactionSubmitted(address indexed sender, uint256 indexed transactionId)
func (_SecurityCouncil *SecurityCouncilFilterer) WatchTransactionSubmitted(opts *bind.WatchOpts, sink chan<- *SecurityCouncilTransactionSubmitted, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SecurityCouncil.contract.WatchLogs(opts, "TransactionSubmitted", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SecurityCouncilTransactionSubmitted)
				if err := _SecurityCouncil.contract.UnpackLog(event, "TransactionSubmitted", log); err != nil {
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

// ParseTransactionSubmitted is a log parse operation binding the contract event 0x1f50cd00b6a6fe3928bf4a5f2f23829e9a1c9396573b828b5fa14d95aae7e775.
//
// Solidity: event TransactionSubmitted(address indexed sender, uint256 indexed transactionId)
func (_SecurityCouncil *SecurityCouncilFilterer) ParseTransactionSubmitted(log types.Log) (*SecurityCouncilTransactionSubmitted, error) {
	event := new(SecurityCouncilTransactionSubmitted)
	if err := _SecurityCouncil.contract.UnpackLog(event, "TransactionSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SecurityCouncilValidationRequestedIterator is returned from FilterValidationRequested and is used to iterate over the raw logs and unpacked data for ValidationRequested events raised by the SecurityCouncil contract.
type SecurityCouncilValidationRequestedIterator struct {
	Event *SecurityCouncilValidationRequested // Event containing the contract specifics and raw log

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
func (it *SecurityCouncilValidationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SecurityCouncilValidationRequested)
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
		it.Event = new(SecurityCouncilValidationRequested)
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
func (it *SecurityCouncilValidationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SecurityCouncilValidationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SecurityCouncilValidationRequested represents a ValidationRequested event raised by the SecurityCouncil contract.
type SecurityCouncilValidationRequested struct {
	TransactionId *big.Int
	OutputRoot    [32]byte
	L2BlockNumber *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidationRequested is a free log retrieval operation binding the contract event 0x00ef5106e82a682c776fd7748be042f406a9ee0feaaea86ae9029477c2b91f2a.
//
// Solidity: event ValidationRequested(uint256 indexed transactionId, bytes32 outputRoot, uint256 l2BlockNumber)
func (_SecurityCouncil *SecurityCouncilFilterer) FilterValidationRequested(opts *bind.FilterOpts, transactionId []*big.Int) (*SecurityCouncilValidationRequestedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SecurityCouncil.contract.FilterLogs(opts, "ValidationRequested", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &SecurityCouncilValidationRequestedIterator{contract: _SecurityCouncil.contract, event: "ValidationRequested", logs: logs, sub: sub}, nil
}

// WatchValidationRequested is a free log subscription operation binding the contract event 0x00ef5106e82a682c776fd7748be042f406a9ee0feaaea86ae9029477c2b91f2a.
//
// Solidity: event ValidationRequested(uint256 indexed transactionId, bytes32 outputRoot, uint256 l2BlockNumber)
func (_SecurityCouncil *SecurityCouncilFilterer) WatchValidationRequested(opts *bind.WatchOpts, sink chan<- *SecurityCouncilValidationRequested, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _SecurityCouncil.contract.WatchLogs(opts, "ValidationRequested", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SecurityCouncilValidationRequested)
				if err := _SecurityCouncil.contract.UnpackLog(event, "ValidationRequested", log); err != nil {
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

// ParseValidationRequested is a log parse operation binding the contract event 0x00ef5106e82a682c776fd7748be042f406a9ee0feaaea86ae9029477c2b91f2a.
//
// Solidity: event ValidationRequested(uint256 indexed transactionId, bytes32 outputRoot, uint256 l2BlockNumber)
func (_SecurityCouncil *SecurityCouncilFilterer) ParseValidationRequested(log types.Log) (*SecurityCouncilValidationRequested, error) {
	event := new(SecurityCouncilValidationRequested)
	if err := _SecurityCouncil.contract.UnpackLog(event, "ValidationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
