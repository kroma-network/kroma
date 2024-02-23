package predeploys

import (
	"github.com/ethereum/go-ethereum/common"
)

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
	L2ERC721Bridge             = "0x420000000000000000000000000000000000000A"
	KromaMintableERC20Factory  = "0x420000000000000000000000000000000000000B"
	KromaMintableERC721Factory = "0x420000000000000000000000000000000000000C"
	MintManager                = "0x4200000000000000000000000000000000000070"
	GovernanceToken            = "0x42000000000000000000000000000000000000FF"
	Create2Deployer            = "0x13b0D85CcB8bf860b6b79AF3029fCA081AE9beF2"
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
	MintManagerAddr                = common.HexToAddress(MintManager)
	Create2DeployerAddr            = common.HexToAddress(Create2Deployer)

	Predeploys          = make(map[string]*Predeploy)
	PredeploysByAddress = make(map[common.Address]*Predeploy)
)

// IsProxied returns true for that will sit behind a proxy contract
func IsProxied(predeployAddr common.Address) bool {
	switch predeployAddr {
	case WETH9Addr:
	case GovernanceTokenAddr:
	default:
		return true
	}
	return false
}

func init() {
	Predeploys["ProxyAdmin"] = &Predeploy{Address: ProxyAdminAddr}
	Predeploys["WETH9"] = &Predeploy{Address: WETH9Addr, ProxyDisabled: true}
	Predeploys["L1Block"] = &Predeploy{Address: L1BlockAddr}
	Predeploys["L2ToL1MessagePasser"] = &Predeploy{Address: L2ToL1MessagePasserAddr}
	Predeploys["L2CrossDomainMessenger"] = &Predeploy{Address: L2CrossDomainMessengerAddr}
	Predeploys["GasPriceOracle"] = &Predeploy{Address: GasPriceOracleAddr}
	Predeploys["ProtocolVault"] = &Predeploy{Address: ProtocolVaultAddr}
	Predeploys["L1FeeVault"] = &Predeploy{Address: L1FeeVaultAddr}
	Predeploys["ValidatorRewardVault"] = &Predeploy{Address: ValidatorRewardVaultAddr}
	Predeploys["L2StandardBridge"] = &Predeploy{Address: L2StandardBridgeAddr}
	Predeploys["GovernanceToken"] = &Predeploy{
		Address: GovernanceTokenAddr,
		Enabled: func(config DeployConfig) bool {
			return config.GovernanceEnabled()
		},
	}
	Predeploys["L2ERC721Bridge"] = &Predeploy{Address: L2ERC721BridgeAddr}
	Predeploys["KromaMintableERC20Factory"] = &Predeploy{Address: KromaMintableERC20FactoryAddr}
	Predeploys["KromaMintableERC721Factory"] = &Predeploy{Address: KromaMintableERC721FactoryAddr}
	Predeploys["MintManager"] = &Predeploy{
		Address: MintManagerAddr,
		Enabled: func(config DeployConfig) bool {
			return config.GovernanceEnabled()
		},
	}
	Predeploys["Create2Deployer"] = &Predeploy{
		Address:       Create2DeployerAddr,
		ProxyDisabled: true,
		Enabled: func(config DeployConfig) bool {
			canyonTime := config.CanyonTime(0)
			return canyonTime != nil && *canyonTime == 0
		},
	}

	for _, predeploy := range Predeploys {
		PredeploysByAddress[predeploy.Address] = predeploy
	}
}
