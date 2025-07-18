
package main

import (
	"log"
	"net/http"
	"web4app.io/api"
)

func main() {
	mux := http.NewServeMux()

	// Serve static files from /public
	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/", fs)

	// API routes
	mux.HandleFunc("/api/hello", api.HelloHandler)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Server failed: ", err)
	}
}
