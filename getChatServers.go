package hitGox

import (
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

// ChatServer just a sinlge item list about chat server IP.
type ChatServer struct {
	ServerIP string `json:"server_ip,omitempty"`
}

// GetChatServers return array with server addresses.
func GetChatServers() ([]ChatServer, error) {
	_, body, err := fasthttp.Get(nil, API+"/chat/servers")
	if err != nil {
		return nil, err
	}
	var obj []ChatServer
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return nil, err
	}
	return obj, nil
}
