package config

import (
	"belajar-go-echo/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	// return gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	connectionString := fmt.Sprintf("root:Admin123@tcp(192.168.0.113:3306)/clean")

	var err error
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, err
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
	)
}
