package main

import (
	"context"
	"errors"
	pb "flash/calculator/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedOperationsServer
}

func (s *server) Add(ctx context.Context, req *pb.NumberRequest) (*pb.NumberResponse, error) {
	return &pb.NumberResponse{
		Result: req.FirstNum + req.SecondNum,
	}, nil
}

func (s *server) Subtract(ctx context.Context, req *pb.NumberRequest) (*pb.NumberResponse, error) {
	return &pb.NumberResponse{
		Result: req.FirstNum - req.SecondNum,
	}, nil
}

func (s *server) Multiply(ctx context.Context, req *pb.NumberRequest) (*pb.NumberResponse, error) {
	return &pb.NumberResponse{
		Result: req.FirstNum * req.SecondNum,
	}, nil
}

func (s *server) Division(ctx context.Context, req *pb.NumberRequest) (*pb.NumberResponse, error) {
	if req.SecondNum == 0 {
		return &pb.NumberResponse{
			Result: 0,
		}, errors.New("Cannot divide by zero")
	}

	return &pb.NumberResponse{
		Result: req.FirstNum / req.SecondNum,
	}, nil
}

func (s *server) Power(ctx context.Context, req *pb.NumberRequest) (*pb.NumberResponse, error) {
	result := req.FirstNum
	for i := 1; i < int(req.SecondNum); i++ {
		result *= req.FirstNum
	}

	return &pb.NumberResponse{
		Result: result,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOperationsServer(s, &server{})

	log.Println("Starting server on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}