package handlers

import (
    "net/http"
	"encoding/json"
)

func GetHeatlh(w http.ResponseWriter, r *http.Request) {
    response := map[string]int{"status": 200}
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK) // Explicitly setting the status code to 200 OK
    json.NewEncoder(w).Encode(response)
}
