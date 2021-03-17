package main

import (
	"io/ioutil"
	"log"

	"github.com/jmoiron/sqlx"
)

func (app *application) initTestDB(db *sqlx.DB) error {

	data, err := ioutil.ReadFile("sql/init_db.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(data))
	if err != nil {
		return err
	}

	app.players.Insert("ae1", "password1", "Albert Einstein")
	app.players.Insert("mc1", "password2", "Marie Curie")
	app.players.Insert("alove", "password3", "Ada Lovelace")
	app.players.Insert("inew", "password4", "Isaac Newton")
	app.players.Insert("at1", "password5", "Alan Turing")

	data, err = ioutil.ReadFile("sql/populate_db.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(data))
	return err
}

// This will initialize the mock DB. It is needed in order to
// insert correctly hashed passwords for example users. We can also
// use this space to run any automated tests for DB operations.
func (app *application) runTests(db *sqlx.DB) bool {
	err := app.initTestDB(db)
	if err != nil {
		log.Fatal(err)
	}
	return true
}
