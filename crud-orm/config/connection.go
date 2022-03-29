package config

import (
	"alterra/go_alterra/crud-orm/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase() {
	fmt.Println("connet to db")
	dsn := fmt.Sprintf("root:%v@tcp(127.0.0.1:3306)/belajar_orm", model.PASSWORD_DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var user model.Users
	db.AutoMigrate(&user, &model.Product{})
}
