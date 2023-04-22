package predeploys

import "github.com/ethereum/go-ethereum/common"

const (
	DevProxyAdmin                = "0x6900000000000000000000000000000000000000"
	DevWETH9                     = "0x6900000000000000000000000000000000000001"
	DevSystemConfig              = "0x6900000000000000000000000000000000000002"
	DevKromaPortal               = "0x6900000000000000000000000000000000000003"
	DevL2OutputOracle            = "0x6900000000000000000000000000000000000004"
	DevValidatorPool             = "0x6900000000000000000000000000000000000005"
	DevL1CrossDomainMessenger    = "0x6900000000000000000000000000000000000006"
	DevL1StandardBridge          = "0x6900000000000000000000000000000000000007"
	DevL1ERC721Bridge            = "0x6900000000000000000000000000000000000008"
	DevKromaMintableERC20Factory = "0x6900000000000000000000000000000000000009"
	DevPoseidon2                 = "0x690000000000000000000000000000000000000A"
	DevZKMerkleTrie              = "0x690000000000000000000000000000000000000B"
	DevZKVerifier                = "0x690000000000000000000000000000000000000C"
	DevColosseum                 = "0x690000000000000000000000000000000000000D"
)

var (
	DevProxyAdminAddr                = common.HexToAddress(DevProxyAdmin)
	DevWETH9Addr                     = common.HexToAddress(DevWETH9)
	DevSystemConfigAddr              = common.HexToAddress(DevSystemConfig)
	DevKromaPortalAddr               = common.HexToAddress(DevKromaPortal)
	DevL2OutputOracleAddr            = common.HexToAddress(DevL2OutputOracle)
	DevValidatorPoolAddr             = common.HexToAddress(DevValidatorPool)
	DevL1CrossDomainMessengerAddr    = common.HexToAddress(DevL1CrossDomainMessenger)
	DevL1StandardBridgeAddr          = common.HexToAddress(DevL1StandardBridge)
	DevL1ERC721BridgeAddr            = common.HexToAddress(DevL1ERC721Bridge)
	DevKromaMintableERC20FactoryAddr = common.HexToAddress(DevKromaMintableERC20Factory)
	DevPoseidon2Addr                 = common.HexToAddress(DevPoseidon2)
	DevZKMerkleTrieAddr              = common.HexToAddress(DevZKMerkleTrie)
	DevZKVerifierAddr                = common.HexToAddress(DevZKVerifier)
	DevColosseumAddr                 = common.HexToAddress(DevColosseum)

	DevPredeploys = make(map[string]*common.Address)
)

func init() {
	DevPredeploys["Admin"] = &DevProxyAdminAddr
	DevPredeploys["WETH9"] = &DevWETH9Addr
	DevPredeploys["SystemConfig"] = &DevSystemConfigAddr
	DevPredeploys["KromaPortal"] = &DevKromaPortalAddr
	DevPredeploys["L2OutputOracle"] = &DevL2OutputOracleAddr
	DevPredeploys["ValidatorPool"] = &DevValidatorPoolAddr
	DevPredeploys["L1CrossDomainMessenger"] = &DevL1CrossDomainMessengerAddr
	DevPredeploys["L1StandardBridge"] = &DevL1StandardBridgeAddr
	DevPredeploys["L1ERC721Bridge"] = &DevL1ERC721BridgeAddr
	DevPredeploys["KromaMintableERC20Factory"] = &DevKromaMintableERC20FactoryAddr
	DevPredeploys["Poseidon2"] = &DevPoseidon2Addr
	DevPredeploys["ZKMerkleTrie"] = &DevZKMerkleTrieAddr
	DevPredeploys["ZKVerifier"] = &DevZKVerifierAddr
	DevPredeploys["Colosseum"] = &DevColosseumAddr
}
