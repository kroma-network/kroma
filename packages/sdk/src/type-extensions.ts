import 'hardhat/types/runtime'
import { DeploymentsExtension } from 'hardhat-deploy/types'

declare module 'hardhat/types/runtime' {
  interface HardhatRuntimeEnvironment {
    deployments: DeploymentsExtension
  }
}
