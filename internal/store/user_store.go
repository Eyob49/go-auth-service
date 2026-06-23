package store

import (
	"github.com/Eyob49/go-auth-service/internal/models"
	"database/sql"
	"log"
)

type UserStore struct {
	DB *sql.DB
}

func (s *UserStore) CreateUser(user models.User) (*models.User, error) {
	err := s.DB.QueryRow(`INSERT INTO users (email, password) VALUES($1, $2)
	                   RETURNING id, email, created_at`, user.Email, user.Password).Scan(&user.ID, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	log.Println("User created successfully")
	return &user, nil
}

func (s *UserStore) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.DB.QueryRow(`SELECT id, email, password, created_at FROM users WHERE email = $1
	                      `, email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
