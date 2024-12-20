package predeploys

import (
	"github.com/ethereum/go-ethereum/common"

	oppredeploys "github.com/ethereum-optimism/optimism/op-bindings/predeploys"
)

const (
	ProxyAdmin                   = "0x4200000000000000000000000000000000000000"
	WETH9                        = "0x4200000000000000000000000000000000000001"
	KromaL1Block                 = "0x4200000000000000000000000000000000000002"
	L2ToL1MessagePasser          = "0x4200000000000000000000000000000000000003"
	L2CrossDomainMessenger       = "0x4200000000000000000000000000000000000004"
	GasPriceOracle               = "0x4200000000000000000000000000000000000005"
	ProtocolVault                = "0x4200000000000000000000000000000000000006"
	L1FeeVault                   = "0x4200000000000000000000000000000000000007"
	ValidatorRewardVault         = "0x4200000000000000000000000000000000000008"
	L2StandardBridge             = "0x4200000000000000000000000000000000000009"
	L2ERC721Bridge               = "0x420000000000000000000000000000000000000A"
	KromaMintableERC20Factory    = "0x420000000000000000000000000000000000000B"
	KromaMintableERC721Factory   = "0x420000000000000000000000000000000000000C"
	Create2Deployer              = "0x13b0D85CcB8bf860b6b79AF3029fCA081AE9beF2"
	MultiCall3                   = "0xcA11bde05977b3631167028862bE2a173976CA11"
	Safe_v130                    = "0x69f4D1788e39c87893C980c06EdF4b7f686e2938"
	SafeL2_v130                  = "0xfb1bffC9d739B8D520DaF37dF666da4C687191EA"
	MultiSendCallOnly_v130       = "0xA1dabEF33b3B82c7814B6D82A79e50F4AC44102B"
	SafeSingletonFactory         = "0x914d7Fec6aaC8cd542e72Bca78B30650d45643d7"
	DeterministicDeploymentProxy = "0x4e59b44847b379578588920cA78FbF26c0B4956C"
	MultiSend_v130               = "0x998739BFdAAdde7C933B942a68053933098f9EDa"
	Permit2                      = "0x000000000022D473030F116dDEE9F6B43aC78BA3"
	SenderCreator                = "0x7fc98430eaedbb6070b35b39d798725049088348"
	EntryPoint                   = "0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789"
)

var (
	ProxyAdminAddr                   = common.HexToAddress(ProxyAdmin)
	WETH9Addr                        = common.HexToAddress(WETH9)
	KromaL1BlockAddr                 = common.HexToAddress(KromaL1Block)
	L2ToL1MessagePasserAddr          = common.HexToAddress(L2ToL1MessagePasser)
	L2CrossDomainMessengerAddr       = common.HexToAddress(L2CrossDomainMessenger)
	GasPriceOracleAddr               = common.HexToAddress(GasPriceOracle)
	ProtocolVaultAddr                = common.HexToAddress(ProtocolVault)
	L1FeeVaultAddr                   = common.HexToAddress(L1FeeVault)
	ValidatorRewardVaultAddr         = common.HexToAddress(ValidatorRewardVault)
	L2StandardBridgeAddr             = common.HexToAddress(L2StandardBridge)
	L2ERC721BridgeAddr               = common.HexToAddress(L2ERC721Bridge)
	KromaMintableERC20FactoryAddr    = common.HexToAddress(KromaMintableERC20Factory)
	KromaMintableERC721FactoryAddr   = common.HexToAddress(KromaMintableERC721Factory)
	Create2DeployerAddr              = common.HexToAddress(Create2Deployer)
	MultiCall3Addr                   = common.HexToAddress(MultiCall3)
	Safe_v130Addr                    = common.HexToAddress(Safe_v130)
	SafeL2_v130Addr                  = common.HexToAddress(SafeL2_v130)
	MultiSendCallOnly_v130Addr       = common.HexToAddress(MultiSendCallOnly_v130)
	SafeSingletonFactoryAddr         = common.HexToAddress(SafeSingletonFactory)
	DeterministicDeploymentProxyAddr = common.HexToAddress(DeterministicDeploymentProxy)
	MultiSend_v130Addr               = common.HexToAddress(MultiSend_v130)
	Permit2Addr                      = common.HexToAddress(Permit2)
	SenderCreatorAddr                = common.HexToAddress(SenderCreator)
	EntryPointAddr                   = common.HexToAddress(EntryPoint)

	Predeploys          = make(map[string]*oppredeploys.Predeploy)
	PredeploysByAddress = make(map[common.Address]*oppredeploys.Predeploy)
)

// IsProxied returns true for predeploys that will sit behind a proxy contract
func IsProxied(predeployAddr common.Address) bool {
	switch predeployAddr {
	case WETH9Addr:
	/* [Kroma: START]
	case GovernanceTokenAddr:
	[Kroma: END] */
	default:
		return true
	}
	return false
}

func init() {
	Predeploys["ProxyAdmin"] = &oppredeploys.Predeploy{Address: ProxyAdminAddr}
	Predeploys["WETH9"] = &oppredeploys.Predeploy{Address: WETH9Addr, ProxyDisabled: true}
	Predeploys["KromaL1Block"] = &oppredeploys.Predeploy{Address: KromaL1BlockAddr}
	Predeploys["L2ToL1MessagePasser"] = &oppredeploys.Predeploy{Address: L2ToL1MessagePasserAddr}
	Predeploys["L2CrossDomainMessenger"] = &oppredeploys.Predeploy{Address: L2CrossDomainMessengerAddr}
	Predeploys["GasPriceOracle"] = &oppredeploys.Predeploy{Address: GasPriceOracleAddr}
	Predeploys["ProtocolVault"] = &oppredeploys.Predeploy{Address: ProtocolVaultAddr}
	Predeploys["L1FeeVault"] = &oppredeploys.Predeploy{Address: L1FeeVaultAddr}
	Predeploys["ValidatorRewardVault"] = &oppredeploys.Predeploy{Address: ValidatorRewardVaultAddr}
	Predeploys["L2StandardBridge"] = &oppredeploys.Predeploy{Address: L2StandardBridgeAddr}
	/* [Kroma: START]
	Predeploys["GovernanceToken"] = &oppredeploys.Predeploy{
		Address: GovernanceTokenAddr,
		ProxyDisabled: true,
		Enabled: func(config oppredeploys.DeployConfig) bool {
			return config.GovernanceEnabled()
		},
	}
	*/
	Predeploys["L2ERC721Bridge"] = &oppredeploys.Predeploy{Address: L2ERC721BridgeAddr}
	Predeploys["KromaMintableERC20Factory"] = &oppredeploys.Predeploy{Address: KromaMintableERC20FactoryAddr}
	Predeploys["KromaMintableERC721Factory"] = &oppredeploys.Predeploy{Address: KromaMintableERC721FactoryAddr}
	Predeploys["Create2Deployer"] = &oppredeploys.Predeploy{
		Address:       Create2DeployerAddr,
		ProxyDisabled: true,
	}
	Predeploys["MultiCall3"] = &oppredeploys.Predeploy{
		Address:       MultiCall3Addr,
		ProxyDisabled: true,
	}
	Predeploys["Safe_v130"] = &oppredeploys.Predeploy{
		Address:       Safe_v130Addr,
		ProxyDisabled: true,
	}
	Predeploys["SafeL2_v130"] = &oppredeploys.Predeploy{
		Address:       SafeL2_v130Addr,
		ProxyDisabled: true,
	}
	Predeploys["MultiSendCallOnly_v130"] = &oppredeploys.Predeploy{
		Address:       MultiSendCallOnly_v130Addr,
		ProxyDisabled: true,
	}
	Predeploys["SafeSingletonFactory"] = &oppredeploys.Predeploy{
		Address:       SafeSingletonFactoryAddr,
		ProxyDisabled: true,
	}
	Predeploys["DeterministicDeploymentProxy"] = &oppredeploys.Predeploy{
		Address:       DeterministicDeploymentProxyAddr,
		ProxyDisabled: true,
	}
	Predeploys["MultiSend_v130"] = &oppredeploys.Predeploy{
		Address:       MultiSend_v130Addr,
		ProxyDisabled: true,
	}
	Predeploys["Permit2"] = &oppredeploys.Predeploy{
		Address:       Permit2Addr,
		ProxyDisabled: true,
	}
	Predeploys["SenderCreator"] = &oppredeploys.Predeploy{
		Address:       SenderCreatorAddr,
		ProxyDisabled: true,
	}
	Predeploys["EntryPoint"] = &oppredeploys.Predeploy{
		Address:       EntryPointAddr,
		ProxyDisabled: true,
	}

	for _, predeploy := range Predeploys {
		PredeploysByAddress[predeploy.Address] = predeploy
	}
}
