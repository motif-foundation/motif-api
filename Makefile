# project related vars
PROJECT := $(shell basename "$(PWD)")

# go related vars
GO_BASE := $(shell pwd)
GO_BIN := $(CURDIR)/build

# compile time variables will be injected into the app
APP_VERSION := 1.1.0
BUILD_DATE := $(shell date)
BUILD_COMPILER := $(shell go version)
BUILD_COMMIT := $(shell git show --format="%H" --no-patch)
BUILD_COMMIT_TIME := $(shell git show --format="%cD" --no-patch)

## server: Make the API server as build/apiserver
server:
	go build \
	-ldflags="-X 'motif-api/cmd/apiserver/build.Version=$(APP_VERSION)' -X 'motif-api/cmd/apiserver/build.Time=$(BUILD_DATE)' -X 'motif-api/cmd/apiserver/build.Compiler=$(BUILD_COMPILER)' -X 'motif-api/cmd/apiserver/build.Commit=$(BUILD_COMMIT)' -X 'motif-api/cmd/apiserver/build.CommitTime=$(BUILD_COMMIT_TIME)'" \
	-o $(GO_BIN)/apiserver \
	./cmd/apiserver

test:
	go test \
	-ldflags="-X 'motif-api/cmd/apiserver/build.Version=$(APP_VERSION)' -X 'motif-api/cmd/apiserver/build.Time=$(BUILD_DATE)' -X 'motif-api/cmd/apiserver/build.Compiler=$(BUILD_COMPILER)' -X 'motif-api/cmd/apiserver/build.Commit=$(BUILD_COMMIT)' -X 'motif-api/cmd/apiserver/build.CommitTime=$(BUILD_COMMIT_TIME)'" \
	./...

.PHONY: help test
all: help
help: Makefile
	@echo
	@echo "Choose a make command in "$(PROJECT)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
