all: get lint test

lint:
	go fmt ./...
test:
	go test ./... -cover
deps:
	go list -m -u all
