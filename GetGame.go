package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// Game is about category of streams.
type Game struct {
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
}

// GetGame return information about game category.
func GetGame(game interface{}) (*Game, error) {
	var seo bool
	switch game.(type) {
	case string:
		seo = true
	case int:
		seo = false
	default:
		return nil, fmt.Errorf("game mast be only as string or int")
	}

	var args f.Args
	args.Add("seo", strconv.FormatBool(seo))

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("game/", game))
	resp, err := just.GET(url, &args)
	if err != nil {
		return nil, err
	}

	var obj = struct {
		Request struct {
			This string `json:"this"`
		} `json:"request"`
		Game `json:"category"`
	}{}
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return &obj.Game, nil
}
