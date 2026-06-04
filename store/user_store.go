package store

import (
	"auth/internal/models"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
	               id SERIAL PRIMARY KEY,
				   email VARCHAR(255) UNIQUE NOT NULL,
				   password VARCHAR(255) NOT NULL,
				   created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP

	)`)
	if err != nil {
		return err
	}
	log.Println("Database initialized successfully")
	return nil
}

func CreateUser(db *sql.DB, email, password string) (*models.User, error) {
	var user models.User
	err := db.QueryRow(`INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id, email, created_at`, email, password).Scan(&user.ID, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	log.Println("User created successfully")
	return &user, nil
}

func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
	var user models.User
	err := db.QueryRow(`SELECT id, email, created_at FROM users WHERE email = $1`, email).Scan(&user.ID, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}