package service

// import "grpc/services/user_service/internal/repository"

// func GetUserById(id int) (map[string]interface{}, error) {
// 	user, err := repository.GetUserById(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return map[string]interface{}{
// 		"id":       user.ID,
// 		"username": user.Username,
// 		"name":     user.Name,
// 	}, nil
// }

// func GetAllUsers() ([]map[string]interface{}, error) {
// 	users, err := repository.GetAllUsers()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var usersData []map[string]interface{}
// 	for _, user := range users {
// 		usersData = append(usersData, map[string]interface{}{
// 			"id":       user.ID,
// 			"username": user.Username,
// 			"name":     user.Name,
// 		})
// 	}

// 	return usersData, nil
// }
