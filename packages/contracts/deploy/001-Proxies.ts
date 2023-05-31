import { DeployFunction } from 'hardhat-deploy/dist/types'

import { deployProxy, getDeploymentAddress } from '../src/deploy-utils'

const PROXY_NAMES = [
  'SystemConfigProxy',
  'KromaPortalProxy',
  'ValidatorPoolProxy',
  'L2OutputOracleProxy',
  'L1CrossDomainMessengerProxy',
  'L1StandardBridgeProxy',
  'L1ERC721BridgeProxy',
  'KromaMintableERC20FactoryProxy',
  'ColosseumProxy',
]

const deployFn: DeployFunction = async (hre) => {
  const proxyAdmin = await getDeploymentAddress(hre, 'ProxyAdmin')

  for (const proxyName of PROXY_NAMES) {
    await deployProxy(hre, proxyName, proxyAdmin)
  }
}

deployFn.tags = [...PROXY_NAMES, 'setup', 'l1']

export default deployFn
