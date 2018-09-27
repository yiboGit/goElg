DEPENDENCIES := github.com/gorilla/websocket \
	github.com/go-redis/redis
all: build silent-test
GOPATH := ~/go
build:
	go build -o bin/im-test *.go
build-ubuntu:
	cd $(GOPATH) && docker run --rm -v $(GOPATH):/go -w /go golang go build -o src/eglass.com/im/bin/im eglass.com/im
test:
	go test -v $(PACKAGES)

silent-test:
	go test $(PACKAGES)

deps:
	go get $(DEPENDENCIES)