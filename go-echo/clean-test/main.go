package main

import (
	"belajar-go-echo/config"
	"belajar-go-echo/controller"
	"belajar-go-echo/middleware"
	"belajar-go-echo/repository"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}
	if err = config.MigrateDB(db); err != nil {
		panic(err)
	}

	iUserRepo := repository.NewUserRepo(db) // --> layer data   <-- unit test

	loginController := controller.NewLoginController(iUserRepo) // --> layer controller
	userController := controller.NewUserController(iUserRepo)   // --> layer controller

	app := echo.New()

	app.POST("/login", loginController.Login)

	app.GET("/users", userController.GetAllData, middleware.ValidateJwt())
	app.GET("/user", userController.GetSingleData, middleware.ValidateJwt())
	app.POST("/users", userController.Create, middleware.ValidateJwt())
	app.Start(":8080")
}
