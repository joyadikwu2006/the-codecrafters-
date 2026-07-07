package main

import (
	"fmt"
	"net/http"
)

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only Get Method Allowed", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s!\n", name)
}
