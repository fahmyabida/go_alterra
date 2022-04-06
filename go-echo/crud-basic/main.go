package main

import (
	"alterra/go_alterra/crud-basic/config"
	"alterra/go_alterra/crud-basic/routes"

	"github.com/labstack/echo"
)

func init() {
	config.InitDB()
	config.InitialMigration()
}

func main() {
	e := echo.New()

	routes.Route(e)

	e.Logger.Fatal(e.Start(":8000"))
}
