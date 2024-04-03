package main

import (
	"context"
	"flag"
	"time"

	"grpc-demo/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-demo/gen/proto/calc"
)

const (
	defaultName = "world"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func main() {
	flag.Parse()

	log := logger.NewZeroLogLogger()

	log.Info("Client starting...")

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("did not connect: %v", err)
		return
	}
	defer conn.Close()
	c := pb.NewCalculatorClient(conn)

	log.Info("Client connected")

	req := &pb.MultiplyRequest{
		ArgumentA: 12,
		ArgumentB: 16,
	}

	log.Info("Send request: %v", req)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Multiply(ctx, req)
	if err != nil {
		log.Error("could not get response: %v", err)
		return
	}

	log.Info("Multiplication result: %v", r.GetResult())
}
