package chat

import (
	"encoding/json"
	"strings"
)

func (ws *Connection) EndGiveaway(channel string) error {
	var endRaffle Message
	endRaffle.Name = message
	endRaffle.Args = append(endRaffle.Args, Args{
		Method: "endRaffle",
		Params: map[string]interface{}{
			"channel": strings.ToLower(channel),
		},
	})

	body, err := json.Marshal(endRaffle)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
