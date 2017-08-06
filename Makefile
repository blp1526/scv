.PHONY: test build zip clean

VERSION = $(shell ./shellscripts/version.sh)
LDFLAGS = -ldflags "-X github.com/blp1526/scv.Version="$(VERSION)

test:
	@go test -v

build: test
	@mkdir -p tmp
	@go build $(LDFLAGS) -o tmp/scv cmd/scv/scv.go

zip: build
	@./shellscripts/zip.sh

clean:
	@rm -rf tmp
