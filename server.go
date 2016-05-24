package main

import (
    "net/http"
)

func main() {
    InitRoutes()
    http.ListenAndServe(":8001", nil)
}
