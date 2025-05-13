package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
	"golang.org/x/crypto/acme/autocert"
)

// Define a WebSocket upgrader to handle connections
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP request to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	// Send a greeting message every 5 seconds
	for {
		message := fmt.Sprintf("Hello at %s", time.Now().Format(time.RFC1123))
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("Error sending message:", err)
			break
		}
		time.Sleep(5 * time.Second)
	}
}

func main() {
	// Use TLS with autocert (ACME certificate management)
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("localhost"), // Replace with your domain
		Cache:      autocert.DirCache("/tmp/certs"),
	}

	// Register the WebSocket handler
	http.HandleFunc("/ws", handleWebSocket)
	
	// Start the HTTPS server with the generated certificates
	log.Println("Secure WebSocket server started at wss://localhost:8080/ws")
	log.Fatal(http.ListenAndServeTLS(":8080", "", "", &certManager))
