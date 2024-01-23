package main

import (
	"fmt"
	"log"
	"net"

	"github.com/victor8titov/go-grpc-example/internal/service/grpc/base"
	"github.com/victor8titov/go-grpc-example/internal/service/grpc/basepb/v1"
	"google.golang.org/grpc"
)

func main() {
	port := 50051

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	server := base.NewServer()
	basepb.RegisterBaseServer(s, server)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
