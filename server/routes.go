package main

import "github.com/labstack/echo/v4"

func (app *application) registerRoutes(e *echo.Echo) {
	e.POST("/player", app.createPlayer)
	e.GET("/player/:username", app.retrievePlayer)
	e.POST("/auth/login", app.loginPlayer)
}
