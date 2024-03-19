package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {

	dsn := "user:@tcp(localhost:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		// panic(&(middlewares.ServerInternalError{Message: err.Error()}))
		fmt.Println(map[string]string{
			"Message": err.Error(),
		})
	} else {
		fmt.Println("db is connecting")
	}
}
