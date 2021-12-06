all: update lint test

update:
	go get ./...
	go mod tidy

lint:
	golangci-lint run

test: update
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out
