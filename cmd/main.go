package main

import (
	"auth/internal/database"
	"auth/internal/handlers"
	"auth/internal/middleware"
	"context"
	"net/http"
	"log"
	"os"
	"os/signal"
	"time"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	db, err := database.ConnectDB(dbURL)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize DB (if InitDB fails, close connection and exit)
	if err := database.InitDB(db); err != nil {
		db.Close()
		log.Fatal(err)
	}

	handler := &handlers.AuthHandler{DB: db}
	mux := http.NewServeMux()
	mux.HandleFunc("/register", handler.Register)
	mux.HandleFunc("/login", handler.Login)
	mux.Handle("/profile", middleware.AuthMiddleware(http.HandlerFunc(handler.Profile)))
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// Start server
	go func() {
		log.Printf("Server is running on http://localhost:%s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	// Graceful shutdown on interrupt
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	if err := db.Close(); err != nil {
		log.Printf("Error closing DB: %v", err)
	}

	log.Println("Server exiting")
}
