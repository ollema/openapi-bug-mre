package api

import (
	"encoding/json"
	"net/http"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

// Server implements the ServerInterface for OpenAPI operations.
type Server struct{}

// NewServer creates a new Server instance.
func NewServer() Server {
	return Server{}
}

// GetPingUserID handles the GET /ping/{userID} endpoint.
func (Server) GetPingUserID(w http.ResponseWriter, _ *http.Request, userID string) {
	resp := Pong{
		Ping: "Hello, " + userID + "!",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
