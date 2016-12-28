package hitGox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"strconv"
)

// Game is about category of streams.
type Game struct {
	Request struct {
		This string `json:"this"`
	} `json:"request"`
	Category struct {
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
	} `json:"category"`
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
		return nil, errors.New("game mast be only as string or int")
	}

	var args fasthttp.Args
	args.Add("seo", strconv.FormatBool(seo))

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("game/", game))
	resp, err := get(url, &args)
	if err != nil {
		return nil, err
	}

	var obj Game
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return &obj, nil
}
