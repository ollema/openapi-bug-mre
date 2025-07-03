package api

import (
	"encoding/json"
	"net/http"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() Server {
	return Server{}
}

// (GET /ping/{userID})
func (Server) GetPingUserID(w http.ResponseWriter, r *http.Request, userID string) {
	resp := Pong{
		Ping: "Hello, " + userID + "!",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
