package chat

import (
	"encoding/json"
	"strings"
)

func (ws *Connection) Ban(channel, name, token string, banIP bool) error {
	var banUser Message
	banUser.Name = message
	banUser.Args = append(banUser.Args, Args{
		Method: "banUser",
		Params: map[string]interface{}{
			"channel": strings.ToLower(channel),
			"name":    name,
			"token":   token,
			"banIP":   banIP,
		},
	})

	body, err := json.Marshal(banUser)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
