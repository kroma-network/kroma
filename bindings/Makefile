SHELL := /bin/bash

pkg := bindings

all: version mkdir bindings more

bindings: basefee-vault-bindings \
		colosseum-bindings \
		erc20-bindings \
		gas-price-oracle-bindings \
		kanvas-mintable-erc20-bindings \
		kanvas-mintable-erc20-factory-bindings \
		kanvas-mintable-erc721-factory-bindings \
		kanvas-portal-bindings \
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
KanvasMintableERC20Factory,\
KanvasMintableERC721Factory,\
KanvasPortal,\
L1Block,\
L1CrossDomainMessenger,\
L2CrossDomainMessenger,\
L2ERC721Bridge,\
L2StandardBridge,\
L2ToL1MessagePasser,\
ProposerFeeVault,\
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
	bash ./gen_bindings.sh contracts/L2/BaseFeeVault.sol:BaseFeeVault $(pkg)

colosseum-bindings: compile
	bash ./gen_bindings.sh contracts/L1/Colosseum.sol:Colosseum $(pkg)

erc20-bindings: compile
	bash ./gen_bindings.sh node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol:ERC20 $(pkg)

gas-price-oracle-bindings: compile
	bash ./gen_bindings.sh contracts/L2/GasPriceOracle.sol:GasPriceOracle $(pkg)

kanvas-mintable-erc20-bindings: compile
	bash ./gen_bindings.sh contracts/universal/KanvasMintableERC20.sol:KanvasMintableERC20 $(pkg)

kanvas-mintable-erc20-factory-bindings: compile
	bash ./gen_bindings.sh contracts/universal/KanvasMintableERC20Factory.sol:KanvasMintableERC20Factory $(pkg)

kanvas-mintable-erc721-factory-bindings: compile
	bash ./gen_bindings.sh contracts/universal/KanvasMintableERC721Factory.sol:KanvasMintableERC721Factory $(pkg)

kanvas-portal-bindings: compile
	bash ./gen_bindings.sh contracts/L1/KanvasPortal.sol:KanvasPortal $(pkg)

l1-cross-domain-messenger-bindings: compile
	bash ./gen_bindings.sh contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger $(pkg)

l1-erc721-bridge-bindings: compile
	bash ./gen_bindings.sh contracts/L1/L1ERC721Bridge.sol:L1ERC721Bridge $(pkg)

l1-fee-vault-bindings: compile
	bash ./gen_bindings.sh contracts/L2/L1FeeVault.sol:L1FeeVault $(pkg)

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
	bash ./gen_bindings.sh contracts/L2/ProposerFeeVault.sol:ProposerFeeVault $(pkg)

proxy-admin-bindings: compile
	bash ./gen_bindings.sh contracts/universal/ProxyAdmin.sol:ProxyAdmin $(pkg)

proxy-bindings: compile
	bash ./gen_bindings.sh contracts/universal/Proxy.sol:Proxy $(pkg)

system-config-bindings: compile
	bash ./gen_bindings.sh contracts/L1/SystemConfig.sol:SystemConfig $(pkg)

weth9-bindings: compile
	bash ./gen_bindings.sh contracts/vendor/WETH9.sol:WETH9 $(pkg)

zk-merkle-trie-bindings: compile
	bash ./gen_bindings.sh contracts/L1/ZKMerkleTrie.sol:ZKMerkleTrie $(pkg)

zk-verifier-bindings: compile
	bash ./gen_bindings.sh contracts/L1/ZKVerifier.sol:ZKVerifier $(pkg)
