package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"strconv"
)

// ListGames contains search results by games.
type ListGames struct {
	Request struct {
		This string `json:"this"`
	} `json:"request"`
	Categories []struct {
		CategoryID         string `json:"category_id"`
		CategoryName       string `json:"category_name"`
		CategoryNameShort  string `json:"category_name_short"`
		CategorySeoKey     string `json:"category_seo_key"`
		CategoryViewers    string `json:"category_viewers"`
		CategoryMediaCount string `json:"category_media_count"`
		CategoryChannels   string `json:"category_channels"`
		CategoryLogoSmall  string `json:"category_logo_small"`
		CategoryLogoLarge  string `json:"category_logo_large"`
		CategoryUpdated    string `json:"category_updated"`
	} `json:"categories"`
}

// GetListGames returns a list games sorted by the number of viewers.
func GetListGames(query string, limit int, liveOnly bool) (*ListGames, error) {
	var args fasthttp.Args
	switch {
	case query != "":
		args.Add("q", query)
		fallthrough
	case limit > 0 && limit <= 100:
		args.Add("limit", strconv.Itoa(limit))
		fallthrough
	default:
		args.Add("liveonly", strconv.FormatBool(liveOnly))
	}

	url := fmt.Sprintf(APIEndpoint, "/games")
	resp, err := get(url, &args)
	if err != nil {
		return nil, err
	}

	var obj ListGames
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
