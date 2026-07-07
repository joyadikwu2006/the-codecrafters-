package main

import (
	"fmt"
	"net/http"
)

func MethodHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You made a %v request", r.Method)
}
