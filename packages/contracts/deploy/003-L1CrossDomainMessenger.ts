import { DeployFunction } from 'hardhat-deploy/dist/types'

import {
  assertContractVariable,
  deploy,
  getContractFromArtifact,
} from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const KanvasPortalProxy = await getContractFromArtifact(
    hre,
    'KanvasPortalProxy'
  )

  await deploy(hre, 'L1CrossDomainMessenger', {
    args: [KanvasPortalProxy.address],
    isProxyImpl: true,
    initArgs: [],
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'PORTAL',
        KanvasPortalProxy.address
      )
    },
  })
}

deployFn.tags = ['L1CrossDomainMessenger', 'setup']

export default deployFn
