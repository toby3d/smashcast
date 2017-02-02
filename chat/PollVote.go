package chat

import (
	"encoding/json"
	"strings"
)

func (ws *Connection) PollVote(channel, name, token string, choice int) error {
	var voteMsg Message
	voteMsg.Name = message
	voteMsg.Args = append(voteMsg.Args, Args{
		Method: "voteMsg",
		Params: map[string]interface{}{
			"channel": strings.ToLower(channel),
			"name":    name,
			"choice":  choice - 1,
			"token":   token,
		},
	})

	body, err := json.Marshal(voteMsg)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
