package chat

import (
	"encoding/json"
	"strings"
)

// MediaLog get a list of images posted in a channel.
//
// mType must be "image" or "link".
func (ws *Connection) MediaLog(channel, name, token string) error {
	var mediaLog Message
	mediaLog.Name = message
	mediaLog.Args = append(mediaLog.Args, Args{
		Method: "mediaLog",
		Params: map[string]interface{}{
			"channel": strings.ToLower(channel),
			"type":    "image",
			"name":    name,
			"token":   token,
		},
	})

	body, err := json.Marshal(mediaLog)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
