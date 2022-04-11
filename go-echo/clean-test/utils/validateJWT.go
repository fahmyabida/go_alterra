package utils

import (
	"belajar-go-echo/model"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var secretKey_HMAC = []byte("abcde12345")

func ParsingJWT(c echo.Context) (user model.User, err error) {
	token := c.Request().Header.Get("Authorization")
	arrToken := strings.Split(token, " ")
	if len(arrToken) < 2 {
		err = errors.New("header Authorization invalid value")
		return user, err
	}
	tokenJwt, err := jwt.Parse(arrToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey_HMAC, nil
	})
	if err != nil {
		return user, err
	}

	if !tokenJwt.Valid {
		err = errors.New("token is invalid")
		return user, err
	}
	payloadJWT := tokenJwt.Claims.(jwt.MapClaims)
	floatId := payloadJWT["id"].(float64)
	user.Id = int(floatId)
	user.Role = payloadJWT["role"].(string)
	user.Name = payloadJWT["name"].(string)
	user.Username = payloadJWT["username"].(string)
	return user, err
}

func CreateJWT(user model.User) string {
	timeNow := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, // header
		jwt.MapClaims{ // payload
			"id":         user.Id,
			"username":   user.Username,
			"name":       user.Name,
			"role":       user.Role,
			"created_at": timeNow,
			"expired_at": timeNow.Add(1 * time.Hour),
		})
	tokenString, err := token.SignedString(secretKey_HMAC)
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}
