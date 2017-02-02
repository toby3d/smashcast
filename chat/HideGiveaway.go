package chat

import (
	"encoding/json"
	"strings"
)

func (ws *Connection) HideGiveaway(channel string) error {
	var hideRaffle Message
	hideRaffle.Name = message
	hideRaffle.Args = append(hideRaffle.Args, Args{
		Method: "hideRaffle",
		Params: map[string]interface{}{
			"channel": strings.ToLower(channel),
		},
	})

	body, err := json.Marshal(hideRaffle)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
