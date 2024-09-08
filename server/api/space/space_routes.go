package space

import (
	"github.com/gorilla/mux"
)

func RegisterSpaceRoutes(router *mux.Router, spaceController *SpaceController) {
	spaceRouter := router.PathPrefix("/api/space").Subrouter()

	spaceRouter.HandleFunc("/", spaceController.GetSpaces).Methods("GET")
	spaceRouter.HandleFunc("/", spaceController.CreateSpace).Methods("POST")
	spaceRouter.HandleFunc("/{id:[0-9]+}", spaceController.DeleteSpace).Methods("DELETE")
}