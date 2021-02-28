package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

func main() {
	portPtr := flag.Int("p", 8080, "Port Number")
	configFile := flag.String("cfg", "config.yml", "Configuration YAML File")
	flag.Parse()

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	cfg := createConfigFromFile(*configFile)

	db := initDBConn(cfg.CreatePostgreSQLDBConnString(false))
	fmt.Println(db)

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
			<h1>Hello World!</h1>
		`)
	})
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(*portPtr)))
}
