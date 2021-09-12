package main

import (
	"draco/models"
	"draco/models/postgresql"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type application struct {
	echoInstance  *echo.Echo
	jwtSigningKey string
	players       interface {
		Insert(username string, password string, name string) error
		Authenticate(username string, password string) (string, error)
		Get(username string) (*models.Player, error)
		UpdatePassword(username string, newPassword string) error
		Delete(username string) error
	}
	characters interface {
		Insert(c models.Character) (int, error)
		Get(id int) (*models.Character, error)
		GetAllUserCharacters(username string) (*[]models.Character, error)
		Update(c models.Character) error
		Delete(id int) error
	}
	spells interface {
		Insert(s models.Spell) error
		Get(characterID int, spellName string) (*models.Spell, error)
		GetAllCharacterSpells(characterID int) (*[]models.Spell, error)
		Delete(characterID int, spellName string) error
		GetCountSpellsPerSchool(characterID int) (*[]models.SpellSchoolCountType, error)
	}
	items interface {
		Insert(i models.Item) error
		Get(characterID int, itemName string) (*models.Item, error)
		GetAllCharacterItems(characterID int) (*[]models.Item, error)
		Delete(characterID int, itemName string) error
		GetItemStats(characterID int) (*models.ItemStats, error)
	}
	campaigns interface {
		Insert(c models.Campaign, characterIDs []int) (int, error)
		Get(id int) (*models.Campaign, error)
		Update(id int, state string, location string) error
		Delete(id int) error
		GetPlayersCreatedCampaigns(dungeonMaster string) (*[]models.Campaign, error)
		GetAllCharacterCampaigns(characterID int) (*[]models.Campaign, error)
		GetCampaignParticpants(id int) (*[]models.CampaignParticipants, error)
		GetPlayersAttendedAll(dungeonMaster string) (*[]string, error)
	}
	milestones interface {
		Insert(campaignID int, milestone string) error
		GetAllForCampaign(campaignID int) (*[]string, error)
	}
	belongsTo interface {
		Insert(c models.BelongsTo) error
		GetAllCharacterCampaigns(characterID int) (*[]models.Campaign, error)
		GetAllCampaignCharacters(campaignID int) (*[]models.Character, error)
	}
	stats interface {
		GetAll() (*models.Stats, error)
	}
}

func (app *application) withDB(db *sqlx.DB) *application {
	app.players = &postgresql.PlayerModel{DB: db}
	app.characters = &postgresql.CharacterModel{DB: db}
	app.spells = &postgresql.SpellModel{DB: db}
	app.items = &postgresql.ItemModel{DB: db}
	app.campaigns = &postgresql.CampaignModel{DB: db}
	app.milestones = &postgresql.MilestoneModel{DB: db}
	app.belongsTo = &postgresql.BelongsToModel{DB: db}
	app.stats = &postgresql.StatsModel{DB: db}
	return app
}

func (app *application) withJWTSigningKey(key string) *application {
	key = strings.TrimSpace(key)
	if key == "" {
		app.echoInstance.Logger.Fatal("error: JWT signing key is empty/invalid")
	}

	app.jwtSigningKey = key
	return app
}

func (app *application) withEchoInstance(e *echo.Echo) *application {
	app.echoInstance = e
	return app
}

func (app *application) run(port int) {
	app.echoInstance.Logger.Fatal(
		app.echoInstance.Start(":" + strconv.Itoa(port)),
	)
}

func startServer(configFile *string) {
	cfg := createConfigFromFile(*configFile)
	db := initDBConn(cfg.CreatePostgreSQLDBConnString(false))

	var app application

	app.withDB(db).
		withJWTSigningKey(cfg.JWTSigningKey).
		withEchoInstance(echo.New())

	app.registerMiddleware()
	app.registerRoutes()

	app.run(cfg.HTTPServer.Port)
}
