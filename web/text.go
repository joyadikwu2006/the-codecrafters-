package main

import (
	"io"
	"net/http"
	"strconv"
)

func TextCounter(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("Send a POST request with text to count words"))
	}

	if r.Method == http.MethodPost {
		Body, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("error reading file"))
		}
		length := len(Body)
		w.Write([]byte(strconv.Itoa(length)))
	}
	w.Write([]byte("method not allowed"))
}
