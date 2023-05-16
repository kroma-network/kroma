import dotenv from 'dotenv'
import { ethers } from 'ethers'
import { HardhatUserConfig } from 'hardhat/config'

// Hardhat plugins
import '@foundry-rs/hardhat-forge'
import '@kroma-network/hardhat-deploy-config'
import '@nomiclabs/hardhat-ethers'
import 'hardhat-deploy'

// Hardhat tasks
import './tasks'

// Deploy configuration
import { deployConfigSpec } from './src/deploy-config'

// Load environment variables
dotenv.config()

const PRIVATE_KEY_DEPLOYER_DEVNET =
  'ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80'

const config: HardhatUserConfig = {
  networks: {
    hardhat: {
      live: false,
    },
    mainnet: {
      url: process.env.L1_RPC_MAINNET || '',
      accounts: [
        process.env.PRIVATE_KEY_DEPLOYER_MAINNET || ethers.constants.HashZero,
      ],
    },
    sepolia: {
      chainId: 11155111,
      url: process.env.L1_RPC_SEPOLIA || '',
      accounts: [
        process.env.PRIVATE_KEY_DEPLOYER_SEPOLIA || ethers.constants.HashZero,
      ],
      companionNetworks: {
        l2: 'kromaSepolia',
      },
    },
    kromaSepolia: {
      chainId: 2357,
      url: process.env.L2_RPC_KROMA_SEPOLIA || '',
      accounts: [
        process.env.PRIVATE_KEY_DEPLOYER_SEPOLIA || ethers.constants.HashZero,
      ],
      companionNetworks: {
        l1: 'sepolia',
      },
    },
    devnetL1: {
      live: false,
      url: 'http://localhost:8545',
      accounts: [PRIVATE_KEY_DEPLOYER_DEVNET],
    },
    devnetL2: {
      live: false,
      url: process.env.RPC_URL || 'http://localhost:9545',
      accounts: [PRIVATE_KEY_DEPLOYER_DEVNET],
    },
    easel: {
      chainId: 7789,
      url: process.env.L1_RPC_EASEL || '',
      accounts: [PRIVATE_KEY_DEPLOYER_DEVNET],
    },
  },
  foundry: {
    buildInfo: true,
  },
  paths: {
    deploy: './deploy',
    deployments: './deployments',
    deployConfig: './deploy-config',
  },
  namedAccounts: {
    deployer: {
      default: 0,
    },
  },
  deployConfigSpec,
  external: {
    contracts: [
      {
        artifacts: '../contracts/artifacts',
      },
    ],
    deployments: {
      mainnet: ['../contracts/deployments/mainnet'],
      sepolia: ['../contracts/deployments/sepolia'],
      easel: ['../contracts/deployments/easel'],
    },
  },
  solidity: {
    compilers: [
      {
        version: '0.8.15',
        settings: {
          optimizer: { enabled: true, runs: 10_000 },
        },
      },
      {
        version: '0.5.17', // Required for WETH9
        settings: {
          optimizer: { enabled: true, runs: 10_000 },
        },
      },
    ],
    settings: {
      metadata: {
        bytecodeHash:
          process.env.FOUNDRY_PROFILE === 'echidna' ? 'ipfs' : 'none',
      },
      outputSelection: {
        '*': {
          '*': ['metadata', 'storageLayout'],
        },
      },
    },
  },
}

export default config
