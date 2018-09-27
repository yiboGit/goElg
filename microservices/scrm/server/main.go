package main

import (
	"eglass.com/microservices"
	"eglass.com/microservices/scrm"
	"google.golang.org/grpc"
)

type ScrmServer struct{}

func (s ScrmServer) Port() string {
	return ":30031"
}
func (s ScrmServer) Name() string {
	return "scrm"
}

func (s ScrmServer) Start(gs *grpc.Server) {
	scrm.RegisterScrmServiceServer(gs, scrm.Server{})
}

func main() {
	microservices.StartGrpcServer(ScrmServer{})
}
