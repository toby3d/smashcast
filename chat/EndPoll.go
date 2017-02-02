package chat

import (
	"encoding/json"
	"strings"
)

func (ws *Connection) EndPoll(channel, token string) error {
	var endPoll Message
	endPoll.Name = message
	endPoll.Args = append(endPoll.Args, Args{
		Method: "endPoll",
		Params: map[string]interface{}{
			"channel": strings.ToLower(channel),
			"token":   token,
		},
	})

	body, err := json.Marshal(endPoll)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
