package main

import (
	"flag"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	portPtr := flag.Int("p", 8080, "Port Number")
	flag.Parse()

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
			<h1>Hello World!</h1>
		`)
	})
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(*portPtr)))
}
