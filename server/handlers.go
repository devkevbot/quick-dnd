package main

import (
	"draco/examples"
	"net/http"

	"github.com/labstack/echo/v4"
)

// TestHandler is a function which will be called when a request to the
// "/test" endpoint is made.
func TestHandler(c echo.Context) error {
	u := &examples.User{
		Name:  "Kevin",
		Email: "kevin@email.com",
	}
	return c.JSON(http.StatusOK, u)
}
