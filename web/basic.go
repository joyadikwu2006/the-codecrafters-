package main

import (
	"net/http"
)

func pongHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
