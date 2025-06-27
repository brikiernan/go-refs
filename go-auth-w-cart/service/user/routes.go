package user

import (
	"fmt"
	"net/http"

	"github.com/brikiernan/go-auth-w-cart/types"
	"github.com/brikiernan/go-auth-w-cart/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /users/{id}", h.handleUser)
	router.HandleFunc("POST /login", h.handleLogin)
	router.HandleFunc("POST /register", h.handleRegister)
}

func (h *Handler) handleUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("id")
	w.Write([]byte("User ID: " + userID))
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Path: " + r.URL.Path))
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		errMessage := fmt.Errorf("user with email %s already exists", payload.Email)
		utils.WriteError(w, http.StatusBadRequest, errMessage)
		return
	}

	hasedPassword := ""

	err = h.store.CreateUser(types.User{
		First:    payload.First,
		Last:     payload.Last,
		Email:    payload.Email,
		Password: hasedPassword,
		Role:     payload.Role,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
