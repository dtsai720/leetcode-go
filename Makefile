.PHONY: test
test:
	@go test -race -v ./... 2>&1

.PHONY: lint
lint:
	@golangci-lint run ./...
