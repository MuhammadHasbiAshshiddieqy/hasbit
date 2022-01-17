package restapi

import (
	"external.v1/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SetupRouter - rest api routing
func SetupRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.GET("/movies/:search/:page", service.GetMovies)
	e.GET("/movie/:id", service.GetMovieDetail)

	return e
}