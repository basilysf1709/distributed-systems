package handlers

import (
    "encoding/json"
    "net/http"
    "go-api/internal/services/nameservice"
)

func GetNames(w http.ResponseWriter, r *http.Request) {
    names := nameservice.GenerateRandomNames(5, 5)
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(names)
}
