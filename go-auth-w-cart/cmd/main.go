package main

import (
	"fmt"
	"log"
	"os"

	"github.com/brikiernan/go-auth-w-cart/cmd/api"
	"github.com/brikiernan/go-auth-w-cart/db"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	server := api.NewAPIServer(port, db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
