package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func addUsersAPI(e *echo.Echo) {
	e.GET("/users/:id", getUser)
	e.POST("/users", createUser)
}

func getUser(c echo.Context) error {
	id := c.Param("id") + "a"
	return c.String(http.StatusOK, id)
}

func createUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	address := c.FormValue("address")
	return c.String(http.StatusAccepted, "name="+name+", email="+email+", address="+address)
}
