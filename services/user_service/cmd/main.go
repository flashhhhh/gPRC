package main

import (
	"fmt"
	"grpc/pkg/env"
	"grpc/services/user_service/internal/repository"
	"grpc/services/user_service/internal/service"
	pb "grpc/services/user_service/pb"
	"net"

	"google.golang.org/grpc"
)

func main() {
	env.LoadEnv("grpc/services/user_service/config/user.env")
	repository.InitDB()

	lis, err := net.Listen("tcp", env.GetEnv("USER_PORT", ":50051"))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &service.Server{})

	fmt.Println("Starting server at " + env.GetEnv("USER_PORT", ":50051"))
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}