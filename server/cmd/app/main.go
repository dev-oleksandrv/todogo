package main

import (
	"log"
	"net/http"

	"github.com/dev-oleksandrv/api/tasks"
	"github.com/dev-oleksandrv/config"
	"github.com/dev-oleksandrv/db"
	"github.com/dev-oleksandrv/repository"
	"github.com/gorilla/mux"
)


func main() {
	// Loading configuration from godotenv
	config.LoadConfig()
	// Creating a connection to the database
	db.CreateConnection()
	// Creating a router
	r := mux.NewRouter()
	// Tasks Controller
	taskRepository := repository.NewGORMTaskRepository(db.DB)
	taskService := tasks.NewTaskService(*taskRepository)
	taskController := tasks.NewTaskController(taskService)
	tasks.RegisterTaskRoutes(r, taskController)
	// Run HTTP server
	log.Println("Running server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}