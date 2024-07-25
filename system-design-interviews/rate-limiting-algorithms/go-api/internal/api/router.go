package api

import (
    "time"
    "net/http"
    "go-api/internal/api/handlers"
    "go-api/internal/api/middleware"
    "github.com/gorilla/mux"
    "go-api/internal/services/nameservice"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    tb := nameservice.NewTokenBucket(1, time.Second) // Token Bucket middleware

    router.Use(middleware.LoggingMiddleware)
    router.HandleFunc("/health", handlers.GetHeatlh).Methods("GET")
    router.HandleFunc("/", handlers.GetHeatlh).Methods("GET")
    router.Handle("/getNames", middleware.RateLimitMiddleware(tb)(http.HandlerFunc(handlers.GetNames))).Methods("GET")

    return router
}
