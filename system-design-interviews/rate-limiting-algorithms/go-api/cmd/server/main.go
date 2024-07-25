package main

import (
    "log"
    "net/http"
    "go-api/internal/api"
)

func main() {
    router := api.NewRouter()
    log.Fatal(http.ListenAndServe(":8080", router))
}
