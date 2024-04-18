import '@nomiclabs/hardhat-ethers'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { deploy } from '../../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  await deploy(hre, 'ZKVerifier', {
    args: [
      hre.deployConfig.zkVerifierHashScalar,
      hre.deployConfig.zkVerifierM56Px,
      hre.deployConfig.zkVerifierM56Py,
    ],
    isProxyImpl: true,
  })
}

deployFn.tags = ['ZKVerifier', 'setup', 'l1']

export default deployFn
