package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"log"
	"golang.org/x/crypto/bcrypt"
	"auth/internal/store"
	"auth/internal/models"
)

type AuthHandler struct {
	DB *sql.DB
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {

	type RegisterRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" {
		http.Error(w, "Missing Required Fields", http.StatusUnprocessableEntity)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost,)
	if err != nil {
		log.Println("Error hashing password:", err)
		return
	}

	user := models.User {
		Email: req.Email,
		Password: string(hashedPassword), 
	}

    userStore := store.UserStore{
		DB: h.DB,
	}
	createdUser, err := userStore.CreateUser(user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)

}