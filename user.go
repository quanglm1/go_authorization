package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func createUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	address := c.FormValue("address")
	return c.String(http.StatusAccepted, "name="+name+", email="+email+", address="+address)
}
