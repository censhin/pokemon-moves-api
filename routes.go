package main

import (
	"github.com/gorilla/mux"
)

func InitRoutes() (*mux.Router) {
	router := mux.NewRouter()

	router.HandleFunc("/moves", MovesResource)
	router.HandleFunc("/moves/{id}", MoveResource)

	return router
}
