package main

import (
	"log"

	"github.com/dev-oleksandrv/config"
	"github.com/dev-oleksandrv/db"
)

func main() {
	config.LoadConfig()
	db.CreateConnection()
	log.Println("Migration started")
	if err := db.DB.AutoMigrate(&db.Task{}); err != nil {
		log.Fatalf("Migration failed with error: %v", err)
		return
	}
	log.Println("Migration completed")
}