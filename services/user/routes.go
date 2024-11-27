package user

import (
	"fmt"
	"net/http"

	"github.com/brumble9401/golang-authentication/services/auth"
	"github.com/brumble9401/golang-authentication/types"
	"github.com/brumble9401/golang-authentication/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriterError(w, http.StatusBadRequest, err)
	}

	_, err := h.store.GetUserByEmailOrUsername(payload.Email, payload.Username)

	if err == nil {
		utils.WriterError(w, http.StatusBadRequest, fmt.Errorf("username or email already exists"))
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriterError(w, http.StatusInternalServerError, err)
	}
	err = h.store.CreateUser(types.User{
		Username: payload.Username,
		Email:     payload.Email,
		PasswordHash:  hashedPassword,
		FullName: payload.FullName,
	})
	if err != nil {
		utils.WriterError(w, http.StatusInternalServerError, fmt.Errorf("failed to create user: %v", err))
		return
	}

	// Respond with success (you can include more details in the response if needed)
	utils.WriteJSON(w, http.StatusCreated, "User created successfully")
}