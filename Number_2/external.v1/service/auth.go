package service

import (
	"github.com/labstack/echo/v4"

	"github.com/dgrijalva/jwt-go"
	"os"
)

// Authorize - check if token is valid for request
func Authorize(c echo.Context) (bool, string) {
	reqHead := c.Request().Header
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file

	token, err := jwt.Parse(reqHead.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return false, ""
	}

	claims := token.Claims.(jwt.MapClaims)
	auth := claims["authorized"].(bool)
	uID := claims["user_id"].(string)
	if auth {
		return auth, uID
	}

	return false, ""
}