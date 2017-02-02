package chat

import (
	"encoding/json"
	"strings"
)

func (ws *Connection) CleanupGiveaway(channel string) error {
	var cleanupRaffle Message
	cleanupRaffle.Name = message
	cleanupRaffle.Args = append(cleanupRaffle.Args, Args{
		Method: "cleanupRaffle",
		Params: map[string]interface{}{
			"channel": strings.ToLower(channel),
		},
	})

	body, err := json.Marshal(cleanupRaffle)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
