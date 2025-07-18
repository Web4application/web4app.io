package web4app.io

import (
	"fmt"
	"sync"
)

// EventHandler is an interface for handling specific events.
type EventHandler interface {
	// Type returns the type of the event this handler handles.
	Type() string

	// Handle processes the event when triggered.
	Handle(*Session, interface{})
}

// EventInterfaceProvider is an interface that provides an instance for an event.
type EventInterfaceProvider interface {
	// Type returns the event type.
	Type() string

	// New returns a new instance of the struct for the event.
	New() interface{}
}

// interfaceEventType is a constant used to represent generic interface{} events.
const interfaceEventType = "__web4app__"

// Session manages the WebSocket session, event handlers, and state.
type Session struct {
	handlers     map[string][]EventHandler
	onceHandlers map[string][]EventHandler
	handlersMu   sync.RWMutex
	SyncEvents   bool
}

// NewSession creates a new Session.
func NewSession() *Session {
	return &Session{
		handlers:     make(map[string][]EventHandler),
		onceHandlers: make(map[string][]EventHandler),
	}
}

// AddHandler adds a persistent event handler.
func (s *Session) AddHandler(handler EventHandler) func() {
	s.handlersMu.Lock()
	defer s.handlersMu.Unlock()

	// Add the handler to the handlers map
	s.handlers[handler.Type()] = append(s.handlers[handler.Type()], handler)

	// Return a function to remove the handler
	return func() {
		s.removeEventHandler(handler.Type(), handler)
	}
}

// AddHandlerOnce adds an event handler that fires only once.
func (s *Session) AddHandlerOnce(handler EventHandler) func() {
	s.handlersMu.Lock()
	defer s.handlersMu.Unlock()

	// Add the handler to the onceHandlers map
	s.onceHandlers[handler.Type()] = append(s.onceHandlers[handler.Type()], handler)

	// Return a function to remove the handler
	return func() {
		s.removeEventHandlerOnce(handler.Type(), handler)
	}
}

// removeEventHandler removes a persistent event handler.
func (s *Session) removeEventHandler(t string, handler EventHandler) {
	s.handlersMu.Lock()
	defer s.handlersMu.Unlock()

	handlers := s.handlers[t]
	for i := 0; i < len(handlers); i++ {
		if handlers[i] == handler {
			s.handlers[t] = append(handlers[:i], handlers[i+1:]...)
			return
		}
	}
}

// removeEventHandlerOnce removes a one-time event handler.
func (s *Session) removeEventHandlerOnce(t string, handler EventHandler) {
	s.handlersMu.Lock()
	defer s.handlersMu.Unlock()

	handlers := s.onceHandlers[t]
	for i := 0; i < len(handlers); i++ {
		if handlers[i] == handler {
			s.onceHandlers[t] = append(handlers[:i], handlers[i+1:]...)
			return
		}
	}
}

// DispatchEvent triggers all registered handlers for a specific event type.
func (s *Session) DispatchEvent(t string, data interface{}) {
	s.handlersMu.RLock()
	defer s.handlersMu.RUnlock()

	// Call all persistent handlers
	for _, handler := range s.handlers[t] {
		if s.SyncEvents {
			handler.Handle(s, data)
		} else {
			go handler.Handle(s, data)
		}
	}

	// Call all one-time handlers and then remove them
	if handlers, ok := s.onceHandlers[t]; ok {
		for _, handler := range handlers {
			if s.SyncEvents {
				handler.Handle(s, data)
			} else {
				go handler.Handle(s, data)
			}
		}
		// Clear once handlers after they are fired
		s.onceHandlers[t] = nil
	}
}

// EventHandlerFactory creates handlers for specific events.
func EventHandlerFactory(eventType string, handler EventHandler) EventInterfaceProvider {
	// A factory for generating event handlers can be implemented here.
	// This can be expanded to create handlers based on specific event data types.
	return nil // Placeholder for future implementation
}
