package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func ConnectDB(dbURL string) (*sql.DB, error) {
	
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database!")

	return db,nil


}

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

	log.Println("Database intialized successfully")

	return nil

}