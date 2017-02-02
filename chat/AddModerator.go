package chat

import (
	"encoding/json"
	"strings"
)

func (ws *Connection) AddModerator(channel, name, token string) error {
	var makeMod Message
	makeMod.Name = message
	makeMod.Args = append(makeMod.Args, Args{
		Method: "makeMod",
		Params: map[string]interface{}{
			"channel": strings.ToLower(channel),
			"name":    name,
			"token":   token,
		},
	})

	body, err := json.Marshal(makeMod)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
