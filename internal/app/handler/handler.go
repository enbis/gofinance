package handler

import (
	"fmt"
	"net/http"
)

func Request(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler")
}
