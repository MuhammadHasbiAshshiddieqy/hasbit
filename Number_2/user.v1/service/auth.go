package service

import (
	"net/http"
	"crypto/sha256"
    "encoding/hex"
	"encoding/json"

	"github.com/labstack/echo/v4"
	"user.v1/repository/mysql"
	"user.v1/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"os"
)

// Auth - for requesting token
func Auth(c echo.Context) error {
	req := &model.User{}
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &model.Auth{Message:"Please provide valid login details", Token:""})
	}

    h := sha256.New()
    h.Write([]byte(req.Password))
    sha256Hash := hex.EncodeToString(h.Sum(nil))

	db := mysql.DbManager() //.Debug
	u := &model.User{}
	err = db.Where("email = ? AND password = ?", req.Email, string(sha256Hash)).First(&u).Error
	if err != nil {
		return c.JSON(http.StatusUnauthorized, &model.Auth{Message:"Please signup", Token:""})
	}
	token, err := CreateToken(u.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, &model.Auth{Message:err.Error(), Token:""})
	}

	return c.JSON(http.StatusOK, &model.Auth{Message:"OK", Token:token})
}

// CreateToken - generate token for login
func CreateToken(userID uuid.UUID) (string, error) {
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	// atClaims["exp"] = time.Now().Add(time.Hour * 12).Unix() // uncomment if want to set expire
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
	   return "", err
	}
	return token, nil
  }