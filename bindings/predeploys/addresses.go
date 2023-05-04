package predeploys

import "github.com/ethereum/go-ethereum/common"

const (
	ProxyAdmin                 = "0x4200000000000000000000000000000000000000"
	WETH9                      = "0x4200000000000000000000000000000000000001"
	L1Block                    = "0x4200000000000000000000000000000000000002"
	L2ToL1MessagePasser        = "0x4200000000000000000000000000000000000003"
	L2CrossDomainMessenger     = "0x4200000000000000000000000000000000000004"
	GasPriceOracle             = "0x4200000000000000000000000000000000000005"
	ProtocolVault              = "0x4200000000000000000000000000000000000006"
	ProposerRewardVault        = "0x4200000000000000000000000000000000000007"
	ValidatorRewardVault       = "0x4200000000000000000000000000000000000008"
	L2StandardBridge           = "0x4200000000000000000000000000000000000009"
	L2ERC721Bridge             = "0x420000000000000000000000000000000000000A"
	KromaMintableERC20Factory  = "0x420000000000000000000000000000000000000B"
	KromaMintableERC721Factory = "0x420000000000000000000000000000000000000C"
)

var (
	ProxyAdminAddr                 = common.HexToAddress(ProxyAdmin)
	WETH9Addr                      = common.HexToAddress(WETH9)
	L1BlockAddr                    = common.HexToAddress(L1Block)
	L2ToL1MessagePasserAddr        = common.HexToAddress(L2ToL1MessagePasser)
	L2CrossDomainMessengerAddr     = common.HexToAddress(L2CrossDomainMessenger)
	GasPriceOracleAddr             = common.HexToAddress(GasPriceOracle)
	ProtocolVaultAddr              = common.HexToAddress(ProtocolVault)
	ProposerRewardVaultAddr        = common.HexToAddress(ProposerRewardVault)
	ValidatorRewardVaultAddr       = common.HexToAddress(ValidatorRewardVault)
	L2StandardBridgeAddr           = common.HexToAddress(L2StandardBridge)
	L2ERC721BridgeAddr             = common.HexToAddress(L2ERC721Bridge)
	KromaMintableERC20FactoryAddr  = common.HexToAddress(KromaMintableERC20Factory)
	KromaMintableERC721FactoryAddr = common.HexToAddress(KromaMintableERC721Factory)

	Predeploys = make(map[string]*common.Address)
)

func init() {
	Predeploys["ProxyAdmin"] = &ProxyAdminAddr
	Predeploys["WETH9"] = &WETH9Addr
	Predeploys["L1Block"] = &L1BlockAddr
	Predeploys["L2ToL1MessagePasser"] = &L2ToL1MessagePasserAddr
	Predeploys["L2CrossDomainMessenger"] = &L2CrossDomainMessengerAddr
	Predeploys["GasPriceOracle"] = &GasPriceOracleAddr
	Predeploys["ProtocolVault"] = &ProtocolVaultAddr
	Predeploys["ProposerRewardVault"] = &ProposerRewardVaultAddr
	Predeploys["ValidatorRewardVault"] = &ValidatorRewardVaultAddr
	Predeploys["L2StandardBridge"] = &L2StandardBridgeAddr
	Predeploys["L2ERC721Bridge"] = &L2ERC721BridgeAddr
	Predeploys["KromaMintableERC20Factory"] = &KromaMintableERC20FactoryAddr
	Predeploys["KromaMintableERC721Factory"] = &KromaMintableERC721FactoryAddr
}
