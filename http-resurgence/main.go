package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mainMux := http.NewServeMux()
	apiMux := http.NewServeMux()

	mainMux.Handle("/api/", http.StripPrefix("/api", apiMux))

	mainMux.HandleFunc("/status", StatusCodeFactory)
	mainMux.HandleFunc("/form", FormDecoder)
	mainMux.HandleFunc("/headers", HeaderDetective)
	mainMux.HandleFunc("/echo", EchoChamber)
	mainMux.HandleFunc("/method-inspector", MethodHandler)
	apiMux.HandleFunc("/v1/ping", pingHandler)
	apiMux.HandleFunc("/v1/greet", greetHandler)

	fmt.Println("server listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", mainMux))
}
