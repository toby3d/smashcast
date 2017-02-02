package chat

import (
	"encoding/json"
	"strings"
)

func (ws *Connection) PauseGiveaway(channel string) error {
	var pauseRaffle Message
	pauseRaffle.Name = message
	pauseRaffle.Args = append(pauseRaffle.Args, Args{
		Method: "pauseRaffle",
		Params: map[string]interface{}{
			"channel": strings.ToLower(channel),
		},
	})

	body, err := json.Marshal(pauseRaffle)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
