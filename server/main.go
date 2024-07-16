package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// var client *http.Client

type CatFact struct {
	Fact   string
	Length int64
}

func handleRetrieveFact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	url := "https://catfact.ninja/fa"
	var catFact CatFact
	resp, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
		}{
			Message: "Unknown error",
		})
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&catFact)

	json.NewEncoder(w).Encode(catFact)
}

// func GetJson(url string, target interface{}) error {
// 	client := &http.Client{Timeout: 10 * time.Second}
// 	resp, err := client.Get(url)
// 	if err != nil {
// 		return err
// 	}

// 	defer resp.Body.Close()

// 	return json.NewDecoder(resp.Body).Decode(target)
// }

func main() {
	mux := http.NewServeMux()

	mux.Handle("/fact", http.HandlerFunc(handleRetrieveFact))

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		log.Fatal("PORT is undefined")
	}

	log.Print("Listening on port:", httpPort)
	err := http.ListenAndServe(":"+httpPort, mux)
	log.Fatal(err)
}

// finalHandler := http.HandlerFunc(final)
// 	mux.Handle("/", middlewareOne(middlewareTwo(finalHandler)))

// func middlewareOne(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Print("Executing middlewareOne")
// 		next.ServeHTTP(w, r)
// 		log.Print("Executing middlewareOne again")
// 	})
// }

// func middlewareTwo(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Print("Executing middlewareTwo")
// 		if r.URL.Path == "/foo" {
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 		log.Print("Executing middlewareTwo again")
// 	})
// }

// func final(w http.ResponseWriter, r *http.Request) {
// 	log.Print("Executing finalHandler")
// 	w.Write([]byte("OK"))
// }
