package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func GetChatServers() ([]string, error) {
	url := fmt.Sprintf(APIEndpoint, "chat/servers")
	resp, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	var obj = []struct {
		ServerIP string `json:"server_ip"`
	}{}
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	var servers []string
	for i := range obj {
		servers = append(servers, obj[i].ServerIP)
	}

	return servers, nil
}
