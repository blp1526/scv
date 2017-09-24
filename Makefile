LOCAL_PACKAGES = $(shell cd $(CURDIR) && go list ./... | grep -v '/vendor/')
VERSION = $(shell git describe --dirty --tags)
LDFLAGS = -ldflags "-X github.com/blp1526/scv.Version="$(VERSION)

.PHONY: all
all: build

.PHONY: vet
vet:
	go vet $(LOCAL_PACKAGES)

.PHONY: test
test: vet
	go test -v --cover $(LOCAL_PACKAGES)

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
