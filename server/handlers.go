package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type playerCreationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (app *application) createPlayer(c echo.Context) error {
	var req playerCreationRequest
	if err := c.Bind(&req); err != nil {
		log.Error(err)
		return sendJSONResponse(c, http.StatusUnprocessableEntity, "Player creation", "Could not process request", nil)
	}

	if err := app.players.Insert(req.Username, req.Password, req.Name); err != nil {
		log.Error(err)
		return sendJSONResponse(c, http.StatusInternalServerError, "Player creation", "Creation failed", nil)
	}

	return sendJSONResponse(c, http.StatusCreated, "Player creation", "Creation successful", nil)
}

// LoginRequest encapsulates a standard login request used to
// authenticate a player.
type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (app *application) loginPlayer(c echo.Context) error {
	var req loginRequest
	if err := c.Bind(&req); err != nil {
		log.Error(err)
		return sendJSONResponse(c, http.StatusUnprocessableEntity, "Player login", "Could not process request", nil)
	}

	username, err := app.players.Authenticate(req.Username, req.Password)
	if err != nil {
		log.Error(err)
		return sendJSONResponse(c, http.StatusUnauthorized, "Player login", "Login failed", nil)
	}

	token, err := app.createJWT(username)
	if err != nil {
		log.Error(err)
		return sendJSONResponse(c, http.StatusUnauthorized, "Player login", "Login failed", nil)
	}

	return sendJSONResponse(c, http.StatusOK, "Player login", "Login successful",
		struct {
			Username string `json:"username"`
			Token    string `json:"token"`
		}{
			username,
			token,
		},
	)
}

func (app *application) retrievePlayer(c echo.Context) error {
	requestedUsername := c.Param("username")
	tokenUsername := getUsernameFromToken(c)
	if tokenUsername != requestedUsername {
		return sendJSONResponse(c, http.StatusUnauthorized, "Player retrieval", "Access denied", nil)
	}

	player, err := app.players.Get(requestedUsername)
	if err != nil {
		log.Error(err)
		return sendJSONResponse(c, http.StatusNotFound, "Player retrieval", "Retrieval failed", nil)
	}

	return sendJSONResponse(c, http.StatusOK, "Player retrieval", "Retrieval successful", player)
}

type characterCreationRequest struct {
	Name          string `json:"name"`
	Weight        int    `json:"weight"`
	Height        int    `json:"height"`
	Alignment     string `json:"alignment"`
	Sex           string `json:"sex"`
	Background    string `json:"background"`
	Race          string `json:"race"`
	Speed         int    `json:"speed"`
	Strength      int    `json:"strength"`
	Dexterity     int    `json:"dexterity"`
	Intelligence  int    `json:"intelligence"`
	Wisdom        int    `json:"wisdom"`
	Charisma      int    `json:"charisma" `
	Constitution  int    `json:"constitution"`
	HPMax         int    `json:"hp_max"`
	AbilityPoints int    `json:"ability_points"`
	XPPoints      int    `json:"xp_points"`
	Class         string `json:"class"`
}

func (app *application) createCharacter(c echo.Context) error {
	var req characterCreationRequest
	if err := c.Bind(&req); err != nil {
		log.Error(err)
		return sendJSONResponse(c, http.StatusUnprocessableEntity, "Character creation", "Could not process request", nil)
	}

	creatorUsername := getUsernameFromToken(c)

	err := app.characters.Insert(
		req.Name, req.Weight, req.Height,
		req.Alignment, req.Sex, req.Background,
		req.Race, req.Speed, req.Strength,
		req.Dexterity, req.Intelligence,
		req.Wisdom, req.Charisma, req.Constitution,
		req.HPMax, req.AbilityPoints, req.XPPoints,
		req.Class, creatorUsername,
	)
	if err != nil {
		log.Error(err)
		return sendJSONResponse(c, http.StatusInternalServerError, "Character creation", "Creation failed", nil)
	}

	return sendJSONResponse(c, http.StatusCreated, "Character creation", "Creation successful", nil)
}

func (app *application) retrieveCharacter(c echo.Context) error {
	requestedName := c.Param("name")
	character, err := app.characters.Get(requestedName)
	if err != nil {
		log.Error(err)
		return sendJSONResponse(c, http.StatusNotFound, "Character retrieval", "Retrieval failed", nil)
	}

	return sendJSONResponse(c, http.StatusOK, "Character retrieval", "Retrieval successful", character)
}
