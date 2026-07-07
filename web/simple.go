package main

import (
	"net/http"
)

func version2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to version 2"))
}

func SimpleRedirector(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/v2", http.StatusMovedPermanently)
}