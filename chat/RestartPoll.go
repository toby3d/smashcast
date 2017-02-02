package chat

import (
	"encoding/json"
	"strings"
)

func (ws *Connection) RestartPoll(channel, token string) error {
	var startPoll Message
	startPoll.Name = message
	startPoll.Args = append(startPoll.Args, Args{
		Method: "startPoll",
		Params: map[string]interface{}{
			"channel": strings.ToLower(channel),
			"token":   token,
		},
	})

	body, err := json.Marshal(startPoll)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
