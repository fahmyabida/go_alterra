package controller

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type LoginController struct {
}

func NewLoginController() LoginController {
	return LoginController{}
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
	if username != "joe" && password != "secret" {
		return c.JSON(400, map[string]interface{}{
			"message": "username / password not registered",
		})
	}
	jwToken := ctrl.CreateJWT(username)
	return c.JSON(200, map[string]interface{}{
		"token": jwToken,
	})
}

func (ctrl LoginController) CreateJWT(username string) string {
	timeNow := time.Now()
	fmt.Println(timeNow)
	fmt.Println(timeNow.Format(time.RFC3339))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, // header
		jwt.MapClaims{ // payload
			"username":   username,
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
