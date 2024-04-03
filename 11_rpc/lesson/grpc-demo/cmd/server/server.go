package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpc-demo/gen/proto/calc"
)

var port = flag.Int("port", 50051, "The server port")

type server struct {
	pb.UnimplementedCalculatorServer
}

// Multiply implements calculator.CalculatorServer
func (s *server) Multiply(ctx context.Context, in *pb.MultiplyRequest) (*pb.MultiplyResponse, error) {
	log.Printf("Received arguments: %v, %v", in.ArgumentA, in.ArgumentB)
	return &pb.MultiplyResponse{Result: in.ArgumentA * in.ArgumentB}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
