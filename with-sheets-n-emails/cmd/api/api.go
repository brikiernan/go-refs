package api

import (
	"log"
	"net/http"
	"with-sheets-n-emails/middleware"
	"with-sheets-n-emails/service/email"
	"with-sheets-n-emails/service/google"
	sampleapps "with-sheets-n-emails/service/sample-apps"
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
	smtpAddr, smtpAuth := email.Auth()

	sampleAppEmails := sampleapps.NewEmails(smtpAddr, smtpAuth)
	sampleAppSheets := sampleapps.NewSheets(ctx, src)
	sampleAppHandler := sampleapps.NewHandler(sampleAppSheets, sampleAppEmails)
	sampleAppHandler.RegisterRoutes(router)

	subrouter := http.NewServeMux()
	subrouter.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	handler := cors.Handler(subrouter)

	server := http.Server{Addr: s.addr, Handler: handler}

	log.Printf("Server listening on port%s", s.addr)

	return server.ListenAndServe()
}
