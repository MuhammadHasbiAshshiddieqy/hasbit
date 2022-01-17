package service

import (
	"net/http"
	"crypto/sha256"
    "encoding/hex"
	"encoding/json"

	"github.com/labstack/echo/v4"
	"user.v1/repository/mysql"
	"user.v1/model"
)

// AddUser - add new user
func AddUser(c echo.Context) error {
	req := &model.User{}
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Please provide valid signup details")
	}

    h := sha256.New()
    h.Write([]byte(req.Password))
    sha256Hash := hex.EncodeToString(h.Sum(nil))
	req.Password = sha256Hash

	db := mysql.DbManager() //.Debug
	err = db.Create(&req).Error
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to create user")
	}

	return c.JSON(http.StatusOK, "Success")
}