package main

import (
	"fmt"
	"net/http"

	"github.com/jdpillaris/matrix-ops/handlers"
)

func main() {
	setupRoutes()

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func setupRoutes() {
	http.HandleFunc("/echo", handlers.EchoMatrix)
	http.HandleFunc("/invert", handlers.TransposeMatrix)
}
