package chat

import (
	"encoding/json"
	"strings"
)

func (ws *Connection) RemoveModerator(channel, name, token string) error {
	var removeMod Message
	removeMod.Name = message
	removeMod.Args = append(removeMod.Args, Args{
		Method: "removeMod",
		Params: map[string]interface{}{
			"channel": strings.ToLower(channel),
			"name":    name,
			"token":   token,
		},
	})

	body, err := json.Marshal(removeMod)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
