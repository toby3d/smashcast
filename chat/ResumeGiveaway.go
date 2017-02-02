package chat

import (
	"encoding/json"
	"strings"
)

func (ws *Connection) ResumeGiveaway(channel string) error {
	var startRaffle Message
	startRaffle.Name = message
	startRaffle.Args = append(startRaffle.Args, Args{
		Method: "startRaffle",
		Params: map[string]interface{}{
			"channel": strings.ToLower(channel),
		},
	})

	body, err := json.Marshal(startRaffle)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
