// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const GasPriceOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"isEcotone\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_bool\"}],\"types\":{\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"}}}"

var GasPriceOracleStorageLayout = new(solc.StorageLayout)

var GasPriceOracleDeployedBin = "0x608060405234801561001057600080fd5b506004361061011b5760003560e01c806354fd4d50116100b2578063c598591811610081578063f45e65d811610066578063f45e65d81461023e578063f820614014610246578063fe173b971461021557600080fd5b8063c598591814610223578063de26c4a11461022b57600080fd5b806354fd4d50146101af57806368d5dca6146101f85780636ef25c3a146102155780638cca67621461021b57600080fd5b80634621e226116100ee5780634621e2261461015457806349948e0e146101875780634ef6e2241461019a578063519b4bd3146101a757600080fd5b80630c18c1621461012057806322b90ab31461013b5780632e0f262514610145578063313ce5671461014d575b600080fd5b61012861024e565b6040519081526020015b60405180910390f35b61014361036f565b005b610128600681565b6006610128565b7f8f72cb8c9ce0db6b33874dcafb4caaa75b2406f18992b3856856f86d8949de5c545b6040519015158152602001610132565b61012861019536600461102c565b610592565b6000546101779060ff1681565b6101286105b6565b6101eb6040518060400160405280600581526020017f312e312e3000000000000000000000000000000000000000000000000000000081525081565b60405161013291906110fb565b6102006106a3565b60405163ffffffff9091168152602001610132565b48610128565b6101436107b4565b6102006109ec565b61012861023936600461102c565b610ad9565b610128610b8d565b610128610c80565b6000805460ff16156102e7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f47617350726963654f7261636c653a206f76657268656164282920697320646560448201527f707265636174656400000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b73420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff16638b239f736040518163ffffffff1660e01b8152600401602060405180830381865afa158015610346573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036a919061116e565b905090565b73420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff1663e591b2826040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103f29190611187565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104d2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f47617350726963654f7261636c653a206f6e6c7920746865206465706f73697460448201527f6f72206163636f756e742063616e2073657420697345636f746f6e6520666c6160648201527f6700000000000000000000000000000000000000000000000000000000000000608482015260a4016102de565b60005460ff1615610565576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f47617350726963654f7261636c653a2045636f746f6e6520616c72656164792060448201527f616374697665000000000000000000000000000000000000000000000000000060648201526084016102de565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b6000805460ff16156105ad576105a782610d6d565b92915050565b6105a782610e11565b60006105e07f8f72cb8c9ce0db6b33874dcafb4caaa75b2406f18992b3856856f86d8949de5c5490565b156106445773420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16635cf249696040518163ffffffff1660e01b8152600401602060405180830381865afa158015610346573d6000803e3d6000fd5b73420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff16635cf249696040518163ffffffff1660e01b8152600401602060405180830381865afa158015610346573d6000803e3d6000fd5b60006106cd7f8f72cb8c9ce0db6b33874dcafb4caaa75b2406f18992b3856856f86d8949de5c5490565b156107555773420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff166368d5dca66040518163ffffffff1660e01b8152600401602060405180830381865afa158015610731573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036a91906111bd565b73420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff166368d5dca66040518163ffffffff1660e01b8152600401602060405180830381865afa158015610731573d6000803e3d6000fd5b73420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff1663e591b2826040518163ffffffff1660e01b8152600401602060405180830381865afa158015610813573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108379190611187565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610917576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604260248201527f47617350726963654f7261636c653a206f6e6c7920746865206465706f73697460448201527f6f72206163636f756e742063616e207365742069734b726f6d614d505420666c60648201527f6167000000000000000000000000000000000000000000000000000000000000608482015260a4016102de565b7f8f72cb8c9ce0db6b33874dcafb4caaa75b2406f18992b3856856f86d8949de5c54156109c6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f47617350726963654f7261636c653a204b726f6d61204d505420616c7265616460448201527f792061637469766500000000000000000000000000000000000000000000000060648201526084016102de565b60017f8f72cb8c9ce0db6b33874dcafb4caaa75b2406f18992b3856856f86d8949de5c55565b6000610a167f8f72cb8c9ce0db6b33874dcafb4caaa75b2406f18992b3856856f86d8949de5c5490565b15610a7a5773420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff1663c59859186040518163ffffffff1660e01b8152600401602060405180830381865afa158015610731573d6000803e3d6000fd5b73420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff1663c59859186040518163ffffffff1660e01b8152600401602060405180830381865afa158015610731573d6000803e3d6000fd5b600080610ae583610f6d565b60005490915060ff1615610af95792915050565b73420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff16638b239f736040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b58573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b7c919061116e565b610b869082611212565b9392505050565b6000805460ff1615610c21576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f47617350726963654f7261636c653a207363616c61722829206973206465707260448201527f656361746564000000000000000000000000000000000000000000000000000060648201526084016102de565b73420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff16639e8c49666040518163ffffffff1660e01b8152600401602060405180830381865afa158015610346573d6000803e3d6000fd5b6000610caa7f8f72cb8c9ce0db6b33874dcafb4caaa75b2406f18992b3856856f86d8949de5c5490565b15610d0e5773420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff1663f82061406040518163ffffffff1660e01b8152600401602060405180830381865afa158015610346573d6000803e3d6000fd5b73420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff1663f82061406040518163ffffffff1660e01b8152600401602060405180830381865afa158015610346573d6000803e3d6000fd5b600080610d7983610f6d565b90506000610d856105b6565b610d8d6109ec565b610d9890601061122a565b63ffffffff16610da89190611256565b90506000610db4610c80565b610dbc6106a3565b63ffffffff16610dcc9190611256565b90506000610dda8284611212565b610de49085611256565b9050610df26006600a6113b3565b610dfd906010611256565b610e0790826113bf565b9695505050505050565b600080610e1d83610f6d565b9050600073420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff16639e8c49666040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e80573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ea4919061116e565b610eac6105b6565b73420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff16638b239f736040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f0b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f2f919061116e565b610f399085611212565b610f439190611256565b610f4d9190611256565b9050610f5b6006600a6113b3565b610f6590826113bf565b949350505050565b80516000908190815b81811015610ff057848181518110610f9057610f906113fa565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016600003610fd057610fc9600484611212565b9250610fde565b610fdb601084611212565b92505b80610fe881611429565b915050610f76565b50610f6582610440611212565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006020828403121561103e57600080fd5b813567ffffffffffffffff8082111561105657600080fd5b818401915084601f83011261106a57600080fd5b81358181111561107c5761107c610ffd565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156110c2576110c2610ffd565b816040528281528760208487010111156110db57600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208083528351808285015260005b818110156111285785810183015185820160400152820161110c565b8181111561113a576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60006020828403121561118057600080fd5b5051919050565b60006020828403121561119957600080fd5b815173ffffffffffffffffffffffffffffffffffffffff81168114610b8657600080fd5b6000602082840312156111cf57600080fd5b815163ffffffff81168114610b8657600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115611225576112256111e3565b500190565b600063ffffffff8083168185168183048111821515161561124d5761124d6111e3565b02949350505050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561128e5761128e6111e3565b500290565b600181815b808511156112ec57817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156112d2576112d26111e3565b808516156112df57918102915b93841c9390800290611298565b509250929050565b600082611303575060016105a7565b81611310575060006105a7565b816001811461132657600281146113305761134c565b60019150506105a7565b60ff841115611341576113416111e3565b50506001821b6105a7565b5060208310610133831016604e8410600b841016171561136f575081810a6105a7565b6113798383611293565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156113ab576113ab6111e3565b029392505050565b6000610b8683836112f4565b6000826113f5577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361145a5761145a6111e3565b506001019056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(GasPriceOracleStorageLayoutJSON), GasPriceOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["GasPriceOracle"] = GasPriceOracleStorageLayout
	deployedBytecodes["GasPriceOracle"] = GasPriceOracleDeployedBin
	immutableReferences["GasPriceOracle"] = false
}
