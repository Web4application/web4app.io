package main

import (
  "encoding/json"
  "net/http"
)

func main() {
  http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Hello from Go backend!"})
  })

  http.Handle("/", http.FileServer(http.Dir("."))) // Serve static files

  http.ListenAndServe(":8080", nil)
}
