package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/v2", version2)
	http.HandleFunc("/legacy", SimpleRedirector)
	http.HandleFunc("/dashboard", SecureDashboard)
	http.HandleFunc("/agent", UserAgentEchoHandler)
	http.HandleFunc("/calculate", BasicMathAPI)
	http.HandleFunc("/count", TextCounter)
	http.HandleFunc("/ping", pongHandler)
	http.HandleFunc("/hello", QueryHandler)
	fmt.Println("server listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
