package scrm

import (
	"context"

	grpc "google.golang.org/grpc"
)

type Server struct{}

// SayHello implements GreeterServer
func (s Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "Hello " + in.Name + ": " + in.Sex}, nil
}

const port = ":30031"

func NewClient(address string) (ScrmServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := NewScrmServiceClient(conn)
	return c, nil
}
