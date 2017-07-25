.PHONY: all zip build clean test

all: build

zip: build
	@./zip

build:
	@go build -ldflags "-X github.com/blp1526/scv/cmd.version="$(shell ./version)

clean:
	@go clean
	@rm -rf archives

test:
	@go test -v ./...
