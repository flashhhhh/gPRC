package repository

import (
	"errors"
	"grpc/pkg/hash"
)

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"unique, not null"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
}

func GetUserById(id int) (User, error) {
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		return User{}, errors.New("User not found")
	}

	return user, nil
}

func GetAllUsers() ([]User, error) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, errors.New("failed to get all users")
	}

	return users, nil
}

func CreateUser(username, password, name string) (User, error) {
	hashedPassword := hash.HashString(password)

	user := User{
		Username: username,
		Password: hashedPassword,
		Name:     name,
	}
	result := db.Create(&user)
	if result.Error != nil {
		return User{}, errors.New("failed to create user")
	}

	return user, nil
}