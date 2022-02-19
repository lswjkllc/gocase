// Package main implements a server for Greeter service.
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "gocase/rpc/grpc/protos/helloworld"

	"google.golang.org/grpc"
)

const (
	host     = "0.0.0.0"
	grpcPort = 50051
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Grpc Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// 服务 host:port
	grpcAddr := fmt.Sprintf("%s:%d", host, grpcPort)

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterGreeterServer(s, &server{})
	// Serve gRPC server
	log.Printf("Serving gRPC on %s\n", grpcAddr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
