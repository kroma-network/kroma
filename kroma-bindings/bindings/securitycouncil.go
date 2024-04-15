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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_colosseum\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_governor\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"COLOSSEUM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOVERNOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractUpgradeGovernor\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"confirmTransaction\",\"inputs\":[{\"name\":\"_transactionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"confirmations\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"confirmationCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeTransaction\",\"inputs\":[{\"name\":\"_transactionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"generateTransactionId\",\"inputs\":[{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConfirmationCount\",\"inputs\":[{\"name\":\"_transactionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVotes\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isConfirmed\",\"inputs\":[{\"name\":\"_transactionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isConfirmedBy\",\"inputs\":[{\"name\":\"_transactionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"outputsDeleteRequested\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quorum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestDeletion\",\"inputs\":[{\"name\":\"_outputIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_force\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestValidation\",\"inputs\":[{\"name\":\"_outputRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_l2BlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeConfirmation\",\"inputs\":[{\"name\":\"_transactionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitTransaction\",\"inputs\":[{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transactionCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transactions\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ConfirmationRevoked\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"transactionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DeletionRequested\",\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"outputIndex\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransactionConfirmed\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"transactionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransactionExecuted\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"transactionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransactionSubmitted\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"transactionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidationRequested\",\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"outputRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"l2BlockNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60c06040523480156200001157600080fd5b50604051620021b5380380620021b5833981016040819052620000349162000065565b6001600160a01b039081166080521660a052620000a4565b6001600160a01b03811681146200006257600080fd5b50565b600080604083850312156200007957600080fd5b825162000086816200004c565b602084015190925062000099816200004c565b809150509250929050565b60805160a0516120ba620000fb60003960008181610379015281816105a2015261065301526000818161025b0152818161075a015281816108160152818161089801528181610dad0152610ed901526120ba6000f3fe608060405234801561001057600080fd5b506004361061016c5760003560e01c80638b51d13f116100cd578063b77bf60011610081578063c01a8c8411610066578063c01a8c84146103c4578063c6427474146103d7578063ee22610b146103ea57600080fd5b8063b77bf6001461039b578063b9774f7b146103a457600080fd5b80639ab24eb0116100b25780639ab24eb01461033e5780639ace38c2146103515780639e45e8f41461037457600080fd5b80638b51d13f146102ff57806391ddadf41461031f57600080fd5b806349ae963d116101245780636dc0ae22116101095780636dc0ae2214610256578063784547a7146102a25780638a8e784c146102b557600080fd5b806349ae963d146101fa57806354fd4d501461020d57600080fd5b80631703a018116101555780631703a0181461019957806320ea8d86146101b45780632a758595146101c757600080fd5b80630192337114610171578063080b91ee14610186575b600080fd5b61018461017f366004611a5d565b6103fd565b005b610184610194366004611b6c565b61063b565b6101a1610737565b6040519081526020015b60405180910390f35b6101846101c2366004611bbc565b6109ba565b6101ea6101d5366004611bbc565b60396020526000908152604090205460ff1681565b60405190151581526020016101ab565b6101a1610208366004611bf7565b610c9c565b6102496040518060400160405280600581526020017f312e312e3000000000000000000000000000000000000000000000000000000081525081565b6040516101ab9190611ca5565b61027d7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101ab565b6101ea6102b0366004611bbc565b610d85565b6101ea6102c3366004611cbf565b600082815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845260010190915290205460ff1692915050565b6101a161030d366004611bbc565b60009081526034602052604090205490565b610327610da9565b60405165ffffffffffff90911681526020016101ab565b6101a161034c366004611ce4565b610ed5565b61036461035f366004611bbc565b610ffe565b6040516101ab9493929190611d01565b61027d7f000000000000000000000000000000000000000000000000000000000000000081565b6101a160385481565b6101a16103b2366004611bbc565b60346020526000908152604090205481565b6101846103d2366004611bbc565b6110dc565b6101a16103e5366004611bf7565b61133e565b6101846103f8366004611bbc565b6113d3565b33600061040982610ed5565b116104815760405162461bcd60e51b815260206004820152603b60248201527f546f6b656e4d756c746953696757616c6c65743a206f6e6c7920616c6c6f776560448201527f6420746f20676f7665726e616e636520746f6b656e206f776e6572000000000060648201526084015b60405180910390fd5b60008381526039602052604090205460ff16158061049c5750815b6105355760405162461bcd60e51b8152602060048201526044602482018190527f5365637572697479436f756e63696c3a20746865206f75747075742068617320908201527f616c7265616479206265656e2072657175657374656420746f2062652064656c60648201527f6574656400000000000000000000000000000000000000000000000000000000608482015260a401610478565b6040805160248082018690528251808303909101815260449091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe39a219c0000000000000000000000000000000000000000000000000000000017905260006105c87f0000000000000000000000000000000000000000000000000000000000000000828461133e565b90506105d3816110dc565b60008581526039602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905551869183917fc63c84660a471a970585c7cab9d0601af8e717ff0822a2ea049a3542fc5aa55a9190a35050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146106e65760405162461bcd60e51b815260206004820152603c60248201527f5365637572697479436f756e63696c3a206f6e6c792074686520636f6c6f737360448201527f65756d20636f6e74726163742063616e20626520612073656e646572000000006064820152608401610478565b60006106f433600084611741565b604080518681526020810186905291925082917eef5106e82a682c776fd7748be042f406a9ee0feaaea86ae9029477c2b91f2a910160405180910390a250505050565b6000806001610744610da9565b61074e9190611d77565b65ffffffffffff1690507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166397c3d3346040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107e79190611d9e565b6040517f60c4247f000000000000000000000000000000000000000000000000000000008152600481018390527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906360c4247f90602401602060405180830381865afa158015610872573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108969190611d9e565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663fc0c546a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610901573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109259190611db7565b73ffffffffffffffffffffffffffffffffffffffff16638e539e8c846040518263ffffffff1660e01b815260040161095f91815260200190565b602060405180830381865afa15801561097c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109a09190611d9e565b6109aa9190611dd4565b6109b49190611e11565b91505090565b3360006109c682610ed5565b11610a395760405162461bcd60e51b815260206004820152603b60248201527f546f6b656e4d756c746953696757616c6c65743a206f6e6c7920616c6c6f776560448201527f6420746f20676f7665726e616e636520746f6b656e206f776e657200000000006064820152608401610478565b600082815260336020526040902054829073ffffffffffffffffffffffffffffffffffffffff16610ad25760405162461bcd60e51b815260206004820152602f60248201527f546f6b656e4d756c746953696757616c6c65743a207472616e73616374696f6e60448201527f20646f6573206e6f7420657869737400000000000000000000000000000000006064820152608401610478565b600083815260336020526040902054839074010000000000000000000000000000000000000000900460ff1615610b715760405162461bcd60e51b815260206004820152602560248201527f546f6b656e4d756c746953696757616c6c65743a20616c72656164792065786560448201527f63757465640000000000000000000000000000000000000000000000000000006064820152608401610478565b600084815260346020908152604080832033845260010190915290205460ff16610c035760405162461bcd60e51b815260206004820152602660248201527f546f6b656e4d756c746953696757616c6c65743a206e6f7420636f6e6669726d60448201527f65642079657400000000000000000000000000000000000000000000000000006064820152608401610478565b60008481526034602090815260408083203380855260018201909352922080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055610c5090610ed5565b816000016000828254610c639190611e4c565b9091555050604051859033907f795394da21278ca39d59bb3ca00efeebdc0679acc420916c7385c2c5d942656f90600090a35050505050565b60008373ffffffffffffffffffffffffffffffffffffffff8116610d285760405162461bcd60e51b815260206004820152602960248201527f546f6b656e4d756c746953696757616c6c65743a20616464726573732069732060448201527f6e6f742076616c696400000000000000000000000000000000000000000000006064820152608401610478565b848484610d33610da9565b604051602001610d469493929190611e63565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012095945050505050565b6000610d8f610737565b600092835260346020526040909220549190911015919050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663fc0c546a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e16573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e3a9190611db7565b73ffffffffffffffffffffffffffffffffffffffff166391ddadf46040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610ebe575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252610ebb91810190611eb1565b60015b610ed057610ecb43611968565b905090565b919050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663fc0c546a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f42573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f669190611db7565b6040517f9ab24eb000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301529190911690639ab24eb090602401602060405180830381865afa158015610fd4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ff89190611d9e565b92915050565b60336020526000908152604090208054600182015460028301805473ffffffffffffffffffffffffffffffffffffffff8416947401000000000000000000000000000000000000000090940460ff1693919061105990611ed9565b80601f016020809104026020016040519081016040528092919081815260200182805461108590611ed9565b80156110d25780601f106110a7576101008083540402835291602001916110d2565b820191906000526020600020905b8154815290600101906020018083116110b557829003601f168201915b5050505050905084565b3360006110e882610ed5565b1161115b5760405162461bcd60e51b815260206004820152603b60248201527f546f6b656e4d756c746953696757616c6c65743a206f6e6c7920616c6c6f776560448201527f6420746f20676f7665726e616e636520746f6b656e206f776e657200000000006064820152608401610478565b600082815260336020526040902054829073ffffffffffffffffffffffffffffffffffffffff166111f45760405162461bcd60e51b815260206004820152602f60248201527f546f6b656e4d756c746953696757616c6c65743a207472616e73616374696f6e60448201527f20646f6573206e6f7420657869737400000000000000000000000000000000006064820152608401610478565b6000838152603460209081526040808320338452600181019092529091205460ff16156112895760405162461bcd60e51b815260206004820152602660248201527f546f6b656e4d756c746953696757616c6c65743a20616c726561647920636f6e60448201527f6669726d656400000000000000000000000000000000000000000000000000006064820152608401610478565b3360008181526001838101602052604090912080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690911790556112ce90610ed5565b8160000160008282546112e19190611f2c565b9091555050604051849033907ff8a17c9136a3ae33364fac05eb088a3cbafee10c1889c88593e20ee2d8e4eb8890600090a361131b610737565b6000858152603460205260409020541061133857611338846113d3565b50505050565b600033600061134c82610ed5565b116113bf5760405162461bcd60e51b815260206004820152603b60248201527f546f6b656e4d756c746953696757616c6c65743a206f6e6c7920616c6c6f776560448201527f6420746f20676f7665726e616e636520746f6b656e206f776e657200000000006064820152608401610478565b6113ca858585611741565b95945050505050565b6113db6119ea565b600081815260336020526040902054819073ffffffffffffffffffffffffffffffffffffffff166114745760405162461bcd60e51b815260206004820152602f60248201527f546f6b656e4d756c746953696757616c6c65743a207472616e73616374696f6e60448201527f20646f6573206e6f7420657869737400000000000000000000000000000000006064820152608401610478565b600082815260336020526040902054829074010000000000000000000000000000000000000000900460ff16156115135760405162461bcd60e51b815260206004820152602560248201527f546f6b656e4d756c746953696757616c6c65743a20616c72656164792065786560448201527f63757465640000000000000000000000000000000000000000000000000000006064820152608401610478565b61151c83610d85565b61158e5760405162461bcd60e51b815260206004820152602760248201527f546f6b656e4d756c746953696757616c6c65743a2071756f72756d206e6f742060448201527f72656163686564000000000000000000000000000000000000000000000000006064820152608401610478565b600083815260336020526040812080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff81167401000000000000000000000000000000000000000017825590919061168f9073ffffffffffffffffffffffffffffffffffffffff165a846001015485600201805461160c90611ed9565b80601f016020809104026020016040519081016040528092919081815260200182805461163890611ed9565b80156116855780601f1061165a57610100808354040283529160200191611685565b820191906000526020600020905b81548152906001019060200180831161166857829003601f168201915b5050505050611a43565b9050806117045760405162461bcd60e51b815260206004820152602c60248201527f546f6b656e4d756c746953696757616c6c65743a2063616c6c207472616e736160448201527f6374696f6e206661696c656400000000000000000000000000000000000000006064820152608401610478565b604051859033907f4e86ad0da28cbaaaa7e93e36c43b32696e970535225b316f1b84fbf30bdc04e890600090a35050505061173e60018055565b50565b60008373ffffffffffffffffffffffffffffffffffffffff81166117cd5760405162461bcd60e51b815260206004820152602960248201527f546f6b656e4d756c746953696757616c6c65743a20616464726573732069732060448201527f6e6f742076616c696400000000000000000000000000000000000000000000006064820152608401610478565b60006117da868686610c9c565b60008181526033602052604090205490915073ffffffffffffffffffffffffffffffffffffffff16156118755760405162461bcd60e51b815260206004820152602f60248201527f546f6b656e4d756c746953696757616c6c65743a207472616e73616374696f6e60448201527f20616c72656164792065786973747300000000000000000000000000000000006064820152608401610478565b6040805160808101825273ffffffffffffffffffffffffffffffffffffffff8089168252600060208084018281528486018b8152606086018b8152888552603390935295909220845181549351151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00000000000000000000000000000000000000000090941694169390931791909117825592516001820155915190919060028201906119269082611f93565b505060388054600101905550604051819033907f1f50cd00b6a6fe3928bf4a5f2f23829e9a1c9396573b828b5fa14d95aae7e77590600090a395945050505050565b600065ffffffffffff8211156119e65760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203460448201527f38206269747300000000000000000000000000000000000000000000000000006064820152608401610478565b5090565b600260015403611a3c5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610478565b6002600155565b600080600080845160208601878a8af19695505050505050565b60008060408385031215611a7057600080fd5b8235915060208301358015158114611a8757600080fd5b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112611ad257600080fd5b813567ffffffffffffffff80821115611aed57611aed611a92565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715611b3357611b33611a92565b81604052838152866020858801011115611b4c57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600060608486031215611b8157600080fd5b8335925060208401359150604084013567ffffffffffffffff811115611ba657600080fd5b611bb286828701611ac1565b9150509250925092565b600060208284031215611bce57600080fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff8116811461173e57600080fd5b600080600060608486031215611c0c57600080fd5b8335611c1781611bd5565b925060208401359150604084013567ffffffffffffffff811115611ba657600080fd5b6000815180845260005b81811015611c6057602081850181015186830182015201611c44565b81811115611c72576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611cb86020830184611c3a565b9392505050565b60008060408385031215611cd257600080fd5b823591506020830135611a8781611bd5565b600060208284031215611cf657600080fd5b8135611cb881611bd5565b73ffffffffffffffffffffffffffffffffffffffff851681528315156020820152826040820152608060608201526000611d3e6080830184611c3a565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600065ffffffffffff83811690831681811015611d9657611d96611d48565b039392505050565b600060208284031215611db057600080fd5b5051919050565b600060208284031215611dc957600080fd5b8151611cb881611bd5565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611e0c57611e0c611d48565b500290565b600082611e47577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b600082821015611e5e57611e5e611d48565b500390565b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152608060408201526000611e986080830185611c3a565b905065ffffffffffff8316606083015295945050505050565b600060208284031215611ec357600080fd5b815165ffffffffffff81168114611cb857600080fd5b600181811c90821680611eed57607f821691505b602082108103611f26577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b60008219821115611f3f57611f3f611d48565b500190565b601f821115611f8e57600081815260208120601f850160051c81016020861015611f6b5750805b601f850160051c820191505b81811015611f8a57828155600101611f77565b5050505b505050565b815167ffffffffffffffff811115611fad57611fad611a92565b611fc181611fbb8454611ed9565b84611f44565b602080601f8311600181146120145760008415611fde5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611f8a565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b8281101561206157888601518255948401946001909101908401612042565b508582101561209d57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b0190555056fea164736f6c634300080f000a",
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
