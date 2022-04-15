package config

import (
	"belajar-go-echo/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Category struct {
	Id   string
	Name string
}

func (Category) TableName() string {
	return "category"
}

func ConnectDB() (*gorm.DB, error) {
	// return gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	connectionString := fmt.Sprintf("root:Admin123@tcp(localhost:3306)/clean")

	var err error
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	categoryNew := Category{
		Id:   "idBAru",
		Name: "Apaja deh",
	}
	db.Create(&categoryNew)

	obj := []Category{}
	if err = db.Find(&obj).Debug().Error; err != nil {
		panic(err)
	}
	for i, row := range obj {
		fmt.Println("dapet datanya", i, row)
	}
	
	return db, err
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
	)
}
