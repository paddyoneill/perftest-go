package main

import (
	"log"
    "context"
    "time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/paddyoneill/perftest-go/pkg/perftest"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := perftest.NewPerftestClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.GetPort(ctx, &perftest.PerftestPortRequest{})
	if err != nil {
		log.Fatalf("failed to run GetPort: %v", err)
	}

	log.Printf("Reponse from gRPC server has port %d", resp.GetValue())
}
