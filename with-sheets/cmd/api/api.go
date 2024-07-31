package api

import (
	"log"
	"net/http"
	"with-sheets/middleware"
	"with-sheets/service/google"
	sampleapps "with-sheets/service/sample-apps"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{addr: addr}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	cors := middleware.Cors()
	ctx, src := google.NewGoogleSheetsService()

	sampleAppSheets := sampleapps.NewSheets(ctx, src)
	sampleAppHandler := sampleapps.NewHandler(sampleAppSheets)
	sampleAppHandler.RegisterRoutes(router)

	subrouter := http.NewServeMux()
	subrouter.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	handler := cors.Handler(subrouter)

	server := http.Server{Addr: s.addr, Handler: handler}

	log.Printf("Server listening on port%s", s.addr)

	return server.ListenAndServe()
}
