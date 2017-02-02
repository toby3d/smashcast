package chat

import (
	"encoding/json"
	"strings"
)

func (ws *Connection) VoteGiveaway(channel, name string, choice int) error {
	var voteRaffle Message
	voteRaffle.Name = message
	voteRaffle.Args = append(voteRaffle.Args, Args{
		Method: "voteRaffle",
		Params: map[string]interface{}{
			"name":    name,
			"channel": strings.ToLower(channel),
			"choice":  choice - 1,
		},
	})

	body, err := json.Marshal(voteRaffle)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
