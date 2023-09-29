
gofmt:
	gofumpt -l -w .

lint:
	golangci-lint --config ./ci/.golangci-lint.yml run

.PHONY: test
test:
	go test ./...
