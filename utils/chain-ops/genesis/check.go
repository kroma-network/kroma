package genesis

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/wemixkanvas/kanvas/bindings/predeploys"
)

type StorageCheckMap = map[common.Hash]common.Hash

var (
	L2XDMOwnerSlot      = common.Hash{31: 0x33}
	ProxyAdminOwnerSlot = common.Hash{}

	// ExpectedStorageSlots is a map of predeploy addresses to the storage slots and values that are
	// expected to be set in those predeploys after the migration. It does not include any predeploys
	// that were not wiped. It also accounts for the 2 EIP-1967 storage slots in each contract.
	// It does _not_ include L1Block. L1Block is checked separately.
	ExpectedStorageSlots = map[common.Address]StorageCheckMap{
		predeploys.L2CrossDomainMessengerAddr: {
			// Slot 0x00 (0) is a combination of spacer_0_0_20, _initialized, and _initializing
			common.Hash{}: common.HexToHash("0x0000000000000000000000010000000000000000000000000000000000000000"),
			// Slot 0x33 (51) is _owner. Requires custom check, so set to a garbage value
			L2XDMOwnerSlot: common.HexToHash("0xbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbad0"),
			// Slot 0x97 (151) is _status
			common.Hash{31: 0x97}: common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
			// Slot 0xcc (204) is xDomainMsgSender
			common.Hash{31: 0xcc}: common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000dead"),
			// EIP-1967 storage slots
			AdminSlot:          common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000018"),
			ImplementationSlot: common.HexToHash("0x000000000000000000000000c0d3c0d3c0d3c0d3c0d3c0d3c0d3c0d3c0d30007"),
		},
		predeploys.L2StandardBridgeAddr:           eip1967Slots(predeploys.L2StandardBridgeAddr),
		predeploys.ProposerFeeVaultAddr:           eip1967Slots(predeploys.ProposerFeeVaultAddr),
		predeploys.KanvasMintableERC20FactoryAddr: eip1967Slots(predeploys.KanvasMintableERC20FactoryAddr),
		predeploys.GasPriceOracleAddr:             eip1967Slots(predeploys.GasPriceOracleAddr),
		//predeploys.L1BlockAddr:                       eip1967Slots(predeploys.L1BlockAddr),
		predeploys.L2ERC721BridgeAddr:              eip1967Slots(predeploys.L2ERC721BridgeAddr),
		predeploys.KanvasMintableERC721FactoryAddr: eip1967Slots(predeploys.KanvasMintableERC721FactoryAddr),
		// ProxyAdmin is not a proxy, and only has the _owner slot set.
		predeploys.ProxyAdminAddr: {
			// Slot 0x00 (0) is _owner. Requires custom check, so set to a garbage value
			ProxyAdminOwnerSlot: common.HexToHash("0xbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbadbad0"),

			// EIP-1967 storage slots
			AdminSlot:          common.HexToHash("0x0000000000000000000000004200000000000000000000000000000000000018"),
			ImplementationSlot: common.HexToHash("0x000000000000000000000000c0d3c0d3c0d3c0d3c0d3c0d3c0d3c0d3c0d30018"),
		},
		predeploys.BaseFeeVaultAddr: eip1967Slots(predeploys.BaseFeeVaultAddr),
		predeploys.L1FeeVaultAddr:   eip1967Slots(predeploys.L1FeeVaultAddr),
	}
)

func eip1967Slots(address common.Address) StorageCheckMap {
	codeAddr, err := AddressToCodeNamespace(address)
	if err != nil {
		panic(err)
	}
	return StorageCheckMap{
		AdminSlot:          predeploys.ProxyAdminAddr.Hash(),
		ImplementationSlot: codeAddr.Hash(),
	}
}
