.PHONY: all clean test

all:
	@go build -ldflags "-X github.com/blp1526/scv/cmd.version="$(shell ./version)

clean:
	@go clean

test:
	@go test ./...
