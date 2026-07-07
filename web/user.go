package main

import (
	"net/http"
)

func UserAgentEchoHandler(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("User-Agent")
	if userAgent == "" {
		http.Error(w, "header empty", http.StatusNotFound)
		return
	}
	w.Write([]byte("You are visiting us using: " + userAgent))

}
