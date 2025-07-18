package web4app

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gorilla/websocket"
)

const VERSION = "1.27.1"

// New creates a new web4app session with provided token.
// If the token is for a bot, it must be prefixed with "Bot "
// 		e.g. "Bot ..."
// Or if it is an OAuth2 token, it must be prefixed with "Bearer "
//		e.g. "Bearer ..."

func (s *SomeStruct) Initialize(token string, compress bool, threshold int) error {
    if token == "" {
        return fmt.Errorf("token is required")
    }

    if !isValidToken(token) {
        return fmt.Errorf("invalid token provided")
    }

    // Initialize the Identify Package with defaults
    s.Identify.Compress = compress
    s.Identify.LargeThreshold = threshold
    s.Identify.Properties.OS = runtime.GOOS
    s.Identify.Properties.Browser = "web4app v" + VERSION
    s.Identify.Intents = IntentsAllWithoutPrivileged
    s.Identify.Token = token
    s.Token = token

    return nil
}
