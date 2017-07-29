.PHONY: all zip build clean test tmp

all: build

zip: build
	@./shellscripts/zip.sh

build: tmp
	@mkdir -p tmp
	@go build -o tmp/scv -ldflags "-X github.com/blp1526/scv.version="$(shell ./shellscripts/version.sh) cmd/scv/scv.go

clean:
	@rm -rf tmp

test:
	@go test -v

tmp:
	@mkdir -p tmp
