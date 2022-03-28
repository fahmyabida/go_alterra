package main

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

var users []User

func main() {
	users = append(users, User{1, "fahmy", "abida"})
	users = append(users, User{2, "asa", "firdausi"})
	e := echo.New()
	e.GET("/users", GetAllUser)
	e.GET("/users/:id", GetSingleUser)
	e.POST("/users", CreateUser)
	e.PUT("/users/:id", UpdateUser)
	e.DELETE("/users/:id", DeleteUser)

	e.Start(":12345")
}

func UpdateUser(c echo.Context) error {
	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		return c.JSON(400, "id on url_param is invalid format, should be numeric")
	}
	var updateUser User
	c.Bind(&updateUser)
	dataIsFound := false
	for i, row := range users {
		if row.Id == id {
			// update
			users[i].FirstName = updateUser.FirstName
			users[i].LastName = updateUser.LastName
			updateUser.Id = id
			// end update
			dataIsFound = true
			break
		}
	}
	if dataIsFound == false {
		return c.JSON(200, map[string]interface{}{
			"message": fmt.Sprintf("user with id '%v' is not found", id),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"message": fmt.Sprintf("updae user with id '%v' is successful", id),
		"data":    updateUser,
	})
}
func DeleteUser(c echo.Context) error {
	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		return c.JSON(400, "id on url_param is invalid format, should be numeric")
	}
	dataIsFound := false
	for i, row := range users {
		if row.Id == id {
			users = remove(users, i)
			dataIsFound = true
			break
		}
	}
	if dataIsFound {
		return c.JSON(200, map[string]interface{}{
			"message": fmt.Sprintf("delete user with id '%v' is successful", id),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"message": fmt.Sprintf("user with id '%v' is not found", id),
	})
}

func remove(slice []User, s int) []User {
	return append(slice[:s], slice[s+1:]...)
}

func CreateUser(c echo.Context) error {
	var user User
	fmt.Println("before", user)
	c.Bind(&user)
	fmt.Println("after", user)
	user.Id = users[2-1].Id + 1 // [0]fahmy(1), [1]asa(2) --> users[1].Id +1  | len(2)-1
	fmt.Println("add id on user", user)
	users = append(users, user)
	return c.JSON(200, user)
}

func GetAllUser(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{
		"data":  users,
		"total": len(users),
	})
}

func GetSingleUser(c echo.Context) error {
	sId := c.Param("id")
	id, _ := strconv.Atoi(sId)
	var result User
	for _, user := range users {
		if id == user.Id {
			result = user
		}
	}
	if result.Id == 0 {
		return c.JSON(200, map[string]interface{}{
			"message": fmt.Sprintf("user not found with id %v", id),
		})
	}
	return c.JSON(200, result)
}
