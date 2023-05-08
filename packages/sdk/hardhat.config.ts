import '@nomiclabs/hardhat-ethers'
import '@nomiclabs/hardhat-waffle'
import dotenv from 'dotenv'
import { HardhatUserConfig } from 'hardhat/types'
import 'hardhat-deploy'

import './src/tasks'

// Load environment variables
dotenv.config()

const config: HardhatUserConfig = {
  solidity: {
    version: '0.8.15',
  },
  paths: {
    sources: './test/contracts',
  },
  networks: {
    sepolia: {
      url: process.env.L1_RPC || 'https://rpc.sepolia.org',
      chainId: 11155111,
      accounts: process.env.PRIVATE_KEY ? [process.env.PRIVATE_KEY] : [],
    },
    kromaSepolia: {
      url: 'https://api.sepolia.kroma.network',
      chainId: 2357,
      accounts: process.env.PRIVATE_KEY ? [process.env.PRIVATE_KEY] : [],
    },
    devnetL1: {
      url: 'http://localhost:8545',
      chainId: 900,
      accounts: [
        'ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80',
      ],
    },
    devnetL2: {
      url: 'http://127.0.0.1:9545',
      chainId: 901,
      accounts: [
        'ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80',
      ],
    },
  },
}

export default config
