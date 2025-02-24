package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/paddyoneill/perftest-go/pkg/perftest"
)

type server struct {
	perftest.UnimplementedPerftestServer
}

func (server *server) GetPort(context.Context, *perftest.PerftestPortRequest) (*perftest.PerftestPortResponse, error) {
	return &perftest.PerftestPortResponse{Value: 12345}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to start server on port 8080: %v", err)
	}

	grpcServer := grpc.NewServer()
	perftest.RegisterPerftestServer(grpcServer, &server{})
	log.Printf("gRPC service listening at %s", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
