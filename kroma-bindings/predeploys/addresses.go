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
	L1FeeVault                 = "0x4200000000000000000000000000000000000007"
	ValidatorRewardVault       = "0x4200000000000000000000000000000000000008"
	L2StandardBridge           = "0x4200000000000000000000000000000000000009"
	GovernanceToken            = "0x4200000000000000000000000000000000000010"
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
	L1FeeVaultAddr                 = common.HexToAddress(L1FeeVault)
	ValidatorRewardVaultAddr       = common.HexToAddress(ValidatorRewardVault)
	L2StandardBridgeAddr           = common.HexToAddress(L2StandardBridge)
	GovernanceTokenAddr            = common.HexToAddress(GovernanceToken)
	L2ERC721BridgeAddr             = common.HexToAddress(L2ERC721Bridge)
	KromaMintableERC20FactoryAddr  = common.HexToAddress(KromaMintableERC20Factory)
	KromaMintableERC721FactoryAddr = common.HexToAddress(KromaMintableERC721Factory)

	Predeploys = make(map[string]*common.Address)
)

// IsProxied returns true for predeploys that will sit behind a proxy contract
func IsProxied(predeployAddr common.Address) bool {
	switch predeployAddr {
	case WETH9Addr:
	default:
		return true
	}
	return false
}

func init() {
	Predeploys["ProxyAdmin"] = &ProxyAdminAddr
	Predeploys["WETH9"] = &WETH9Addr
	Predeploys["L1Block"] = &L1BlockAddr
	Predeploys["L2ToL1MessagePasser"] = &L2ToL1MessagePasserAddr
	Predeploys["L2CrossDomainMessenger"] = &L2CrossDomainMessengerAddr
	Predeploys["GasPriceOracle"] = &GasPriceOracleAddr
	Predeploys["ProtocolVault"] = &ProtocolVaultAddr
	Predeploys["L1FeeVault"] = &L1FeeVaultAddr
	Predeploys["ValidatorRewardVault"] = &ValidatorRewardVaultAddr
	Predeploys["L2StandardBridge"] = &L2StandardBridgeAddr
	Predeploys["GovernanceToken"] = &GovernanceTokenAddr
	Predeploys["L2ERC721Bridge"] = &L2ERC721BridgeAddr
	Predeploys["KromaMintableERC20Factory"] = &KromaMintableERC20FactoryAddr
	Predeploys["KromaMintableERC721Factory"] = &KromaMintableERC721FactoryAddr
}
