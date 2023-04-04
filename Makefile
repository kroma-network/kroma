VERSION := $(shell git describe --tags --abbrev=0 --match v* 2> /dev/null || git rev-parse --short HEAD)
GIT_COMMIT := $(shell git rev-parse HEAD)
GIT_DATE := $(shell git show -s --format='%ct')

LD_FLAGS_ARGS +=-X main.Version=$(VERSION)
LD_FLAGS_ARGS +=-X main.GitCommit=$(GIT_COMMIT)
LD_FLAGS_ARGS +=-X main.GitDate=$(GIT_DATE)
LD_FLAGS := -ldflags "$(LD_FLAGS_ARGS)"

build:
	GO111MODULE=on go build -v $(LD_FLAGS) -o bin/kanvas-node ./components/node/cmd/main.go
	GO111MODULE=on go build -v $(LD_FLAGS) -o bin/kanvas-stateviz ./components/node/cmd/stateviz/main.go
	GO111MODULE=on go build -v $(LD_FLAGS) -o bin/kanvas-batcher ./components/batcher/cmd/main.go
	GO111MODULE=on go build -v $(LD_FLAGS) -o bin/kanvas-validator ./components/validator/cmd/main.go
.PHONY: build

clean:
	@rm -rf bin/*
.PHONY: clean

test:
	go test ./bindings/...
	go test ./components/...
	go test ./utils/...
	go test ./e2e/... -timeout 30m # requires a minimum of 30min in a CI
	yarn test
.PHONY: test

lint:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2 \
	&& golangci-lint run
.PHONY: lint

bindings:
	make -C ./bindings
.PHONY: bindings

contracts-snapshot:
	@(cd ./packages/contracts && yarn gas-snapshot && yarn storage-snapshot)
.PHONY: gas-snapshot

devnet-up:
	@bash ./ops-devnet/devnet-up.sh
.PHONY: devnet-up

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
