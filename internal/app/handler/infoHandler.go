package handler

import (
	"fmt"
	"net/http"

	"github.com/enbis/gofinance/internal/app/gofinance/foundamentals"
)

func Info(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("ticker")
	fmt.Println("ticker ricevuto ", t)
	foundamentals.Get(t)
}
