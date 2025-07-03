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

// (GET /ping/{userId})
func (Server) GetPingUserId(w http.ResponseWriter, r *http.Request, userId string) {
	resp := Pong{
		Ping: "Hello, " + userId + "!",
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
