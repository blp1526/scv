.PHONY: all zip build clean test tmp

VERSION = $(shell ./shellscripts/version.sh)
LDFLAGS = -ldflags "-X github.com/blp1526/scv.version="$(VERSION)

all: build

zip: build
	@./shellscripts/zip.sh

build: tmp
	@go build $(LDFLAGS) -o tmp/scv cmd/scv/scv.go

clean:
	@rm -rf tmp

test:
	@go test -v

tmp:
	@mkdir -p tmp
