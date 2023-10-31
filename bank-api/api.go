package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleRetrieveAccount))

	fmt.Printf("Server running on port: %v", s.listenAddr)

	http.ListenAndServe(":"+strconv.Itoa(s.listenAddr), router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleListAccounts(w, r)
	}

	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}

	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("Method %v not allowed", r.Method)
}

func (s *APIServer) handleRetrieveAccount(w http.ResponseWriter, r *http.Request) error {
	paramId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(paramId)

	if err != nil {
		return err
	}

	account, err := s.store.RetrieveAccount(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, &account)
}

func (s *APIServer) handleListAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.ListAccounts()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, &accounts)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	params := CreateAccountParams{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		return err
	}

	account := NewAccount(params.FirstName, params.LastName)
	newAcct, err := s.store.CreateAccount(account)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusCreated, &newAcct)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, &APIError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr int
	store      Storage
}

func NewAPIServer(listenAddr int, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}
