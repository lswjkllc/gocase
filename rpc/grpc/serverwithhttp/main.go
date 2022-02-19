// Package main implements a server for Greeter service.
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "gocase/rpc/grpc/protos/helloworldwithhttp"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	host     = "0.0.0.0"
	grpcPort = 50051
	httpPort = 8000
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Grpc With Http Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// 服务 host:port
	grpcAddr := fmt.Sprintf("%s:%d", host, grpcPort)
	httpAddr := fmt.Sprintf("%s:%d", host, httpPort)

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
	go func() {
		log.Fatalln(s.Serve(lis))
	}()
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		grpcAddr,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = pb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    httpAddr,
		Handler: gwmux,
	}

	log.Printf("Serving http on %s\n", httpAddr)
	log.Fatalln(gwServer.ListenAndServe())
}
