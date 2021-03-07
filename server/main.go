package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

func main() {
	portPtr := flag.Int("p", 3000, "Port Number")
	configFile := flag.String("cfg", "config.yml", "Configuration YAML File")
	flag.Parse()

	cfg := createConfigFromFile(*configFile)

	db := initDBConn(cfg.CreatePostgreSQLDBConnString(false))
	fmt.Println(db)

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.GET("/test", TestHandler)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(*portPtr)))
}
