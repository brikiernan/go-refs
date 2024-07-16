package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"auth-service/types"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	h := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			First:    "Test",
			Last:     "User",
			Email:    "invalidemail",
			Password: "password",
			Role:     "client",
		}

		marshalled, _ := json.Marshal(payload)
		body := bytes.NewBuffer(marshalled)

		req, err := http.NewRequest(http.MethodPost, "/register", body)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("POST /register", h.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should correctly register the user", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			First:    "Test",
			Last:     "User",
			Email:    "valid@email.com",
			Password: "password",
			Role:     "client",
		}

		marshalled, _ := json.Marshal(payload)
		body := bytes.NewBuffer(marshalled)

		req, err := http.NewRequest(http.MethodPost, "/register", body)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("POST /register", h.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}

	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("")
}

func (m *mockUserStore) GetUserByID(id string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}
