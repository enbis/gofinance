package handler

import (
	"net/http"

	"github.com/enbis/gofinance/internal/app/gofinance/foundamentals"
)

func Request(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("data")
	foundamentals.Get(t)
}
