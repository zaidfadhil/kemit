PREFIX ?= /usr/local
VERSION ?= $(shell git describe --tags --abbrev=0)

.PHONY: go-install
go-install:
	@go get -v -t -d ./...

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: test
test:
	@go test -count=1 -v ./...

.PHONY: lint
lint: 
	@golangci-lint run ./...

.PHONY: build
build:
	@go build -ldflags "-X main.version=$(VERSION)" -o bin/kemit main.go

.PHONY: install
install: build
	@cp bin/kemit $(PREFIX)/bin/kemit
	@chmod +x $(PREFIX)/bin/kemit

setup-githook:
	git config core.hooksPath .githooks
