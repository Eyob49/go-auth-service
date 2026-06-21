package auth

import (
	"auth/internal/models"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserID int64 `json:"user_id"`
	Email string  `json:"email"`
	jwt.RegisteredClaims
}



func GenerateJWT(user *models.User, secret string) (string, error){

	claims := UserClaims{
		UserID: user.ID,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "go-auth-service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}