import { ethers } from 'ethers'

/**
 * Predeploys are Solidity contracts that are injected into the initial L2 state and provide
 * various useful functions.
 */
export const predeploys = {
  ProxyAdmin: '0x4200000000000000000000000000000000000000',
  WETH9: '0x4200000000000000000000000000000000000001',
  L1Block: '0x4200000000000000000000000000000000000002',
  L2ToL1MessagePasser: '0x4200000000000000000000000000000000000003',
  L2CrossDomainMessenger: '0x4200000000000000000000000000000000000004',
  GasPriceOracle: '0x4200000000000000000000000000000000000005',
  ProtocolVault: '0x4200000000000000000000000000000000000006',
  L1FeeVault: '0x4200000000000000000000000000000000000007',
  ValidatorRewardVault: '0x4200000000000000000000000000000000000008',
  L2StandardBridge: '0x4200000000000000000000000000000000000009',
  L2ERC721Bridge: '0x420000000000000000000000000000000000000A',
  KromaMintableERC20Factory: '0x420000000000000000000000000000000000000B',
  KromaMintableERC721Factory: '0x420000000000000000000000000000000000000C',
  Create2Deployer: '0x13b0D85CcB8bf860b6b79AF3029fCA081AE9beF2',
}

const uint128Max = ethers.BigNumber.from('0xffffffffffffffffffffffffffffffff')

export const defaultResourceConfig = {
  maxResourceLimit: 20_000_000,
  elasticityMultiplier: 10,
  baseFeeMaxChangeDenominator: 8,
  minimumBaseFee: ethers.utils.parseUnits('1', 'gwei'),
  systemTxMaxGas: 1_000_000,
  maximumBaseFee: uint128Max,
}
