package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/gorilla/websocket"
)

// Errors
var (
	ErrWSAlreadyOpen = errors.New("web socket already opened")
	ErrWSNotFound    = errors.New("web socket not found")
)

var msgID atomic.Uint64

// Request/Response structs
type SubscribeRequest struct {
	UserID    string `json:"user_id"`
	Plan      string `json:"plan"`
	Community string `json:"community"`
}

type JoinCommunityRequest struct {
	UserID    string `json:"user_id"`
	Community string `json:"community"`
}

type DeviceRebornRequest struct {
	UserID string `json:"user_id"`
	Mode   string `json:"mode"` // e.g. "web4"
}

// DAG Client
type DAGClient struct {
	BaseURL string // e.g., "http://subscription-service.mesh.local"
	Client  *http.Client
}

func NewDAGClient(baseURL string) *DAGClient {
	return &DAGClient{
		BaseURL: baseURL,
		Client:  &http.Client{},
	}
}

func (d *DAGClient) CallService(endpoint string, payload any) ([]byte, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/%s", d.BaseURL, endpoint)
	resp, err := d.Client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Message router
func handleWS(ws *websocket.Conn, dag *DAGClient) {
	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("WS closed:", err)
			break
		}

		// Basic envelope: {"action":"subscribe","payload":{...}}
		var envelope map[string]json.RawMessage
		if err := json.Unmarshal(msg, &envelope); err != nil {
			ws.WriteJSON(map[string]string{"error": "invalid envelope"})
			continue
		}

		var action string
		json.Unmarshal(envelope["action"], &action)

		switch action {
		case "subscribe":
			var req SubscribeRequest
			json.Unmarshal(envelope["payload"], &req)
			resp, err := dag.CallService("subscribe", req)
			if err != nil {
				ws.WriteJSON(map[string]string{"error": "subscribe failed"})
				continue
			}
			ws.WriteMessage(websocket.TextMessage, resp)

		case "joinCommunity":
			var req JoinCommunityRequest
			json.Unmarshal(envelope["payload"], &req)
			resp, err := dag.CallService("join", req)
			if err != nil {
				ws.WriteJSON(map[string]string{"error": "join failed"})
				continue
			}
			ws.WriteMessage(websocket.TextMessage, resp)

		case "deviceReborn":
			var req DeviceRebornRequest
			json.Unmarshal(envelope["payload"], &req)
			resp, err := dag.CallService("reborn", req)
			if err != nil {
				ws.WriteJSON(map[string]string{"error": "reborn failed"})
				continue
			}
			ws.WriteMessage(websocket.TextMessage, resp)

		default:
			ws.WriteJSON(map[string]string{"error": "unknown action"})
		}
	}
}

func main() {
	// Example: subscription service mesh endpoint
	dag := NewDAGClient("http://subscription-service.mesh.local")

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "Upgrade failed", http.StatusBadRequest)
			return
		}
		go handleWS(ws, dag)
	})

	log.Println("Web4 WebSocket server running on :8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
