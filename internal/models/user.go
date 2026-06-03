package models

import "time"

type User struct {
	ID     int    `json:"id"`
	Email string `json:"email"`
	Password string `json:"-"` // Exclude password from JSON reponses
	CreatedAt time.Time `json:"created_at"`
}