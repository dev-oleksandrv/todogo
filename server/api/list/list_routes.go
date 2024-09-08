package list

import (
	"github.com/gorilla/mux"
)

func RegisterListRoutes(router *mux.Router, listController *ListController) {
	listRouter := router.PathPrefix("/api/list").Subrouter()

	listRouter.HandleFunc("/", listController.GetLists).Methods("GET")
	listRouter.HandleFunc("/", listController.CreateList).Methods("POST")
	listRouter.HandleFunc("/{id:[0-9]+}", listController.DeleteList).Methods("DELETE")
}