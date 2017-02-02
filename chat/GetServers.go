package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	just "github.com/toby3d/hitGox/tools"
)

// Server just a element of array.
type Server struct {
	IP string `json:"server_ip"`
}

// GetServers return array with objects containing server addresses.
//
// You should pick a server from the response, but you should never save the address. The servers are ever changing and getting the list from the API on every start is the safest method.
//
// At this point you should also get a authentication token for the user unless you are planning to login as a guest.
func GetServers() ([]Server, error) {
	url := fmt.Sprintf(APIEndpoint, "chat/servers")
	resp, err := just.GET(url, nil)
	if err != nil {
		return nil, err
	}

	var obj []Server
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	log.Printf("%#v", obj)

	return obj, nil
}
