package main

import (
	"fmt"
	"net/http"
	"os"

	"runn/function"
)

func main() {
	mux := function.Mux

	fmt.Printf("Server running locally...")

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	http.ListenAndServe(":8080", mux)
}
