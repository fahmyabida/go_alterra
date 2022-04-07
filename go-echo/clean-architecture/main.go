package main

import (
	"belajar-go-echo/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	// db, err := config.ConnectDB()
	// if err != nil {
	// 	panic(err)
	// }

	// err = config.MigrateDB(db)
	// if err != nil {
	// 	panic(err)
	// }

	app := echo.New()

	// iUserRepo := repository.NewUserRepo(db)
	// userController := controller.NewUserController(iUserRepo)
	loginController := controller.NewLoginController()

	app.POST("/login", loginController.Login)

	// app.GET("/users", userController.GetAllData)
	// app.POST("/users", userController.Create)
	app.Start(":8080")
}
