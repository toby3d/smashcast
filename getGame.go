package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

// Game is a response body.
type Game struct {
	Request    Request  `json:"request"`
	Categories Category `json:"category"`
}

// GetGameByID return information about game category by gameID (eg. 1)
func GetGameByID(gameID int) (Game, error) {
	requestURL := fmt.Sprintf("%s/game/%d", API, gameID)
	_, body, err := fasthttp.Get(nil, requestURL)
	if err != nil {
		return Game{}, err
	}
	var obj Game
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return Game{}, err
	}
	return obj, nil
}

// GetGameByKey return information about game category by `gameKey` (eg. league-of-legends)
func GetGameByKey(gameKey string) (Game, error) {
	var args fasthttp.Args
	args.Add("seo", "true") // If using a game name, this must be true.
	requestURL := fmt.Sprintf("%s/game/%s?%s", API, gameKey, args.String())
	_, body, err := fasthttp.Get(nil, requestURL)
	if err != nil {
		return Game{}, err
	}
	var obj Game
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return Game{}, err
	}
	return obj, nil
}
