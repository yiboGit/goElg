## golang dev
### go get 
```go
go get -u -v package
```
如果被墙了，比如 grpc，香港服务器 git clone下来，放到 $GOPATH/src, 运行 go install $package

### Makefile 使用
把常用操作放到 Makefile, 类似npm run操作，可以多行shell脚本, example 参考microservices/scrm