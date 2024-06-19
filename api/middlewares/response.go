package middlewares

import (
	"encoding/json"
	"net/http"
)

func SendErrorToClient(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := map[string]string{"error": message}
	json.NewEncoder(w).Encode(response)
}

func SendDataToClient(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func SendSuccessResponse(w http.ResponseWriter, message string) {
	w.Write([]byte(message))
	w.WriteHeader(http.StatusOK)
}
