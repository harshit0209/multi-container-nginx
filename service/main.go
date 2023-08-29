package main

import (
	"fmt"
	"net/http"
	"os"
)

var port string

func handler(w http.ResponseWriter, r *http.Request) {
	message := "Hello, this is the HTTP server!   >>>port="
	message = fmt.Sprintf("%v%v", message, port)
	w.Write([]byte(message))
}

func main() {
	host := os.Getenv("HTTP_HOST")
	if host == "" {
		host = "localhost"
	}

	port = os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}

	route := os.Getenv("HTTP_ROUTE")
	if route == "" {
		route = "/api"
	}

	addr := host + ":" + port

	http.HandleFunc(route, handler)

	fmt.Println("HTTP server is listening on", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
