package main

import (
	"github.com/labstack/echo/v4/middleware"
)

func (app *application) registerMiddleware() {
	app.echoInstance.Use(middleware.Recover())
	app.echoInstance.Use(middleware.Logger())
	app.echoInstance.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:8080"},
		}))
}

func (app *application) registerRoutes() {
	app.echoInstance.POST("/login", app.loginPlayer)
	app.echoInstance.POST("/register", app.createPlayer)

	// Unprotected character endpoints
	app.echoInstance.GET("/character/:id", app.retrieveCharacter)

	// Unprotected stat endpoints
	app.echoInstance.GET("/stat", app.retrieveAllStats)

	// All routes which require JWT-based authentication
	r := app.echoInstance.Group("/auth")
	r.Use(middleware.JWTWithConfig(app.getJWTConfig()))
	r.GET("/player/:username", app.retrievePlayer)
	r.PUT("/player/me/password", app.changePlayerPassword)
	r.DELETE("/player/me", app.deletePlayerSelf)

	// Protected character endpoints
	r.POST("/character", app.createCharacter)
	r.GET("/character/me", app.retrieveUserCharacters)
	r.GET("/character/:id", app.retrieveCharacter)
	r.PUT("/character/:id", app.updateCharacter)
	r.DELETE("/character/:id", app.deleteCharacter)

	// Protected spell endpoints
	r.POST("/character/:id/spell", app.createSpell)
	r.GET("/character/:id/spell/:name", app.retrieveSpell)
	r.GET("/character/:id/spell", app.retrieveAllCharacterSpells)
	r.DELETE("/character/:id/spell/:name", app.deleteSpell)
	r.GET("/character/:id/spell/count-per-school", app.getCountSpellsPerSchool)

	// Protected item endpoints
	r.POST("/character/:id/item", app.createItem)
	r.GET("/character/:id/item/:name", app.retrieveItem)
	r.GET("/character/:id/item", app.retrieveAllCharacterItems)
	r.DELETE("/character/:id/item/:name", app.deleteItem)
	r.GET("/character/:id/item/stats", app.getItemStats)

	// Protected campaign endpoints
	r.POST("/campaign", app.createCampaign)
	r.PUT("/campaign/:id", app.updateCampaign)
	r.DELETE("/campaign/:id", app.deleteCampaign)
	r.POST("/campaign/milestone", app.createMilestone)
	r.GET("/campaign/:id/milestone", app.getAllMilestonesForCampaign)
	r.GET("/campaign/:id/participants", app.getCampaignParticipants)
	r.GET("/campaign/me/stats/player-attendance", app.getPlayersAttendedAll)
	r.GET("/campaign/me", app.getsPlayersCreatedCampaigns)
	r.GET("/character/:id/campaign", app.getAllCharacterCampaigns)

}
