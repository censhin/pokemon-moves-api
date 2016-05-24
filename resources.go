package main

import (
	"net/http"
	"log"
)

func MovesResource(w http.ResponseWriter, r *http.Request) {
	b, err := MovesService()

	if err != nil {
		log.Fatal(err)
	} else {
		w.Write(b)
	}
}
