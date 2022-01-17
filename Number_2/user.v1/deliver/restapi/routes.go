package restapi

import (
	"user.v1/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SetupRouter - rest api routing
func SetupRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.POST("/login", service.Auth)
	e.POST("/signup", service.AddUser)

	return e
}