package models

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var client *sql.DB

func Init() {
	connStr := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL: %v", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping PostgreSQL: %v", err)
	}

	client = db
}
