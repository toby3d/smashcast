package chat

import (
	"encoding/json"
	"strings"
)

func (ws *Connection) DirectMessage(channel, from, to, nameColor, text string) error {
	var directMsg Message
	directMsg.Name = message
	directMsg.Args = append(directMsg.Args, Args{
		Method: "directMsg",
		Params: map[string]interface{}{
			"channel":   strings.ToLower(channel),
			"from":      from,
			"to":        to,
			"nameColor": nameColor,
			"text":      text,
		},
	})

	body, err := json.Marshal(directMsg)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
