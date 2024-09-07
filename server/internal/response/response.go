package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}