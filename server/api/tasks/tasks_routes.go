package tasks

import "github.com/gorilla/mux"

func RegisterTaskRoutes(router *mux.Router, taskController *TaskController) {
	taskRouter := router.PathPrefix("/api/tasks").Subrouter()

	taskRouter.HandleFunc("/", taskController.GetTasks).Methods("GET")
	taskRouter.HandleFunc("/", taskController.CreateTask).Methods("POST")
	taskRouter.HandleFunc("/{id:[0-9]+}", taskController.DeleteTask).Methods("DELETE")
}