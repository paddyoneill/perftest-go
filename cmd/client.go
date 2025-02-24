package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/paddyoneill/perftest-go/pkg/perftest"
)

var serverAddr string

func main() {
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := perftest.NewPerftestClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	portResp, err := client.GetPort(ctx, &perftest.PortRequest{})
	if err != nil {
		log.Fatalf("failed to run GetPort: %v", err)
	}

	port := portResp.GetValue()
	log.Printf("Reponse from gRPC server has port %d", port)

	perftestResp, err := client.StartPerftest(ctx, &perftest.PerftestRequest{Port: port})
	if err != nil {
		log.Fatal("failed to start perftest")
	}

	log.Printf("Response from server: %s", perftestResp.GetMessage())
}

func init() {
	flag.StringVar(&serverAddr, "server", "localhost:8080", "address of server")
	flag.Parse()
}
