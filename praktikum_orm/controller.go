package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// get all users
func GetAllUserController(c echo.Context) error {
	var users []User

	// SELECT * FROM USERS;
	err := DB.Find(&users).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	var user User

	stringId := c.Param("id")
	// SELECT * FROM users WHERE id = ?;
	err := DB.First(&user, "id = ?", stringId).Debug().Error
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
	user := User{}
	c.Bind(&user)

	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	stringId := c.Param("id")
	err := DB.Delete(&User{}, "id = ?", stringId).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user with id '" + stringId + "'",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)
	stringId := c.Param("id")
	err := DB.Model(&user).Where("id = ?", stringId).Updates(user).Debug().Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}
