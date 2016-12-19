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
		CategoryID         string      `json:"category_id"`
		CategoryName       string      `json:"category_name"`
		CategoryNameShort  interface{} `json:"category_name_short"`
		CategorySeoKey     string      `json:"category_seo_key"`
		CategoryViewers    string      `json:"category_viewers"`
		CategoryMediaCount string      `json:"category_media_count"`
		CategoryChannels   interface{} `json:"category_channels"`
		CategoryLogoSmall  interface{} `json:"category_logo_small"`
		CategoryLogoLarge  string      `json:"category_logo_large"`
		CategoryUpdated    string      `json:"category_updated"`
	} `json:"category"`
}

// GetGame return information about game category.
func GetGame(game interface{}) (*Game, error) {
	var seo bool

	switch game.(type) {
	case string:
		if game == "" {
			return nil, errors.New("game can not be empty")
		}
		seo = true
	case int:
		seo = false
	default:
		return nil, errors.New("game mast be only as string or int")
	}

	var args fasthttp.Args
	args.Add("seo", strconv.FormatBool(seo))

	url := fmt.Sprint(APIEndpoint, "/game/", game)
	resp, err := get(url, &args)
	if err != nil {
		return nil, err
	}

	var obj Game
	if err = json.NewDecoder(bytes.NewReader(resp)).Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
