package service

import (
	"context"
	"grpc/services/user_service/internal/repository"
	pb "grpc/services/user_service/pb"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) GetUserById(ctx context.Context, in *pb.IdRequest) (*pb.UserResponse, error) {
	id := int(in.Id)
	user, err := repository.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:       int32(user.ID),
		Username: user.Username,
		Name:     user.Name,
	}, nil
}

func (s *Server) GetAllUsers(ctx context.Context, in *pb.EmptyRequest) (*pb.UsersResponse, error) {
	users, err := repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var usersArray []*pb.UserResponse
	for _, user := range users {
		usersArray = append(usersArray, &pb.UserResponse{
			Id:       int32(user.ID),
			Username: user.Username,
			Name:     user.Name,
		})
	}

	return &pb.UsersResponse{
		Users: usersArray,
	}, nil
}

func (s *Server) CreateUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	username := in.Username
	password := in.Password
	name := in.Name

	user, err := repository.CreateUser(username, password, name)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		Id:       int32(user.ID),
		Username: user.Username,
		Name:     user.Name,
	}, nil
}