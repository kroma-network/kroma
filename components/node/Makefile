GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')
VERSION := v0.1.0

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X github.com/kroma-network/kroma/node/version.Version=$(VERSION)
LDFLAGSSTRING +=-X github.com/kroma-network/kroma/node/version.Meta=$(VERSION_META)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

build:
	env GO111MODULE=on go build -v $(LDFLAGS) -o ./bin/kroma-node ./cmd/main.go

clean:
	rm bin/kroma-node

test:
	go test -v ./...

lint:
	golangci-lint run -E goimports,sqlclosecheck,bodyclose,asciicheck,misspell,errorlint -e "errors.As" -e "errors.Is"

fuzz:
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzExecutionPayloadUnmarshal ./eth
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzExecutionPayloadMarshalUnmarshal ./eth
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzOBP01 ./eth
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzL1InfoRoundTrip ./rollup/derive
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzL1InfoAgainstContract ./rollup/derive
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzUnmarshallLogEvent ./rollup/derive
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzParseFrames ./rollup/derive
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzFrameUnmarshalBinary ./rollup/derive
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzBatchRoundTrip ./rollup/derive
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzDeriveDepositsRoundTrip ./rollup/derive
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzDeriveDepositsBadVersion ./rollup/derive
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzParseL1InfoDepositTxDataValid ./rollup/derive
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzParseL1InfoDepositTxDataBadLength ./rollup/derive
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzRejectCreateBlockBadTimestamp ./rollup/driver
	go test -run NOTAREALTEST -v -fuzztime 10s -fuzz FuzzDecodeDepositTxDataToL1Info ./rollup/driver


.PHONY: \
	build \
	clean \
	test \
	lint \
	fuzz
