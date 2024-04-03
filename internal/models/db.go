package models

import (
	"dcard-golang-project/middlewares"
	"dcard-golang-project/schemas"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DBInit() {
	USER := os.Getenv("DB_USER")
	HOST := os.Getenv("DB_HOST")
	NAME := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, HOST, NAME)

	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	}); err != nil {
		panic(&(middlewares.ServerInternalError{Message: err.Error()}))
	}

	DB.AutoMigrate(&schemas.Admin{}, &schemas.Country{}, &schemas.Gender{}, &schemas.Platform{})

	fmt.Println("db is connecting")
}
