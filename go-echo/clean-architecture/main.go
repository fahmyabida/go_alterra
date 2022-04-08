package main

import (
	"belajar-go-echo/config"
	"belajar-go-echo/controller"
	"belajar-go-echo/repository"
	"belajar-go-echo/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	iUserRepo := repository.NewUserRepo(db)
	loginController := controller.NewLoginController(iUserRepo)
	userController := controller.NewUserController(iUserRepo)

	app := echo.New()

	app.POST("/login", loginController.Login)

	app.GET("/users", userController.GetAllData, utils.ValidateJwt())
	app.POST("/users", userController.Create, utils.ValidateJwt())
	app.Start(":8080")
}
