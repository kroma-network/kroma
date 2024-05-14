COMPOSEFLAGS=-d
ITESTS_L2_HOST=http://localhost:9545
VERSION := $(shell git describe --tags --abbrev=0 --match v* 2> /dev/null || echo 'v0.0.0')
GIT_COMMIT := $(shell git rev-parse --short=8 HEAD)

# Requires at least Python v3.9; specify a minor version below if needed
PYTHON?=python3

LD_FLAGS_ARGS +=-X main.Version=$(VERSION)
LD_FLAGS_ARGS +=-X main.Meta=$(GIT_COMMIT)
LD_FLAGS := -ldflags "$(LD_FLAGS_ARGS)"

build:
	GO111MODULE=on go build -v $(LD_FLAGS) -o bin/op-node ./op-node/cmd/main.go
	GO111MODULE=on go build -v $(LD_FLAGS) -o bin/op-stateviz ./op-node/cmd/stateviz/main.go
	GO111MODULE=on go build -v $(LD_FLAGS) -o bin/op-batcher ./op-batcher/cmd/main.go
	GO111MODULE=on go build -v $(LD_FLAGS) -o bin/kroma-validator ./kroma-validator/cmd/main.go
.PHONY: build

lint-go:
	golangci-lint run -E goimports,sqlclosecheck,bodyclose,asciicheck,misspell,errorlint --timeout 5m -e "errors.As" -e "errors.Is" ./...
.PHONY: lint-go

test:
	go test ./op-bindings/...
	go test ./op-batcher/...
	go test ./op-node/...
	go test ./op-service/...
	go test ./op-chain-ops/...
	go test ./kroma-bindings/...
	go test ./kroma-chain-ops/...
	go test ./kroma-validator/...
	go test ./op-e2e/... -timeout 30m # a maximum of 30min in a CI
	pnpm test
.PHONY: test

golang-docker:
	# We don't use a buildx builder here, and just load directly into regular docker, for convenience.
	GIT_COMMIT=$$(git rev-parse HEAD) \
	GIT_DATE=$$(git show -s --format='%ct') \
	IMAGE_TAGS=$$(git rev-parse HEAD),latest \
	docker buildx bake \
			--progress plain \
			--load \
			-f docker-bake.hcl \
			op-node op-batcher kroma-validator
.PHONY: golang-docker

submodules:
	git submodule update --init --recursive
.PHONY: submodules

bindings:
	make -C ./kroma-bindings
.PHONY: bindings

contracts-snapshot:
	@(cd ./packages/contracts && pnpm gas-snapshot && pnpm storage-snapshot)
.PHONY: gas-snapshot

mod-tidy:
	# Below GOPRIVATE line allows mod-tidy to be run immediately after
	# releasing new versions. This bypasses the Go modules proxy, which
	# can take a while to index new versions.
	#
	# See https://proxy.golang.org/ for more info.
	export GOPRIVATE="github.com/kroma-network" && go mod tidy
.PHONY: mod-tidy

clean:
	rm -rf ./bin
.PHONY: clean

nuke: clean devnet-clean
	git clean -Xdf
.PHONY: nuke

pre-devnet: submodules
	@if ! [ -x "$(command -v geth)" ]; then \
		make install-geth; \
	fi
.PHONY: pre-devnet

devnet-up: pre-devnet
	./ops/scripts/newer-file.sh .devnet/allocs-l1.json ./packages/contracts \
		|| make devnet-allocs
	PYTHONPATH=./kroma-devnet $(PYTHON) ./kroma-devnet/main.py --monorepo-dir=.
.PHONY: devnet-up

# alias for devnet-up
devnet-up-deploy: devnet-up

devnet-test: pre-devnet
	PYTHONPATH=./kroma-devnet $(PYTHON) ./kroma-devnet/main.py --monorepo-dir=. --test
.PHONY: devnet-test

devnet-down:
	@(cd ./ops-devnet && GENESIS_TIMESTAMP=$(shell date +%s) docker compose stop)
.PHONY: devnet-down

devnet-clean:
	rm -rf ./packages/contracts/deployments/devnetL1
	rm -rf ./.devnet
	cd ./ops-devnet && docker compose down
	docker image ls 'ops-devnet*' --format='{{.Repository}}' | xargs -r docker rmi
	docker volume ls --filter name=ops-devnet --format='{{.Name}}' | xargs -r docker volume rm
.PHONY: devnet-clean

devnet-allocs: pre-devnet
	PYTHONPATH=./kroma-devnet $(PYTHON) ./kroma-devnet/main.py --monorepo-dir=. --allocs

devnet-logs:
	@(cd ./ops-devnet && docker compose logs -f)
.PHONY: devnet-logs

# Remove the baseline-commit to generate a base reading & show all issues
semgrep:
	$(eval DEV_REF := $(shell git rev-parse develop))
	SEMGREP_REPO_NAME=ethereum-optimism/optimism semgrep ci --baseline-commit=$(DEV_REF)
.PHONY: semgrep

clean-node-modules:
	rm -rf node_modules
	rm -rf packages/**/node_modules

update-geth:
	@ETH_GETH=$$(go list -m -f '{{.Path}}@{{.Version}}' github.com/ethereum/go-ethereum); \
	KROMA_GETH=$$(go list -m -f '{{.Path}}@{{.Version}}' github.com/kroma-network/go-ethereum@dev); \
	go mod edit -replace $$ETH_GETH=$$KROMA_GETH
	@go mod tidy
.PHONY: update-geth

install-geth:
	./ops/scripts/geth-version-checker.sh && \
	 	(echo "Geth versions match, not installing geth..."; true) || \
 		(echo "Versions do not match, installing geth!"; \
 			go install -v github.com/ethereum/go-ethereum/cmd/geth@$(shell jq -r .geth < versions.json); \
 			echo "Installed geth!"; true)
.PHONY: install-geth
