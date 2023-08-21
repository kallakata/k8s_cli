GOCMD=go
BINARY_NAME=kubeparser

.PHONY: install build run vendor fmt tools lint

build: 
	mkdir -p bin
	$(GOCMD) build -mod vendor -o bin/$(BINARY_NAME) .

install: 
	$(GOCMD) install

run: 
	mkdir -p bin
	$(GOCMD) build -mod vendor -o bin/$(BINARY_NAME) .
	./bin/kubeparser

vendor:
	$(GOCMD) mod vendor

fmt:
	$(GOCMD) fmt ./...

tools:
	$(GOCMD) get golang.org/x/lint/golint

test:
	$(GOCMD) vet -v ./...