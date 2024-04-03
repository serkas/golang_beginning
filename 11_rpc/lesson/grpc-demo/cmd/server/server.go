package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	"grpc-demo/logger"

	"google.golang.org/grpc"
	pb "grpc-demo/gen/proto/calc"
)

var port = flag.Int("port", 50051, "The server port")

type server struct {
	pb.UnimplementedCalculatorServer

	log logger.Logger
}

// Multiply implements calculator.CalculatorServer
func (s *server) Multiply(ctx context.Context, in *pb.MultiplyRequest) (*pb.MultiplyResponse, error) {
	s.log.Info("Received arguments: %v, %v", in.ArgumentA, in.ArgumentB)
	return &pb.MultiplyResponse{Result: in.ArgumentA * in.ArgumentB}, nil
}

func main() {
	flag.Parse()

	log := logger.NewZeroLogLogger()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Error("failed to listen: %v", err)
		return
	}

	srv := &server{
		log: log,
	}

	grpcSrv := grpc.NewServer()
	pb.RegisterCalculatorServer(grpcSrv, srv)
	log.Info("server listening at %v", lis.Addr())

	if err := grpcSrv.Serve(lis); err != nil {
		log.Error("failed to serve: %v", err)
	}
}
