package function

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

var Mux = newMux()

// F represents cloud function entry point
func WebServer(w http.ResponseWriter, r *http.Request) {
	Mux.ServeHTTP(w, r)

}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/one", one)
	mux.HandleFunc("/two", two)
	mux.HandleFunc("/subroute/three", three)
	fmt.Println("New Mux")

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair)
	}
	return mux
}

func one(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from one"))
}

func two(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from two"))
}

func three(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from three"))
}
