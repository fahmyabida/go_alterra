package routes

import (
	"alterra/go_alterra/praktikum_orm/controller"

	"github.com/labstack/echo"
)

func Route(e *echo.Echo) {
	// just check the echo is running well & library just fine
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, "pong")
	})

	// Route / to handler function
	e.GET("/users", controller.GetAllUserController)
	e.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.PUT("/users/:id", controller.UpdateUserController)
}
