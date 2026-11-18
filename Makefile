BINARY := bin/server$(shell go env GOEXE)

.PHONY: build run build-run test build-test

## build: compile the binary
build:
	go build -o $(BINARY) ./cmd

## run: run the server from source
run:
	go run ./cmd

## build-run: compile and start the compiled binary
build-run: build
	./$(BINARY)

## test: run all unit tests no caching
test:
	go test ./... -v -count=1

## build-test: compile and run all unit tests
build-test: build test
