package chat

import (
	"encoding/json"
	"strconv"
	"strings"
)

func (ws *Connection) Timeout(channel, name, token string, timeout int) error {
	var kickUser Message
	kickUser.Name = message
	kickUser.Args = append(kickUser.Args, Args{
		Method: "kickUser",
		Params: map[string]interface{}{
			"channel": strings.ToLower(channel),
			"name":    name,
			"token":   token,
			"timeout": strconv.Itoa(timeout),
		},
	})

	body, err := json.Marshal(kickUser)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
