.PHONY: all

all: gofmt test examples

gofmt:
	@echo "formating..."
	@gofmt -s -w .

test:
	@echo "run test..."
	@go mod tidy
	@go clean -testcache
	@go test ./... -cover

examples:
	@echo "generate examples..."
	@cd example && go run .
