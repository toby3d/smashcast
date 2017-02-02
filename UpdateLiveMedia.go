package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

type update struct {
	MediaUserName    string   `json:"media_user_name"`
	MediaID          string   `json:"media_id"`
	MediaCategoryID  string   `json:"media_category_id"`
	MediaLiveDelay   string   `json:"media_live_delay"`
	MediaHidden      string   `json:"media_hidden"`
	MediaRecording   string   `json:"media_recording"`
	MediaMature      string   `json:"media_mature"`
	MediaHostedName  string   `json:"media_hosted_name"`
	MediaCountries   []string `json:"media_countries"`
	MediaStatus      string   `json:"media_status"`
	MediaDescription string   `json:"media_description"`
}

func (account *Account) UpdateLiveMedia(channel string, changes Livestream) (*LiveMedia, error) {
	var args f.Args
	args.Add("authToken", account.AuthToken)

	var upd = struct {
		Livestream []update `json:"livestream"`
	}{}
	upd.Livestream = append(upd.Livestream, update{
		MediaUserName:    changes.MediaName,
		MediaID:          changes.MediaID,
		MediaCategoryID:  changes.MediaCategoryID,
		MediaLiveDelay:   changes.MediaLiveDelay,
		MediaRecording:   changes.MediaRecording,
		MediaMature:      changes.MediaMature,
		MediaStatus:      changes.MediaStatus,
		MediaDescription: changes.MediaDescription,
		MediaHidden:      changes.MediaHidden,
	})
	upd.Livestream[0].MediaCountries = append(upd.Livestream[0].MediaCountries, changes.MediaCountries[0])
	if changes.MediaHidden == "" {
		upd.Livestream[0].MediaHidden = "0"
	}
	if len(changes.MediaHostedMedia.Livestream) != 0 {
		upd.Livestream[0].MediaHostedName = changes.MediaHostedMedia.Livestream[0].MediaUserName
	} else {
		upd.Livestream[0].MediaHostedName = "off"
	}

	dst, err := json.Marshal(&upd)
	if err != nil {
		return nil, err
	}

	log.Printf("%#v", string(dst))

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("media/live/", channel))
	resp, err := just.PUT(dst, url, &args)
	if err != nil {
		return nil, err
	}

	var obj LiveMedia
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return &obj, nil
}
