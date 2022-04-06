package controller

import (
	"alterra/go_alterra/praktikum_orm/lib/database"
	"alterra/go_alterra/praktikum_orm/models"
	"net/http"

	"github.com/labstack/echo"
)

// get all users
func GetAllUserController(c echo.Context) error {
	listUser, err := database.GetAllUser()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   listUser,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	user, err := database.GetSingleUserById(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get users",
		"user":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	err := database.CreateNewUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	err := database.DeleteUserById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user with id '" + id + "'",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	err := database.UpdateUserById(c.Param("id"), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}
