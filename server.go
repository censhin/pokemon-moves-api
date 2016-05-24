package main

import (
    "net/http"
    "log"
)

func main() {
    InitRoutes()
    http.ListenAndServe(":8001", nil)
    log.Println("Starting server on port 8001.")
}
