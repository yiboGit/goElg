# 服务器编译
DEPENDENCIES := github.com/gorilla/websocket \
	github.com/go-redis/redis \
    github.com/go-sql-driver/mysql \
    upper.io/db.v3
COMPILER := oa.epeijing.cn:5000/goc:alpine
HOME := "/go/src/eglass.com"
all: build-ws

clean: rm -rf bin/*

build-scheduler:
	docker run --rm -v $(CURDIR):$(HOME) -w $(HOME) $(COMPILER) go build -o bin/scheduler eglass.com/scheduler
	docker build -t scheduler -f scheduler-dockerfile .
run-batch:
	docker run --rm -v $(CURDIR):$(HOME) -w $(HOME) goc go build -o batchjob eglass.com/batch
	prod=true ./batchjob
build-ws:
	docker run --rm -v $(CURDIR):$(HOME) -w $(HOME) $(COMPILER) go build -o bin/ws eglass.com/ws
	docker build -t ws -f ws-dockerfile .
build-card-push:
	docker run --rm -v $(CURDIR):$(HOME) -w $(HOME) $(COMPILER) go build -o bin/card-push eglass.com/card-push
	docker build -t card-push -f push-dockerfile .
build-moniter:
	docker run --rm -v $(CURDIR):$(HOME) -w $(HOME) $(COMPILER) go build -o bin/moniter eglass.com/moniter
	docker build -t moniter -f moniter-dockerfile .

deps:
	go get $(DEPENDENCIES)
	docker pull oa.epeijing.cn:5000/golang-alpine
	docker tag oa.epeijing.cn:5000/golang-alpine golang-alpine
