package main

import (
	"fmt"
	"log"

	"auth-service/cmd/api"
	"auth-service/db"
)

func main() {

	db, err := db.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	port := fmt.Sprintf(":%s", "8080")
	server := api.NewAPIServer(port, db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
