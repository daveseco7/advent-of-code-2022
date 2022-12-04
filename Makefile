.PHONY: build

LINTER_VERSION=v1.31.0

default: unit-tests

get-linter:
	command -v golangci-lint || curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b ${GOPATH}/bin $(LINTER_VERSION)

## lint: download/install golangci-lint and analyse the source code with the configuration in .golangci.yml
lint: get-linter
	golangci-lint run --timeout=10m

unit-tests: lint
	go fmt ./...
	go test -vet all -shuffle=on ./...


mod:
	go mod vendor -v

tidy:
	go mod tidy -v
