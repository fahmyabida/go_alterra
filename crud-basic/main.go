package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	sId := c.Param("id")
	if sId == "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "id not found",
			"user":     nil,
		})
	}
	id, _ := strconv.Atoi(sId)
	var user User
	for _, rowUser := range users {
		if rowUser.Id == id {
			user = rowUser
			break
		}
	}
	if user.Id == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": fmt.Sprintf("failed get user with id '%v' is not found", id),
			"user":     nil,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": fmt.Sprintf("success get user with id %v", id),
		"user":     user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	sId := c.Param("id")
	if sId == "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "id not found",
			"user":     nil,
		})
	}
	id, _ := strconv.Atoi(sId)
	var user User
	for i, rowUser := range users {
		if rowUser.Id == id {
			user = rowUser
			users = append(users[:i], users[i+1:]...)
			break
		}
	}
	if user.Id == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": fmt.Sprintf("failed delete user with id '%v' is not found", id),
			"user":     nil,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": fmt.Sprintf("success get user with id %v", id),
		"user":     user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	var user User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": err.Error(),
			"user":     nil,
		})
	}
	sId := c.Param("id")
	if sId == "" {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "id not found",
			"user":     nil,
		})
	}
	id, _ := strconv.Atoi(sId)
	isFound := false
	for i, rowUser := range users {
		if rowUser.Id == id {
			user.Id = id
			users[i] = user
			user.Id = rowUser.Id
			isFound = true
			break
		}
	}
	if !isFound {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": fmt.Sprintf("failed update user with id '%v' is not found", id),
			"user":     nil,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": fmt.Sprintf("success get user with id %v", id),
		"user":     user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.POST("/users", CreateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
