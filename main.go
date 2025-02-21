package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	catcatcatpb "github.com/holin20/catcatcat/proto/catcatcat"
)

type server struct {
	catcatcatpb.UnimplementedCatcatcatServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *catcatcatpb.ListCatsRequest) (*catcatcatpb.ListCatsResponse, error) {
	return &catcatcatpb.ListCatsResponse{}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	catcatcatpb.RegisterCatcatcatServer(s, &server{})
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatal(s.Serve(lis))
}
