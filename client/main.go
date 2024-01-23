package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/victor8titov/go-grpc-example/internal/service/grpc/basepb/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	grpcClient := basepb.NewBaseClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := grpcClient.SayHello(ctx, &basepb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not say hello: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
