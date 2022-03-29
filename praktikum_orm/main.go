package main

import "github.com/labstack/echo"

// init adalah function yg dijalankan sebelum func main()
func init() {
	InitDB()
	InitialMigration()
}

func main() {
	// create a new echo instance
	e := echo.New()

	// just check the echo is running well & library just fine
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, "pong")
	})

	// Route / to handler function
	e.GET("/users", GetAllUserController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
