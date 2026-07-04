.PHONY: build run clean frontend install dev docker migrate

NAME=am
VERSION=$(shell git describe --tags --always 2>/dev/null || echo "dev")
LDFLAGS=-ldflags "-s -w -X main.Version=$(VERSION)"
GO_BUILD=go build $(LDFLAGS) -o bin/$(NAME) ./cmd/server

SHELL := /bin/bash

export PATH := /usr/local/bin/go/bin:$(HOME)/go/bin:$(PATH)

define NVM_LOAD
export NVM_DIR="$(HOME)/.nvm"; \
[ -s "$$NVM_DIR/nvm.sh" ] && . "$$NVM_DIR/nvm.sh" || true
endef

install:
	@$(NVM_LOAD) && echo "Installing frontend dependencies..." && cd web && npm install && echo "Done."

frontend:
	@$(NVM_LOAD) && echo "Building frontend..." && cd web && npm run build && echo "Frontend done."

build: frontend
	@echo "Building $(NAME)..."
	$(GO_BUILD)
	@echo "Binary: bin/$(NAME)"

build-server:
	@echo "Building server only..."
	$(GO_BUILD)
	@echo "Binary: bin/$(NAME)"

migrate:
	@echo "Building migration tool..."
	go build $(LDFLAGS) -o bin/migrate ./cmd/migrate
	@echo "Binary: bin/migrate"

run:
	go run ./cmd/server

dev:
	go run ./cmd/server

clean:
	rm -rf bin/
	rm -f am.log

docker:
	docker build -t am:$(VERSION) .

.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  install       - Install frontend npm dependencies"
	@echo "  frontend      - Build Vue frontend only"
	@echo "  build         - Build frontend + Go binary"
	@echo "  build-server  - Build Go server only (skip frontend)"
	@echo "  migrate       - Build etcd migration tool"
	@echo "  run           - Run server (dev mode)"
	@echo "  clean         - Remove build artifacts"
	@echo "  docker        - Build Docker image"
