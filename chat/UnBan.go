package chat

import (
	"encoding/json"
	"strings"
)

func (ws *Connection) UnBan(channel, name, token string) error {
	var unbanUser Message
	unbanUser.Name = message
	unbanUser.Args = append(unbanUser.Args, Args{
		Method: "unbanUser",
		Params: map[string]interface{}{
			"channel": strings.ToLower(channel),
			"name":    name,
			"token":   token,
		},
	})

	body, err := json.Marshal(unbanUser)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
