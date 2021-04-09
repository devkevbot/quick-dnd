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

	// Unprotected stat endpoints
	e.GET("/stat", app.retrieveAllStats)

	// All routes which require JWT-based authentication
	r := e.Group("/auth")
	r.Use(middleware.JWTWithConfig(app.getJWTConfig()))
	r.GET("/player/:username", app.retrievePlayer)
	r.PUT("/player/me/password", app.changePlayerPassword)
	r.DELETE("/player/me", app.deletePlayerSelf)

	// Protected character endpoints
	r.POST("/character", app.createCharacter)
	r.GET("/character/me", app.retrieveUserCharacters)
	r.DELETE("/character/:id", app.deleteCharacter)

	// Protected spell endpoints
	r.POST("/character/:id/spell", app.createSpell)
	r.GET("/character/:id/spell/:name", app.retrieveSpell)
	r.GET("/character/:id/spell", app.retrieveAllCharacterSpells)

	// Protected item endpoints
	r.POST("/character/:id/item", app.createItem)
	r.GET("/character/:id/item/:name", app.retrieveItem)
	r.GET("/character/:id/item", app.retrieveAllCharacterItems)
	r.DELETE("/character/:id/item/:name", app.deleteItem)

	// Protected campaign endpoints
	r.POST("/campaign", app.createCampaign)
}
