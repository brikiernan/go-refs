package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresStore() (*sql.DB, error) {
	dbUser := "someuser"
	dbName := "somename"
	dbPass := "somepass"

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", dbUser, dbName, dbPass)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected")

	return db, nil
}
