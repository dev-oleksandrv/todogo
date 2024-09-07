package db

import (
	"database/sql"
	"log"
	"os"
)

func CreateConnection() sql.DB {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return *db
}