package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func BasicMathAPI(w http.ResponseWriter, r *http.Request) {
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")
	op := r.URL.Query().Get("op")
	num1, err := strconv.Atoi(a)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	num2, err := strconv.Atoi(b)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	var Result int

	switch op {
	case "add":
		Result = num1 + num2
	case "subtract":
		Result = num1 - num2
	case "multiply":
		Result = num1 * num2
	default:
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
	}

	fmt.Fprintf(w, "Result: %d", Result)
}
