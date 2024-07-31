package main

import (
	"fmt"
	"log"
	"with-sheets/cmd/api"
	"with-sheets/utils"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	port := utils.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	server := api.NewAPIServer(addr)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
