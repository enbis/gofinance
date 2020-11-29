package handler

import (
	"encoding/json"
	"net/http"

	"github.com/enbis/gofinance/internal/app/gofinance/foundamentals"
)

type responsePayload struct {
	Data string `json:"data"`
}

func Request(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("data")
	data := foundamentals.Get(t)

	w.Header().Add("Content-Type", "application/json")

	// Prepare body
	json.NewEncoder(w).Encode(responsePayload{
		Data: data,
	})
}
