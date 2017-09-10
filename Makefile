VERSION = $(shell git describe --dirty --tags)
LDFLAGS = -ldflags "-X github.com/blp1526/scv.Version="$(VERSION)

.PHONY: all
all: build

.PHONY: test
test:
	go test -v --cover ./...

.PHONY: tmp
tmp: test
	mkdir -p tmp

.PHONY: build
build: tmp
	go build $(LDFLAGS) -o tmp/scv cmd/scv/scv.go

.PHONY: install
install: build
	mv tmp/scv ${GOPATH}/bin/scv

.PHONY: compress
compress: build
	./shellscripts/compress.sh

.PHONY: clean
clean:
	rm -rf tmp
	rm -rf ${GOPATH}/bin/scv
