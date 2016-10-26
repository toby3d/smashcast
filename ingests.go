package hitGox

import (
	"bytes"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

// Ingest is a single list element about streaming server for programs.
type Ingest struct {
	Location string `json:"ingest_location"`
	URL      string `json:"ingest_url"`
}

// Ingests returns a list of servers for live stream ingesting.
func Ingests() ([]Ingest, error) {
	_, body, err := fasthttp.Get(nil, API+"/ingests/default_list")
	if err != nil {
		return nil, err
	}
	var obj []Ingest
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return nil, err
	}
	return obj, nil
}
