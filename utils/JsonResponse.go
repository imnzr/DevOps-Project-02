package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJsonError(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
