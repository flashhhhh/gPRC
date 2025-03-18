package handler

import (
	"encoding/json"
	"grpc/api_gateway/internal/service"
	"net/http"
	"strconv"
)

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, errAtoi := strconv.Atoi(idStr)
	if errAtoi != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id must be an integer"))
		return
	}

	user := service.GetUserById(id)
	userJSON, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error converting user to JSON"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := service.GetAllUsers()

	usersJSON, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error converting users to JSON"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(usersJSON)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	name := r.URL.Query().Get("name")

	user := service.CreateUser(username, password, name)
	userJSON, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error converting user to JSON"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}