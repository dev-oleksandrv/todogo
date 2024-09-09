package main

import (
	"log"
	"net/http"

	"github.com/dev-oleksandrv/api/auth"
	"github.com/dev-oleksandrv/api/list"
	"github.com/dev-oleksandrv/api/space"
	"github.com/dev-oleksandrv/api/tasks"
	"github.com/dev-oleksandrv/db"
	"github.com/dev-oleksandrv/internal/config"
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
	// User Repository
	userRepository := repository.NewGORMUserRepository(db.DB)
	// Auth Controller
	authService := auth.NewAuthService(*userRepository)
	authController := auth.NewAuthController(authService)
	auth.RegisterAuthRoutes(r, authController)
	// Space Controller
	spaceRepository := repository.NewGORMSpaceRepository(db.DB) 
	spaceService := space.NewSpaceService(*spaceRepository)
	spaceController := space.NewSpaceController(spaceService)
	space.RegisterSpaceRoutes(r, spaceController)
	// List Controller
	listRepository := repository.NewGORMListRepository(db.DB)
	listService := list.NewListService(*listRepository)
	listController := list.NewListController(listService)
	list.RegisterListRoutes(r, listController)
	// Tasks Controller
	taskRepository := repository.NewGORMTaskRepository(db.DB)
	taskService := tasks.NewTaskService(*taskRepository)
	taskController := tasks.NewTaskController(taskService)
	tasks.RegisterTaskRoutes(r, taskController)
	// Closing DB after program exit
	sqlDb, err := db.DB.DB()
	if err == nil {
		defer sqlDb.Close()
	}
	// Run HTTP server
	log.Println("Running server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}