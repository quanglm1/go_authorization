package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
