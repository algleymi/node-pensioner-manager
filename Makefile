.PHONY: all lint test
all: lint test

lint:
	go fmt ./...
test:
	go test ./... -cover
