// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const GasPriceOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/GasPriceOracle.sol:GasPriceOracle\",\"label\":\"isEcotone\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_bool\"}],\"types\":{\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"}}}"

var GasPriceOracleStorageLayout = new(solc.StorageLayout)

var GasPriceOracleDeployedBin = "0x608060405234801561001057600080fd5b506004361061011b5760003560e01c806368d5dca6116100b2578063c598591811610081578063f45e65d811610066578063f45e65d814610242578063f82061401461024a578063fe173b97146101f257600080fd5b8063c598591814610227578063de26c4a11461022f57600080fd5b806368d5dca6146101d55780636ef25c3a146101f25780638cca6762146101f8578063a566e1a51461020057600080fd5b806349948e0e116100ee57806349948e0e146101545780634ef6e22414610167578063519b4bd31461018457806354fd4d501461018c57600080fd5b80630c18c1621461012057806322b90ab31461013b5780632e0f262514610145578063313ce5671461014d575b600080fd5b610128610252565b6040519081526020015b60405180910390f35b610143610373565b005b610128600681565b6006610128565b6101286101623660046110c7565b610596565b6000546101749060ff1681565b6040519015158152602001610132565b6101286105ba565b6101c86040518060400160405280600581526020017f312e312e3000000000000000000000000000000000000000000000000000000081525081565b6040516101329190611196565b6101dd6106a7565b60405163ffffffff9091168152602001610132565b48610128565b6101436107b8565b7f8f72cb8c9ce0db6b33874dcafb4caaa75b2406f18992b3856856f86d8949de5c54610174565b6101dd610a87565b61012861023d3660046110c7565b610b74565b610128610c28565b610128610d1b565b6000805460ff16156102eb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f47617350726963654f7261636c653a206f76657268656164282920697320646560448201527f707265636174656400000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b73420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff16638b239f736040518163ffffffff1660e01b8152600401602060405180830381865afa15801561034a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036e9190611209565b905090565b73420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff1663e591b2826040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103d2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103f69190611222565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104d6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f47617350726963654f7261636c653a206f6e6c7920746865206465706f73697460448201527f6f72206163636f756e742063616e2073657420697345636f746f6e6520666c6160648201527f6700000000000000000000000000000000000000000000000000000000000000608482015260a4016102e2565b60005460ff1615610569576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f47617350726963654f7261636c653a2045636f746f6e6520616c72656164792060448201527f616374697665000000000000000000000000000000000000000000000000000060648201526084016102e2565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b6000805460ff16156105b1576105ab82610e08565b92915050565b6105ab82610eac565b60006105e47f8f72cb8c9ce0db6b33874dcafb4caaa75b2406f18992b3856856f86d8949de5c5490565b156106485773420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16635cf249696040518163ffffffff1660e01b8152600401602060405180830381865afa15801561034a573d6000803e3d6000fd5b73420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff16635cf249696040518163ffffffff1660e01b8152600401602060405180830381865afa15801561034a573d6000803e3d6000fd5b60006106d17f8f72cb8c9ce0db6b33874dcafb4caaa75b2406f18992b3856856f86d8949de5c5490565b156107595773420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff166368d5dca66040518163ffffffff1660e01b8152600401602060405180830381865afa158015610735573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036e9190611258565b73420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff166368d5dca66040518163ffffffff1660e01b8152600401602060405180830381865afa158015610735573d6000803e3d6000fd5b73420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff1663e591b2826040518163ffffffff1660e01b8152600401602060405180830381865afa158015610817573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061083b9190611222565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461091b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604260248201527f47617350726963654f7261636c653a206f6e6c7920746865206465706f73697460448201527f6f72206163636f756e742063616e207365742069734b726f6d614d505420666c60648201527f6167000000000000000000000000000000000000000000000000000000000000608482015260a4016102e2565b60005460ff1615156001146109b2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f47617350726963654f7261636c653a2045636f746f6e65206973206e6f74206160448201527f637469766500000000000000000000000000000000000000000000000000000060648201526084016102e2565b7f8f72cb8c9ce0db6b33874dcafb4caaa75b2406f18992b3856856f86d8949de5c5415610a61576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f47617350726963654f7261636c653a204b726f6d61204d505420616c7265616460448201527f792061637469766500000000000000000000000000000000000000000000000060648201526084016102e2565b60017f8f72cb8c9ce0db6b33874dcafb4caaa75b2406f18992b3856856f86d8949de5c55565b6000610ab17f8f72cb8c9ce0db6b33874dcafb4caaa75b2406f18992b3856856f86d8949de5c5490565b15610b155773420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff1663c59859186040518163ffffffff1660e01b8152600401602060405180830381865afa158015610735573d6000803e3d6000fd5b73420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff1663c59859186040518163ffffffff1660e01b8152600401602060405180830381865afa158015610735573d6000803e3d6000fd5b600080610b8083611008565b60005490915060ff1615610b945792915050565b73420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff16638b239f736040518163ffffffff1660e01b8152600401602060405180830381865afa158015610bf3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c179190611209565b610c2190826112ad565b9392505050565b6000805460ff1615610cbc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f47617350726963654f7261636c653a207363616c61722829206973206465707260448201527f656361746564000000000000000000000000000000000000000000000000000060648201526084016102e2565b73420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff16639e8c49666040518163ffffffff1660e01b8152600401602060405180830381865afa15801561034a573d6000803e3d6000fd5b6000610d457f8f72cb8c9ce0db6b33874dcafb4caaa75b2406f18992b3856856f86d8949de5c5490565b15610da95773420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff1663f82061406040518163ffffffff1660e01b8152600401602060405180830381865afa15801561034a573d6000803e3d6000fd5b73420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff1663f82061406040518163ffffffff1660e01b8152600401602060405180830381865afa15801561034a573d6000803e3d6000fd5b600080610e1483611008565b90506000610e206105ba565b610e28610a87565b610e339060106112c5565b63ffffffff16610e4391906112f1565b90506000610e4f610d1b565b610e576106a7565b63ffffffff16610e6791906112f1565b90506000610e7582846112ad565b610e7f90856112f1565b9050610e8d6006600a61144e565b610e989060106112f1565b610ea2908261145a565b9695505050505050565b600080610eb883611008565b9050600073420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff16639e8c49666040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f1b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f3f9190611209565b610f476105ba565b73420000000000000000000000000000000000000273ffffffffffffffffffffffffffffffffffffffff16638b239f736040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fa6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fca9190611209565b610fd490856112ad565b610fde91906112f1565b610fe891906112f1565b9050610ff66006600a61144e565b611000908261145a565b949350505050565b80516000908190815b8181101561108b5784818151811061102b5761102b611495565b01602001517fff000000000000000000000000000000000000000000000000000000000000001660000361106b576110646004846112ad565b9250611079565b6110766010846112ad565b92505b80611083816114c4565b915050611011565b50611000826104406112ad565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000602082840312156110d957600080fd5b813567ffffffffffffffff808211156110f157600080fd5b818401915084601f83011261110557600080fd5b81358181111561111757611117611098565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561115d5761115d611098565b8160405282815287602084870101111561117657600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208083528351808285015260005b818110156111c3578581018301518582016040015282016111a7565b818111156111d5576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60006020828403121561121b57600080fd5b5051919050565b60006020828403121561123457600080fd5b815173ffffffffffffffffffffffffffffffffffffffff81168114610c2157600080fd5b60006020828403121561126a57600080fd5b815163ffffffff81168114610c2157600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156112c0576112c061127e565b500190565b600063ffffffff808316818516818304811182151516156112e8576112e861127e565b02949350505050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156113295761132961127e565b500290565b600181815b8085111561138757817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0482111561136d5761136d61127e565b8085161561137a57918102915b93841c9390800290611333565b509250929050565b60008261139e575060016105ab565b816113ab575060006105ab565b81600181146113c157600281146113cb576113e7565b60019150506105ab565b60ff8411156113dc576113dc61127e565b50506001821b6105ab565b5060208310610133831016604e8410600b841016171561140a575081810a6105ab565b611414838361132e565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156114465761144661127e565b029392505050565b6000610c21838361138f565b600082611490577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036114f5576114f561127e565b506001019056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(GasPriceOracleStorageLayoutJSON), GasPriceOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["GasPriceOracle"] = GasPriceOracleStorageLayout
	deployedBytecodes["GasPriceOracle"] = GasPriceOracleDeployedBin
	immutableReferences["GasPriceOracle"] = false
}
