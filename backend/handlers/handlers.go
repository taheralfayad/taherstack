package handlers

import (
	"backend/responses"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		responses.MethodNotAllowed(w)
		return
	}

	responses.JSON(w, http.StatusOK, responses.Payload{"status": "ok"})
}
