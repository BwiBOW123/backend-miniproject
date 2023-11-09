package handler

import (
	service "BwiBOW123/backend-miniproject/internal/app/services"
	"BwiBOW123/backend-miniproject/internal/domain"
	"BwiBOW123/backend-miniproject/logs"
	"encoding/json"
	"log"
	"net/http"
)

type LoginResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	User    domain.Member `json:"user"`
}

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user domain.Member
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.Error(err)
		return
	}

	if err := h.service.Register(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logs.Error(err)
		return
	}
	logs.Info("201 Create Suscess")
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
	var credentials domain.Member
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.Error(err)
		return
	}

	user, err := h.service.Login(credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, "Login Fail", http.StatusUnauthorized)
		logs.Debug(err.Error())
		return
	}
	resp := LoginResponse{
		Status:  "ok",
		Message: "Logged in",
		User:    *user,
	}
	logs.Info("User logged in:" + user.Username)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
