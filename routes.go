package main

import (
	"log"
	"net/http"

	"z-photos/util"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func home(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func photos(c echo.Context) error {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	data, err := util.GetPhotos()
	if err != nil {
		panic(err)
	}

	c.Response().Header().Set("HX-Redirect", "/photos")
	return c.Render(http.StatusOK, "photos.html", data)
}
