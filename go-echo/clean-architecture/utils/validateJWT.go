package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func ValidateJwt() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			arrToken := strings.Split(token, " ")
			if len(arrToken) < 2 {
				return c.JSON(400, map[string]interface{}{
					"message": "token not found on header authorization",
				})
			}
			tokenJwt, err := jwt.Parse(arrToken[1], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte("abcde12345"), nil
			})
			if err != nil {
				return c.JSON(400, map[string]interface{}{
					"message": "error while parsing jwt",
					"error":   err,
				})
			}

			if !tokenJwt.Valid {
				return c.JSON(400, map[string]interface{}{
					"message": "token invalid",
				})
			}
			payloadJWT := tokenJwt.Claims.(jwt.MapClaims)
			expiredAt, _ := time.Parse(time.RFC3339, payloadJWT["expired_at"].(string))
			if expiredAt.Before(time.Now()) {
				return c.JSON(400, map[string]interface{}{
					"message": "token expired",
				})
			}
			err = next(c)
			return err
		}
	}
}
