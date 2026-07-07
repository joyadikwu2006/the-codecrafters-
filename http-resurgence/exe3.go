package main

import (
	"fmt"
	"net/http"
)

func HeaderDetective(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Custom-Token")
	if token == "" {
		http.Error(w, "X-Custom-Token header is missing", http.StatusBadRequest)
		return
	}
	content := r.Header.Get("Content-Type")
	if content == "" {
		fmt.Fprintf(w, "Token received: %s\nContent-Type: Content-Type not provided", token)
	} else {
		fmt.Fprintf(w, "Token received: %s\nContent-Type: %s", token, content)
	}
}
