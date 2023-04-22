import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getContractFromArtifact,
} from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const KromaPortalProxy = await getContractFromArtifact(
    hre,
    'KromaPortalProxy'
  )

  await deploy(hre, 'L1CrossDomainMessenger', {
    args: [KromaPortalProxy.address],
    isProxyImpl: true,
    initArgs: [],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'PORTAL',
        KromaPortalProxy.address
      )
    },
  })
}

deployFn.tags = ['L1CrossDomainMessenger', 'setup']

export default deployFn
