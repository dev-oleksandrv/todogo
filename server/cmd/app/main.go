package main

import (
	"log"
	"net/http"

	"github.com/dev-oleksandrv/api/tasks"
	"github.com/dev-oleksandrv/config"
	"github.com/gorilla/mux"
)


func main() {
	config.LoadConfig()
	r := mux.NewRouter()
	taskController := tasks.NewTaskController()
	tasks.RegisterTaskRoutes(r, taskController)

	log.Println("Running server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}