package io

// Event represents a generic event.
type Event interface{}

// EventHandler defines the interface for handling events.
type EventHandler interface {
	Handle(event Event)
}

// MessageCreateHandler handles MessageCreate events.
type MessageCreateHandler struct{}

// Handle processes the MessageCreate event.
func (h *MessageCreateHandler) Handle(event Event) {
	// Implement your event handling logic here
}
