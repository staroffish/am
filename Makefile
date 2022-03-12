GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
API_PROTO_FILES=$(shell find api -name *.proto)
APP_CONFIG_DIRS=$(shell find app/*/internal -maxdepth 1 -type d -name conf)
APP_DIRS=$(shell find app -maxdepth 1 -type d ! -name app)
APP_DIRS_CONFIG=$(APP_DIRS:%=%_config)
APP_DIRS_BUILD=$(APP_DIRS:%=%_build)
APP_DIRS_GENERATE=$(APP_DIRS:%=%_generate)

.PHONY: $(APP_DIRS)

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@v0.6.1

.PHONY: errors
# generate errors code
errors:
	protoc --proto_path=. \
               --proto_path=./third_party \
               --go_out=paths=source_relative:. \
               --go-errors_out=paths=source_relative:. \
               $(API_PROTO_FILES)

.PHONY: config $(APP_DIRS_CONFIG)
# generate internal proto
config: $(APP_DIRS_CONFIG)

$(APP_DIRS_CONFIG): 
	make -C $(@:%_config=%) config

.PHONY: api
# generate api proto
api:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
 	       --openapi_out==paths=source_relative:. \
	       $(API_PROTO_FILES)

.PHONY: build $(APP_DIRS_BUILD)
# build
build: $(APP_DIRS_BUILD)

$(APP_DIRS_BUILD):
	 make -C $(@:%_build=%) build

.PHONY: generate $(APP_DIRS_GENERATE)
# generate
generate: $(APP_DIRS_GENERATE)

$(APP_DIRS_GENERATE): 
	make -C $(@:%_generate=%) generate

.PHONY: all
# generate all
all:
	make api;
	make errors;
	make config;
	make generate;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

