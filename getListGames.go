package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"strconv"
	"time"
)

type (
	ListRequest struct {
		Query    string
		Limit    int
		LiveOnly bool
	}

	ListGames struct {
		Request    Request    `json:"request"`
		Categories []Category `json:"categories"`
	}

	Category struct {
		ID         string    `json:"category_id"`
		Name       string    `json:"category_name"`
		NameShort  string    `json:"category_name_short,ommitempty"`
		SEOKey     string    `json:"category_seo_key"`
		Viewers    string    `json:"category_viewers"`
		MediaCount string    `json:"category_media_count"`
		Channels   string    `json:"category_channels,ommitempty"`
		LogoSmall  string    `json:"category_logo_small,ommitempty"`
		LogoLarge  string    `json:"category_logo_large"`
		Updated    time.Time `json:"category_updated"`
	}
)

// GetListGames returns a list games sorted by the number of viewers.
func GetListGames(req ListRequest) (ListGames, error) {
	var args fasthttp.Args
	if req.Query != "" {
		args.Add("q", req.Query) // Search keyword for `category_name`.
	}
	if req.Limit > 0 && req.Limit <= 100 {
		args.Add("limit", strconv.Itoa(req.Limit)) // Maximum number of objects to fetch. Default and maximum is 100.
	}
	args.Add("liveonly", strconv.FormatBool(req.LiveOnly)) // Return only games that have live channels.
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
