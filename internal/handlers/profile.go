package handlers

import (
	"github.com/Eyob49/go-auth-service/internal/auth"
	"fmt"
	"log"
	"net/http"
)

func (h *AuthHandler) Profile(w http.ResponseWriter, r *http.Request) {
	log.Println("Profile handler reached")
	claimsValue := r.Context().Value(auth.UserClaimsKey)
	claims, ok := claimsValue.(*auth.UserClaims)
	if !ok || claims == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, "Welcome user %d! Your email is %s.", claims.UserID, claims.Email)
}
