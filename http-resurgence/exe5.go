package main

import (
	"net/http"
	"strconv"
)

func StatusCodeFactory(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "code parameter is required", http.StatusBadRequest)
		return
	}
	num, err := strconv.Atoi(code)
	if err != nil {
		http.Error(w, "code must be a valid integer", http.StatusBadRequest)
		return
	}
	if num < http.StatusContinue || num > http.StatusInternalServerError {
		http.Error(w, "code must be a valid HTTP status code (100-599)", http.StatusNotFound)
		return
	}
	w.WriteHeader(num)
	w.Write([]byte("Responding with status [code]"))
}
