PACKAGES = main
DEPENDENCIES := github.com/gorilla/websocket \
	github.com/go-redis/redis
all: build silent-test

build-test:
	go build -o bin/scheduler-test *.go
build-prod:
	cd ~/go && docker run --rm -v "$PWD":/go -w /go golang go build -o eglass.com/bin/scheduler eglass.com/scheduler
run:
	go build -o bin/scheduler-test *.go && ./bin/scheduler-test
test:
	go test -v .
docker: 
	docker build -t eglass/scheduler .
silent-test:
	go test $(PACKAGES)

format:
	go fmt $(PACKAGES)
deps:
	go get $(DEPENDENCIES)