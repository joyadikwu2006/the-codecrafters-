package main

import (
	"io"
	"net/http"
)

func EchoChamber(w http.ResponseWriter, r *http.Request) {
	/*switch r.Method {
	case "POST":
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		if len(data) == 0 {
			http.Error(w, "body cannot be empty", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(data))
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}*/
	if r.Method == http.MethodPost {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		if len(data) == 0 {
			http.Error(w, "body cannot be empty", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(data))
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
