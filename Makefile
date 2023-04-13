APP_NAME	:= Fiber Starter
PKG_NAME    := github.com/Redish101/fiber-starter
BIN_NAME	:= ./bin/fiber-starter
VERSION     ?= $(shell git describe --tags --abbrev=0)
COMMIT_HASH := $(shell git rev-parse --short HEAD)
DEV_VERSION := dev-$(COMMIT_HASH)
ARGS        ?= server

all: install build

install:
	go mod tidy

build: 
	@echo "Building $(APP_NAME) $(VERSION)..."
	go build \
    	-ldflags "-s -w -X $(PKG_NAME)/internal/config.Version=$(VERSION) \
        -X $(PKG_NAME)/internal/config.CommitHash=$(COMMIT_HASH)" \
        -o $(BIN_NAME) \
    	$(PKG_NAME)


run: all
	$(BIN_NAME) $(ARGS)

build-debug:
	@echo "Building $(APP_NAME) $(VERSION) for debugging..."
	@go build \
		-ldflags "-X $(PKG_NAME)/internal/config.Version=$(VERSION) \
		  -X $(PKG_NAME)/internal/config.CommitHash=$(COMMIT_HASH)" \
		-gcflags "all=-N -l" \
		-o $(BIN_NAME) \
		$(PKG_NAME)

dev: build-debug
	$(BIN_NAME) $(ARGS)
