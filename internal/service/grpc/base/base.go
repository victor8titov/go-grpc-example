package base

import (
	"context"
	"log"

	"github.com/victor8titov/go-grpc-example/internal/service/grpc/basepb/v1"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	basepb.UnimplementedBaseServer
}

func NewServer() *server {
	return &server{}
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *basepb.HelloRequest) (*basepb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &basepb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
