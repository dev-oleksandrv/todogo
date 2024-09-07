package db

import "log"

func Seed() {
	task := Task{Content: "Test", Description: "TaskDescription"}
	if err := DB.Create(&task).Error; err != nil {
		log.Fatalf("Failed to insert seed task: %v", err)
	}

	log.Println("Seeding was succesfuly completed")
}