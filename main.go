package main

import (
	"net/http"

	"github.com/jdpillaris/matrix-ops/handlers"
)

func main() {
	http.HandleFunc("/echo", handlers.EchoMatrix)
	http.ListenAndServe(":8080", nil)
}
