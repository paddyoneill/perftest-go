package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/paddyoneill/perftest-go/pkg/perftest"
)

type PerftestServer struct {
	perftest.UnimplementedPerftestServer
	WorkerQueue <-chan string
}

func (server *PerftestServer) GetPort(context.Context, *perftest.PortRequest) (*perftest.PortResponse, error) {
	port, err := getAvailablePort()

	if err != nil {
		return nil, err
	}

	return &perftest.PortResponse{Value: port}, nil
}

func (server *PerftestServer) StartPerftest(ctx context.Context, req *perftest.PerftestRequest) (*perftest.PerftestResponse, error) {
	port := req.GetPort()
	log.Printf("Starting perftest on port %d", port)
	return &perftest.PerftestResponse{Message: fmt.Sprintf("starting perftest on port %d", port)}, nil
}

func (server *PerftestServer) GetNic(ctx context.Context, req *perftest.NicRequest) (*perftest.NicResponse, error) {
	select {
	case nic := <-server.WorkerQueue:
		return &perftest.NicResponse{Nic: nic}, nil

	default:
		return &perftest.NicResponse{}, fmt.Errorf("No available worker in queue")
	}
}

func main() {
	nics := []string{"bnxt_re0", "bnxt_re1", "bnxt_re2", "bnxt_re3"}
	perftestServer := &PerftestServer{
		WorkerQueue: initWorkerQueue(nics),
	}

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to start server on port 8080: %v", err)
	}

	grpcServer := grpc.NewServer()
	perftest.RegisterPerftestServer(grpcServer, perftestServer)
	log.Printf("gRPC service listening at %s", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getAvailablePort() (int32, error) {
	addr, err := net.ResolveTCPAddr("tcp", ":0")
	if err != nil {
		return -1, err
	}
	lis, err := net.ListenTCP("tcp", addr)
	defer lis.Close()
	if err != nil {
		return -1, err
	}
	return int32(lis.Addr().(*net.TCPAddr).Port), nil
}
