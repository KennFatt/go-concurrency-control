.PHONY: build clean

APP_BIN=./bin/app
APP_SRC=./cmd/server

GO_ENV=CGO_ENABLED=0 GO111MODULE=on
GO_FLAGS=-ldflags="-extldflags -static -s -w"
GO_BIN=$(GO_ENV) $(shell which go)

setup:
	@$(GO_BIN) install entgo.io/ent/cmd/ent@latest
	@$(GO_BIN) install golang.org/x/tools/cmd/goimports@latest
	@$(GO_BIN) install github.com/mgechev/revive@latest
	@$(GO_BIN) install github.com/alexkohler/prealloc@latest

generate: setup
	@$(GO_BIN) generate ./...

build: generate
	@$(GO_BIN) build  $(GO_FLAGS) -o $(APP_BIN) $(APP_SRC)

clean:
	@$(GO_BIN) clean ./...
	@rm -f $(APP_BIN)

lint: generate
	@goimports -w -l ./cmd/* ./conf/* ./data/* ./domain/* ./handler/*
	@gofmt -w -l ./cmd/* ./conf/* ./data/* ./domain/* ./handler/*

stan: generate
	@revive -config revive.toml -formatter friendly ./...
	@prealloc ./...

run:
	@$(GO_BIN) run $(APP_SRC)

all: clean build