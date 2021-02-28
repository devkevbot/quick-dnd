package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func initDBConn(connString string) *sqlx.DB {
	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal((err))
	}
	fmt.Println("Successfully connected to database")
	return db
}
