SHELL := /bin/bash

pkg := bindings

all: version mkdir bindings more

bindings: basefee-vault-bindings \
		colosseum-bindings \
		erc20-bindings \
		gas-price-oracle-bindings \
		kroma-mintable-erc20-bindings \
		kroma-mintable-erc20-factory-bindings \
		kroma-mintable-erc721-factory-bindings \
		kroma-portal-bindings \
		l1-cross-domain-messenger-bindings \
		l1-erc721-bridge-bindings \
		l1-fee-vault-bindings \
		l1-standard-bridge-bindings \
		l1block-bindings \
		l2-cross-domain-messenger-bindings \
		l2-erc721-bridge-bindings \
		l2-output-oracle-bindings \
		l2-standard-bridge-bindings \
		l2-to-l1-message-passer-bindings \
		proposer-fee-vault-bindings \
		proxy-admin-bindings \
		proxy-bindings \
		system-config-bindings \
		validator-pool-bindings \
		weth9-bindings \
		zk-merkle-trie-bindings \
		zk-verifier-bindings

version:
	forge --version

compile:
	cd ../packages/contracts/ && \
		yarn clean && \
		npx hardhat compile

more:
	go run ./gen/main.go \
		-artifacts ../packages/contracts/artifacts \
		-out ./bindings \
		-contracts \
		Colosseum,\
KromaMintableERC20Factory,\
KromaMintableERC721Factory,\
KromaPortal,\
L1Block,\
L1CrossDomainMessenger,\
L2CrossDomainMessenger,\
L2ERC721Bridge,\
L2StandardBridge,\
L2ToL1MessagePasser,\
ValidatorRewardVault,\
Proxy,\
ProxyAdmin,\
SystemConfig,\
WETH9\
		-package bindings

mkdir:
	mkdir -p bin $(pkg)

clean:
	rm -rf bin $(pkg)

test:
	go test ./...

basefee-vault-bindings: compile
	bash ./gen_bindings.sh contracts/L2/ProtocolVault.sol:ProtocolVault $(pkg)

colosseum-bindings: compile
	bash ./gen_bindings.sh contracts/L1/Colosseum.sol:Colosseum $(pkg)

erc20-bindings: compile
	bash ./gen_bindings.sh node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol:ERC20 $(pkg)

gas-price-oracle-bindings: compile
	bash ./gen_bindings.sh contracts/L2/GasPriceOracle.sol:GasPriceOracle $(pkg)

kroma-mintable-erc20-bindings: compile
	bash ./gen_bindings.sh contracts/universal/KromaMintableERC20.sol:KromaMintableERC20 $(pkg)

kroma-mintable-erc20-factory-bindings: compile
	bash ./gen_bindings.sh contracts/universal/KromaMintableERC20Factory.sol:KromaMintableERC20Factory $(pkg)

kroma-mintable-erc721-factory-bindings: compile
	bash ./gen_bindings.sh contracts/universal/KromaMintableERC721Factory.sol:KromaMintableERC721Factory $(pkg)

kroma-portal-bindings: compile
	bash ./gen_bindings.sh contracts/L1/KromaPortal.sol:KromaPortal $(pkg)

l1-cross-domain-messenger-bindings: compile
	bash ./gen_bindings.sh contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger $(pkg)

l1-erc721-bridge-bindings: compile
	bash ./gen_bindings.sh contracts/L1/L1ERC721Bridge.sol:L1ERC721Bridge $(pkg)

l1-fee-vault-bindings: compile
	bash ./gen_bindings.sh contracts/L2/ProposerRewardVault.sol:ProposerRewardVault $(pkg)

l1-standard-bridge-bindings: compile
	bash ./gen_bindings.sh contracts/L1/L1StandardBridge.sol:L1StandardBridge $(pkg)

l1block-bindings: compile
	bash ./gen_bindings.sh contracts/L2/L1Block.sol:L1Block $(pkg)

l2-cross-domain-messenger-bindings: compile
	bash ./gen_bindings.sh contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger $(pkg)

l2-erc721-bridge-bindings: compile
	bash ./gen_bindings.sh contracts/L2/L2ERC721Bridge.sol:L2ERC721Bridge $(pkg)

l2-output-oracle-bindings: compile
	bash ./gen_bindings.sh contracts/L1/L2OutputOracle.sol:L2OutputOracle $(pkg)

l2-standard-bridge-bindings: compile
	bash ./gen_bindings.sh contracts/L2/L2StandardBridge.sol:L2StandardBridge $(pkg)

l2-to-l1-message-passer-bindings: compile
	bash ./gen_bindings.sh contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser $(pkg)

proposer-fee-vault-bindings: compile
	bash ./gen_bindings.sh contracts/L2/ValidatorRewardVault.sol:ValidatorRewardVault $(pkg)

proxy-admin-bindings: compile
	bash ./gen_bindings.sh contracts/universal/ProxyAdmin.sol:ProxyAdmin $(pkg)

proxy-bindings: compile
	bash ./gen_bindings.sh contracts/universal/Proxy.sol:Proxy $(pkg)

system-config-bindings: compile
	bash ./gen_bindings.sh contracts/L1/SystemConfig.sol:SystemConfig $(pkg)

validator-pool-bindings: compile
	bash ./gen_bindings.sh contracts/L1/ValidatorPool.sol:ValidatorPool $(pkg)

weth9-bindings: compile
	bash ./gen_bindings.sh contracts/vendor/WETH9.sol:WETH9 $(pkg)

zk-merkle-trie-bindings: compile
	bash ./gen_bindings.sh contracts/L1/ZKMerkleTrie.sol:ZKMerkleTrie $(pkg)

zk-verifier-bindings: compile
	bash ./gen_bindings.sh contracts/L1/ZKVerifier.sol:ZKVerifier $(pkg)
