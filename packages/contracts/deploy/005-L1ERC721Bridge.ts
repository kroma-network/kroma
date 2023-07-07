import { DeployFunction } from 'hardhat-deploy/dist/types'

import { predeploys } from '../src'
import {
  assertContractVariable,
  deploy,
  getContractFromArtifact,
} from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const L1CrossDomainMessengerProxy = await getContractFromArtifact(
    hre,
    'L1CrossDomainMessengerProxy'
  )

  await deploy(hre, 'L1ERC721Bridge', {
    args: [L1CrossDomainMessengerProxy.address, predeploys.L2ERC721Bridge],
    isProxyImpl: true,
    postDeployAction: async (contract) => {
      await assertContractVariable(
        contract,
        'MESSENGER',
        L1CrossDomainMessengerProxy.address
      )
    },
  })
}

deployFn.tags = ['L1ERC721Bridge', 'setup', 'l1']

export default deployFn
