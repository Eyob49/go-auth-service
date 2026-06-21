package main

import (
	"auth/internal/database"
	"auth/internal/handlers"
	"net/http"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file:", err)
	}
	dbURL := os.Getenv("DATABASE_URL")
	db, err := database.ConnectDB(dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = database.InitDB(db)
	if err != nil {
		log.Fatal(err)
	}
    
	handler := &handlers.AuthHandler{DB: db}
	mux := http.NewServeMux()
	mux.HandleFunc("POST /register", handler.Register)
    mux.HandleFunc("POST /login", handler.Login)
	log.Printf("Server is running on http://localhost:8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}