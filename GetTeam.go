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
type Team struct {
	Info struct {
		GroupID          string `json:"group_id"`
		FounderName      string `json:"founder_name"`
		GroupName        string `json:"group_name"`
		GroupDisplayName string `json:"group_display_name"`
		GroupText        string `json:"group_text"`
		GroupLogoSmall   string `json:"group_logo_small"`
		GroupLogoLarge   string `json:"group_logo_large"`
		GroupCover       string `json:"group_cover"`
		MembersTotal     int    `json:"members_total"`
	} `json:"info"`
	// Media struct{
	// 	Livestream []??
	// 	Video []???
	// }
	Members []struct {
		Followers       string `json:"followers"`
		Videos          string `json:"videos"`
		Recordings      string `json:"recordings"`
		Teams           string `json:"teams"`
		UserID          string `json:"user_id"`
		UserName        string `json:"user_name"`
		UserStatus      string `json:"user_status"`
		UserLogo        string `json:"user_logo"`
		UserCover       string `json:"user_cover"`
		UserLogoSmall   string `json:"user_logo_small"`
		UserPartner     string `json:"user_partner"`
		Admin           string `json:"admin"`
		Enabled         string `json:"enabled"`
		IsDefault       string `json:"is_default"`
		RevenuesEnabled string `json:"revenues_enabled"`
		GroupRole       string `json:"group_role"`
		GroupAccepted   bool   `json:"group_accepted"`
	} `json:"members"`
}

// GetTeam returns a team object for team.
func GetTeam(team, mediaType string, media, liveOnly, partner bool) (*Team, error) {
	var args f.Args
	switch {
	case media:
		args.Add("media", strconv.FormatBool(media))
		fallthrough
	case mediaType != "":
		args.Add("media_type", mediaType)
		fallthrough
	case liveOnly:
		args.Add("liveonly", strconv.FormatBool(liveOnly))
		fallthrough
	case partner:
		args.Add("partner", strconv.FormatBool(partner))
	}

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("/team/", team))
	resp, err := just.GET(url, &args)
	if err != nil {
		return nil, err
	}

	var obj Team
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return &obj, nil
}
