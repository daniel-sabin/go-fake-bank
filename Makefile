.ONESHELL:
.PHONY: help

BINARY_NAME=demobank-server
OS := $(shell uname)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n" , $$1, $$2}'

clean: ## Go clean
	go clean

test: ## Run test
	go test ./...

test_coverage: ## Run test with coverage
	go test ./... -coverprofile=coverage.out

run:
	./${BINARY_NAME}

build: ## Build application for your os
	GOARCH=amd64 GOOS=${OS,,} go build -o ${BINARY_NAME} main.go

br: clean build run ## Build and run application