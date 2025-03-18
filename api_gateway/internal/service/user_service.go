package service

import (
	"context"
	pb "grpc/api_gateway/pb"
	"grpc/pkg/env"

	"google.golang.org/grpc"
)

var userClient pb.UserServiceClient

func InitUserService() {
	if userClient != nil {
		return
	}

	userServicePort := env.GetEnv("USER_SERVICE_PORT", "50051")

	conn, err := grpc.Dial("localhost:" + userServicePort, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	userClient = pb.NewUserServiceClient(conn)
}

func GetUserById(id int) map[string]interface{} {
	InitUserService()
	ctx := context.Background()

	req := &pb.IdRequest{Id: int32(id)}
	res, err := userClient.GetUserById(ctx, req)
	if err != nil {
		panic(err)
	}

	return map[string]interface{}{
		"id":       res.Id,
		"username": res.Username,
		"name":     res.Name,
	}
}

func GetAllUsers() map[string]interface{} {
	InitUserService()
	ctx := context.Background()

	req := &pb.EmptyRequest{}
	res, err := userClient.GetAllUsers(ctx, req)
	if err != nil {
		panic(err)
	}

	users := make([]map[string]interface{}, 0)
	for _, user := range res.Users {
		users = append(users, map[string]interface{}{
			"id":       user.Id,
			"username": user.Username,
			"name":     user.Name,
		})
	}

	return map[string]interface{}{
		"users": users,
	}
}
