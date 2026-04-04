.PHONY: all build test lint fmt clean install help

BINARY := mmk-cn
DIST := dist
GO := go
MODULE := github.com/pureugong/mmk-cardnews-v2
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")

all: lint test build

build:
	@mkdir -p $(DIST)
	$(GO) build -ldflags "-X main.version=$(VERSION)" -o $(DIST)/$(BINARY) ./cmd/mmk-cn

test:
	$(GO) test -v ./...

test-coverage:
	$(GO) test -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Open coverage.html to view coverage report"

lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		$(GO) vet ./...; \
		gofmt -l .; \
	fi

fmt:
	gofmt -l -w .

clean:
	rm -rf $(DIST)/ coverage.out coverage.html

install: build
	@cp $(DIST)/$(BINARY) $$($(GO) env GOPATH)/bin/
	@echo "Installed $(BINARY) to $$($(GO) env GOPATH)/bin/"

help:
	@echo "Targets:"
	@echo "  build          Build CLI binary to dist/"
	@echo "  test           Run tests"
	@echo "  test-coverage  Run tests with coverage report"
	@echo "  lint           Run linter"
	@echo "  fmt            Format Go files"
	@echo "  clean          Remove build artifacts"
	@echo "  install        Install binary to GOPATH/bin"
