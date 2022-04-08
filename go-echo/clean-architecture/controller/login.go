package controller

import (
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
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
	user, err := ctrl.iUserRepo.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "username / password not registered",
		})
	}
	jwToken := ctrl.CreateJWT(user)
	return c.JSON(200, map[string]interface{}{
		"token": jwToken,
	})
}

func (ctrl LoginController) CreateJWT(user model.User) string {
	timeNow := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, // header
		jwt.MapClaims{ // payload
			"username":   user.Username,
			"name":       user.Name,
			"created_at": timeNow,
			"expired_at": timeNow.Add(1 * time.Hour),
		})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("abcde12345"))
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}
