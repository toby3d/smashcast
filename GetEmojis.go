package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

// Team is about group information.
type EmojiObject struct {
	IconID          string      `json:"icon_id"`
	IconName        string      `json:"icon_name"`
	IconShort       string      `json:"icon_short"`
	IconShortAlt    string      `json:"icon_short_alt"`
	IconPath        string      `json:"icon_path"`
	CategoryID      string      `json:"category_id"`
	CategoryName    string      `json:"category_name"`
	CategoryLogo    string      `json:"category_logo"`
	Channel         interface{} `json:"channel"`
	IconPremiumOnly string      `json:"icon_premium_only,omitempty"`
}

// GetEmojis returns a list of channels a user follows.
func GetEmojis(userName, authToken string, premiumOnly bool) (*[]EmojiObject, error) {
	var args f.Args
	args.Add("authToken", authToken)
	args.Add("premiumOnly", strconv.FormatBool(premiumOnly))

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("chat/emotes/", userName))
	resp, err := just.GET(url, &args)
	if err != nil {
		return nil, err
	}

	var obj []EmojiObject
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return &obj, nil
}
