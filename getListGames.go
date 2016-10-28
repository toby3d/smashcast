package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"strconv"
)

type (
	// ListRequest is filters for searching by game categories.
	ListRequest struct {
		// Search keyword for category_name.
		Query string

		// Maximum number of objects to fetch. Default and maximum is 100.
		Limit int

		// Return only games that have live channels.
		LiveOnly bool
	}

	// ListGames is a response body about find games categories.
	ListGames struct {
		Request    Request    `json:"request"`
		Categories []Category `json:"categories"`
	}

	// Category is a game category information.
	Category struct {
		ID         string    `json:"category_id"`
		Name       string    `json:"category_name"`
		NameShort  string    `json:"category_name_short,omitempty"`
		SEOKey     string    `json:"category_seo_key"`
		Viewers    string    `json:"category_viewers"`
		MediaCount string    `json:"category_media_count"`
		Channels   string    `json:"category_channels,omitempty"`
		LogoSmall  string    `json:"category_logo_small,omitempty"`
		LogoLarge  string    `json:"category_logo_large"`
		Updated    Timestamp `json:"category_updated"`
	}
)

// GetListGames returns a list games sorted by the number of viewers.
func GetListGames(req ListRequest) (ListGames, error) {
	var args fasthttp.Args
	if req.Query != "" {
		args.Add("q", req.Query)
	}
	if req.Limit > 0 && req.Limit <= 100 {
		args.Add("limit", strconv.Itoa(req.Limit))
	}
	args.Add("liveonly", strconv.FormatBool(req.LiveOnly))

	requestURL := fmt.Sprintf("%s/games?%s", API, args.String())
	_, body, err := fasthttp.Get(nil, requestURL)
	if err != nil {
		return ListGames{}, err
	}
	var obj ListGames
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return ListGames{}, err
	}
	return obj, nil
}
