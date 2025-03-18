package repository

import (
	"grpc/pkg/env"
	"grpc/pkg/gorm"

	orm "gorm.io/gorm"
)

var db *orm.DB

func InitDB() {
	if db != nil {
		return
	}

	host := env.GetEnv("USER_DB_HOST", "localhost")
	port := env.GetEnv("USER_DB_PORT", "5432")
	user := env.GetEnv("USER_DB_USER", "postgres")
	password := env.GetEnv("USER_DB_PASSWORD", "12345678")
	dbname := env.GetEnv("USER_DB_NAME", "user_db")

	dsn := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"
	db = gorm.NewGormDB(dsn)
}