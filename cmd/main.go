package main

import (
	"auth/store"
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


func main() {
	godotenv.Load()
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set.")
	}
	// Open the pool structures
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to open database pool: %v", err)
	}
	defer db.Close()

    // Force a real network connection check to ensure the database is reachable
    err = db.Ping()
	if err != nil {
		log.Fatalf("Database is unreachable: %v", err)
	}

	log.Println("Succesfully connected to the database.")



	err = store.InitDB(db)
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
}