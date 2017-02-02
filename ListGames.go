package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// ListGames contains search results by games.
type ListGames struct {
	Request struct {
		This string `json:"this"`
	} `json:"request"`
	Category []struct {
		ID         string `json:"category_id"`
		Name       string `json:"category_name"`
		NameShort  string `json:"category_name_short"`
		SeoKey     string `json:"category_seo_key"`
		Viewers    string `json:"category_viewers"`
		MediaCount string `json:"category_media_count"`
		Channels   string `json:"category_channels"`
		LogoSmall  string `json:"category_logo_small"`
		LogoLarge  string `json:"category_logo_large"`
		Updated    string `json:"category_updated"`
	} `json:"categories"`
}

// GetListGames returns a list games sorted by the number of viewers.
func GetListGames(query string, limit int, liveOnly bool) (*ListGames, error) {
	var args f.Args
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

	url := fmt.Sprintf(APIEndpoint, "games")
	resp, err := just.GET(url, &args)
	if err != nil {
		return nil, err
	}

	var obj ListGames
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
