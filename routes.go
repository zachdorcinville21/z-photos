package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func home(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}
