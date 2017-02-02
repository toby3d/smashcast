package chat

import (
	"encoding/json"
	"strings"
)

func (ws *Connection) PausePoll(channel, token string) error {
	var pausePoll Message
	pausePoll.Name = message
	pausePoll.Args = append(pausePoll.Args, Args{
		Method: "pausePoll",
		Params: map[string]interface{}{
			"channel": strings.ToLower(channel),
			"token":   token,
		},
	})

	body, err := json.Marshal(pausePoll)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
