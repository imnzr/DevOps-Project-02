package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJsonError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Code":    statusCode,
		"Status":  "Error",
		"Message": message,
	})
}
