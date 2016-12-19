package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Ingest is a single list element about streaming server for programs.
type Ingest []struct {
	IngestLocation string `json:"ingest_location"`
	IngestURL      string `json:"ingest_url"`
}

// GetIngests returns a list of servers for live stream ingesting.
func GetIngests() (Ingest, error) {
	url := fmt.Sprint(API, "/ingests/default_list")
	resp, err := get(url, nil)
	if err != nil {
		return nil, err
	}

	var obj Ingest
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return obj, nil
}
