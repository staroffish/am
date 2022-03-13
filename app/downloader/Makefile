GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git rev-parse HEAD)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
NAME=am-downloader

.PHONY: errors

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=. \
 	       --go_out=paths=source_relative:. \
	       $(INTERNAL_PROTO_FILES)

.PHONY: build
# build
build:
	mkdir -p ../../bin/ && go build -mod vendor -ldflags "-X main.Version=$(VERSION)" -o ../../bin/$(NAME) ./cmd

.PHONY: generate
# generate
generate:
	go generate ./...
