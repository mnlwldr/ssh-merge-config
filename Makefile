clean:
	@go clean

build: clean
	@go build .

test:
	golangci-lint run ./...
	staticcheck ./...
	go test ./...
