package main

import (
	"fmt"
	"grpc/api_gateway/internal/handler"
	"grpc/pkg/env"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/users/getUserById", handler.GetUserByIdHandler)
	http.HandleFunc("/api/v1/users/getAllUsers", handler.GetAllUsersHandler)
	http.HandleFunc("/api/v1/users/createUser", handler.CreateUserHandler)

	env.LoadEnv("api_gateway/config/api_gateway.env")

	port := env.GetEnv("SERVER_PORT", "8080")
	fmt.Println("API Gateway is running on port " + port)
	http.ListenAndServe(":"+port, nil)
}
