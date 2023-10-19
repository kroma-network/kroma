import { DeployFunction } from 'hardhat-deploy/dist/types'

import { deployProxy, getDeploymentAddress } from '../../src/deploy-utils'

const PROXY_NAMES = [
  'SecurityCouncilTokenProxy',
  'TimeLockProxy',
  'UpgradeGovernorProxy',
]

const deployFn: DeployFunction = async (hre) => {
  const proxyAdminProxy = await getDeploymentAddress(hre, 'ProxyAdminProxy')

  for (const proxyName of PROXY_NAMES) {
    await deployProxy(hre, proxyName, proxyAdminProxy)
  }
}

deployFn.tags = [...PROXY_NAMES, 'setup', 'l2']

export default deployFn
