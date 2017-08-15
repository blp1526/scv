VERSION = $(shell ./shellscripts/version.sh)
LDFLAGS = -ldflags "-X github.com/blp1526/scv.Version="$(VERSION)

.PHONY: all
all: build

.PHONY: test
test:
	@echo "\n######## TEST ########\n"
	go test -v --cover ./...

.PHONY: tmp
tmp: test
	@echo "\n######## TMP ########\n"
	mkdir -p tmp

.PHONY: build
build: tmp
	@echo "\n######## BUILD ########\n"
	go build $(LDFLAGS) -o tmp/scv cmd/scv/scv.go

.PHONY: install
install: build
	@echo "\n######## INSTALL ########\n"
	mv tmp/scv ${GOPATH}/bin/scv

.PHONY: zip
zip: build
	@echo "\n######## ZIP ########\n"
	./shellscripts/zip.sh

.PHONY: clean
clean:
	rm -rf tmp
	rm -rf ${GOPATH}/bin/scv
