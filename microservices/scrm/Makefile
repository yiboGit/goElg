PACKAGES := \
	github.com/eliasson/foo \
	github.com/eliasson/bar
DEPENDENCIES := github.com/eliasson/acme

all: build silent-test

build:
	go build -o bin/scrm-server server/main.go

gen-proto:
	cd .. && protoc -I scrm scrm/scrm.proto --go_out=plugins=grpc:scrm

test:
	go test -v $(PACKAGES)

silent-test:
	go test $(PACKAGES)

format:
	go fmt $(PACKAGES)

deps:
	go get $(DEPENDENCIES)