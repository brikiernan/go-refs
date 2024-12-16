package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /hey", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET params were:", r.URL.Query())

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(map[string]string{
			"hey": "now",
		})
	})

	server := http.Server{Addr: ":8080", Handler: router}

	log.Printf("Server listening on port%s", server.Addr)

	err := server.ListenAndServe()

	log.Fatal(err)
}
