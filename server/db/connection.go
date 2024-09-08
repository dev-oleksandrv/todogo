package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func CreateConnection() {
	connStr := os.Getenv("DATABASE_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
		PrepareStmt: false,
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	log.Println("Database connection established")
}