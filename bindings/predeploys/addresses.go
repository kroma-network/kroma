package predeploys

import "github.com/ethereum/go-ethereum/common"

const (
	ProxyAdmin                  = "0x4200000000000000000000000000000000000000"
	WETH9                       = "0x4200000000000000000000000000000000000001"
	L1Block                     = "0x4200000000000000000000000000000000000002"
	L2ToL1MessagePasser         = "0x4200000000000000000000000000000000000003"
	L2CrossDomainMessenger      = "0x4200000000000000000000000000000000000004"
	GasPriceOracle              = "0x4200000000000000000000000000000000000005"
	BaseFeeVault                = "0x4200000000000000000000000000000000000006"
	L1FeeVault                  = "0x4200000000000000000000000000000000000007"
	ProposerFeeVault            = "0x4200000000000000000000000000000000000008"
	L2StandardBridge            = "0x4200000000000000000000000000000000000009"
	L2ERC721Bridge              = "0x420000000000000000000000000000000000000A"
	KanvasMintableERC20Factory  = "0x420000000000000000000000000000000000000B"
	KanvasMintableERC721Factory = "0x420000000000000000000000000000000000000C"
)

var (
	ProxyAdminAddr                  = common.HexToAddress(ProxyAdmin)
	WETH9Addr                       = common.HexToAddress(WETH9)
	L1BlockAddr                     = common.HexToAddress(L1Block)
	L2ToL1MessagePasserAddr         = common.HexToAddress(L2ToL1MessagePasser)
	L2CrossDomainMessengerAddr      = common.HexToAddress(L2CrossDomainMessenger)
	GasPriceOracleAddr              = common.HexToAddress(GasPriceOracle)
	BaseFeeVaultAddr                = common.HexToAddress(BaseFeeVault)
	L1FeeVaultAddr                  = common.HexToAddress(L1FeeVault)
	ProposerFeeVaultAddr            = common.HexToAddress(ProposerFeeVault)
	L2StandardBridgeAddr            = common.HexToAddress(L2StandardBridge)
	L2ERC721BridgeAddr              = common.HexToAddress(L2ERC721Bridge)
	KanvasMintableERC20FactoryAddr  = common.HexToAddress(KanvasMintableERC20Factory)
	KanvasMintableERC721FactoryAddr = common.HexToAddress(KanvasMintableERC721Factory)

	Predeploys = make(map[string]*common.Address)
)

func init() {
	Predeploys["ProxyAdmin"] = &ProxyAdminAddr
	Predeploys["WETH9"] = &WETH9Addr
	Predeploys["L1Block"] = &L1BlockAddr
	Predeploys["L2ToL1MessagePasser"] = &L2ToL1MessagePasserAddr
	Predeploys["L2CrossDomainMessenger"] = &L2CrossDomainMessengerAddr
	Predeploys["GasPriceOracle"] = &GasPriceOracleAddr
	Predeploys["BaseFeeVault"] = &BaseFeeVaultAddr
	Predeploys["L1FeeVault"] = &L1FeeVaultAddr
	Predeploys["ProposerFeeVault"] = &ProposerFeeVaultAddr
	Predeploys["L2StandardBridge"] = &L2StandardBridgeAddr
	Predeploys["L2ERC721Bridge"] = &L2ERC721BridgeAddr
	Predeploys["KanvasMintableERC20Factory"] = &KanvasMintableERC20FactoryAddr
	Predeploys["KanvasMintableERC721Factory"] = &KanvasMintableERC721FactoryAddr
}
