package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type Payload map[string]string

func JSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("encode response failed: %v", err)
	}
}

func MethodNotAllowed(w http.ResponseWriter) {
	JSON(w, http.StatusMethodNotAllowed, Payload{"error": "method not allowed"})
}

func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, Payload{"error": message})
}
