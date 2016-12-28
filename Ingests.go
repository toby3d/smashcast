package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Ingest is a single list element about streaming server for programs.
type Ingest []struct {
	Location string `json:"ingest_location"`
	URL      string `json:"ingest_url"`
}

// Ingests returns a list of servers for live stream ingesting.
func Ingests() (Ingest, error) {
	url := fmt.Sprintf(APIEndpoint, "ingests/default_list")
	resp, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	var obj Ingest
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return obj, nil
}
