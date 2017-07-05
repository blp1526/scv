.PHONY: all clean test

all:
	@go build

clean:
	@go clean

test:
	@go test ./...
