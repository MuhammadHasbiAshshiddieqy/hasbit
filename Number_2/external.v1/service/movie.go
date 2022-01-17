package service

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"github.com/labstack/echo/v4"
	"external.v1/model"
	"github.com/google/uuid"
	"external.v1/repository/mysql"
)

// GetMovies - get movies list from omdbapi
func GetMovies(c echo.Context) error {
	uAuth, userID := Authorize(c)
	if  !uAuth{
		return c.JSON(http.StatusUnauthorized, "Please check the token")
	}
	uID := uuid.MustParse(userID)
	search := c.Param("search")
	page := c.Param("page")
	endpoint := fmt.Sprintf("movies/%s/%s", search, page)
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=faf7e5bb&s=%s&page=%s", search, page)

	db := mysql.DbManager() //.Debug

	resp, err := http.Get(url)
	if err != nil {
		l := &model.Log{UserID:uID, URL:url, Endpoint:endpoint, Error:err.Error()}
		db.Create(l)
		return c.JSON(http.StatusBadRequest, "Please provide valid request")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		l := &model.Log{UserID:uID, URL:url, Endpoint:endpoint, Error:err.Error()}
		db.Create(l)
		return c.JSON(http.StatusInternalServerError, "Failed to read response from http://www.omdbapi.com/")
	}

	var dat interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
        l := &model.Log{UserID:uID, URL:url, Endpoint:endpoint, Error:err.Error()}
		db.Create(l)
		return c.JSON(http.StatusInternalServerError, "Failed Unmarshal response")
    }

	l := &model.Log{UserID:uID, URL:url, Endpoint:endpoint}
	db.Create(l)
	return c.JSON(http.StatusOK, dat)
}

// GetMovieDetail - get movie detail from omdbapi
func GetMovieDetail(c echo.Context) error {
	uAuth, userID := Authorize(c)
	if  !uAuth{
		return c.JSON(http.StatusUnauthorized, "Please check the token")
	}
	uID := uuid.MustParse(userID)
	id := c.Param("id")
	endpoint := fmt.Sprintf("movie/%s", id)
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=faf7e5bb&i=%s&plot=full", id)

	db := mysql.DbManager() //.Debug

	resp, err := http.Get(url)
	if err != nil {
		l := &model.Log{UserID:uID, URL:url, Endpoint:endpoint,Error:err.Error()}
		db.Create(l)
		return c.JSON(http.StatusBadRequest, "Please provide valid request")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		l := &model.Log{UserID:uID, URL:url, Endpoint:endpoint, Error:err.Error()}
		db.Create(l)
		return c.JSON(http.StatusInternalServerError, "Failed to read response from http://www.omdbapi.com/")
	}

	var dat interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
        l := &model.Log{UserID:uID, URL:url, Endpoint:endpoint, Error:err.Error()}
		db.Create(l)
		return c.JSON(http.StatusInternalServerError, "Failed Unmarshal response")
    }

	l := &model.Log{UserID:uID, URL:url, Endpoint:endpoint}
	db.Create(l)
	return c.JSON(http.StatusOK, dat)
}