package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	e.GET("/users/:id", getUser)
	e.POST("/users", createUser)
	e.Logger.Fatal(e.Start(":8888"))
}
