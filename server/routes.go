package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (app *application) registerRoutes(e *echo.Echo) {
	e.POST("/login", app.loginPlayer)
	e.POST("/register", app.createPlayer)

	// Unprotected character endpoints
	e.GET("/character/:id", app.retrieveCharacter)

	// All routes which require JWT-based authentication
	r := e.Group("/auth")
	r.Use(middleware.JWTWithConfig(app.getJWTConfig()))
	r.GET("/player/:username", app.retrievePlayer)

	// Protected character endpoints
	r.POST("/character", app.createCharacter)
	r.GET("/character/me", app.retrieveUserCharacters)

	// Protected spell endpoints
	r.POST("/character/:id/spell", app.createSpell)
	r.GET("/character/:id/spell/:name", app.retrieveSpell)
	r.GET("/character/:id/spell", app.retrieveAllCharacterSpells)
}
