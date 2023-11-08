package handler

import (
	service "BwiBOW123/backend-miniproject/internal/app/services"
	"BwiBOW123/backend-miniproject/internal/domain"
	"BwiBOW123/backend-miniproject/logs"
	"encoding/json"
	"log"

	"net/http"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.Error(err)
		return
	}

	if err := h.service.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logs.Error(err)
		return
	}
	log.Println(http.StatusCreated)
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	user, err := h.service.GetUserByUsername(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	log.Println("\033[32mUsername\033[0m: " + user.Username)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials domain.User
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.Error(err)
		return
	}

	user, err := h.service.LoginUser(credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		logs.Error(err)
		return
	}
	logs.Info("User logged in:" + user.Username)
	json.NewEncoder(w).Encode(user)
}
