package main

import (
	"context"
	"log"
	"os"
	"time"

	"eglass.com/microservices/scrm"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:30031"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := scrm.NewScrmServiceClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancel()
	r, err := c.SayHello(ctx, &scrm.HelloRequest{Name: name, Sex: "jdjd"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
