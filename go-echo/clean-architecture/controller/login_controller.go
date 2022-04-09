package controller

import (
	"belajar-go-echo/repository"
	"belajar-go-echo/utils"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

type LoginController struct {
	iUserRepo repository.IUserRepo
}

func NewLoginController(iUserRepo repository.IUserRepo) LoginController {
	return LoginController{iUserRepo}
}

func (ctrl LoginController) Login(c echo.Context) error {
	authorization := c.Request().Header.Get("authorization")
	arrAuth := strings.Split(authorization, " ")
	if len(arrAuth) != 2 {
		fmt.Println("header auth tidak valid")
		return c.JSON(400, map[string]interface{}{
			"message": "header auth tidak valid",
		})
	} else if arrAuth[0] != "Basic" {
		fmt.Println("haeder atuh must be basic")
		return c.JSON(400, map[string]interface{}{
			"message": "haeder atuh must be basic",
		})
	}
	// "Basic am9lOnNlY3JldA=="
	var decodedByte, _ = base64.StdEncoding.DecodeString(arrAuth[1])
	// joe:secret
	arrDecodeString := strings.Split(string(decodedByte), ":")
	username, password := arrDecodeString[0], arrDecodeString[1]
	user, err := ctrl.iUserRepo.GetUserByUsername(username)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "user not registered",
		})
	} else if user.Password != password {
		return c.JSON(400, map[string]interface{}{
			"message": "wrong password",
		})
	}
	jwToken := utils.CreateJWT(user)
	return c.JSON(200, map[string]interface{}{
		"token": jwToken,
	})
}
