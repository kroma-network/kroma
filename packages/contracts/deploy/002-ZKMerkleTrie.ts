import '@nomiclabs/hardhat-ethers'
import { poseidonContract } from 'circomlibjs'
import { DeployFunction } from 'hardhat-deploy/dist/types'

import { deploy, getDeploymentAddress } from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const abi = poseidonContract.generateABI(2)
  const bytecode = poseidonContract.createCode(2)

  await deploy(hre, 'Poseidon2', {
    contract: {
      abi,
      bytecode,
    },
  })

  const poseidon2 = await getDeploymentAddress(hre, 'Poseidon2')

  await deploy(hre, 'ZKMerkleTrie', {
    args: [poseidon2],
  })
}

deployFn.tags = ['ZKMerkleTrie', 'Poseidon2', 'setup', 'l1']

export default deployFn
