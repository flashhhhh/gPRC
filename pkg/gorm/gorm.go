package gorm

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func NewGormDB(dsn string) (*gorm.DB) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}