package main

import (
	"draco/models"
	"draco/models/postgresql"
	"flag"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

type application struct {
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
		GetTotalWeightCharacterItems(characterID int) (int, error)
	}
	campaigns interface {
		Insert(c models.Campaign) (int, error)
		Get(id int) (*models.Campaign, error)
		GetAllCharacterCampaigns(characterID int)
		GetPlayersAttendedAll(dungeonMaster string) (*[]string, error)
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

func main() {
	configFile := flag.String("cfg", "config.yml", "Configuration YAML File")
	flag.Parse()

	cfg := createConfigFromFile(*configFile)

	db := initDBConn(cfg.CreatePostgreSQLDBConnString(false))
	fmt.Println(db)

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
	}))

	if cfg.IsProduction {
		// For production builds, we will want to serve the minified
		// application bundle for frontend codes. This is created using
		// the `npm run build` command from the `/client/` directory of
		// the project.
		e.Static("/", "../client/dist")
	}

	app := &application{
		jwtSigningKey: cfg.JWTSigningKey,
		players:       &postgresql.PlayerModel{DB: db},
		characters:    &postgresql.CharacterModel{DB: db},
		spells:        &postgresql.SpellModel{DB: db},
		items:         &postgresql.ItemModel{DB: db},
		stats:         &postgresql.StatsModel{DB: db},
		campaigns:     &postgresql.CampaignModel{DB: db},
	}

	if cfg.IsTest {
		app.runTests(db)
	}

	app.registerRoutes(e)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(cfg.HTTPServer.Port)))
}
