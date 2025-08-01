package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir("./static")))
    http.HandleFunc("/api/hello", helloHandler)
    http.HandleFunc("/api/time", timeHandler)
    http.HandleFunc("/api/echo", echoHandler)
    http.HandleFunc("/api/status", statusHandler)

    fmt.Println("Server is running at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{"message": "Hello from Go 🚀"})
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{"time": time.Now().UTC().Format(time.RFC3339)})
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
    var body map[string]string
    json.NewDecoder(r.Body).Decode(&body)
    json.NewEncoder(w).Encode(map[string]string{"echo": body["text"]})
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{"status": "ok", "uptime": "10s"})
}
