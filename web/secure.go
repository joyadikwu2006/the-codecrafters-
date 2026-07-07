package main

import (
	"fmt"
	"net/http"
)

func SecureDashboard(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("X-API-Key")
	if key != "secret123" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Welcome to the dashboard")
	//w.Write([]byte("error found"))
}
