package main

import (
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/moves", MovesResource)
}
