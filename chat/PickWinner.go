package chat

import (
	"encoding/json"
	"strings"
)

func (ws *Connection) PickWinner(channel string, answer int) error {
	var winnerRaffle Message
	winnerRaffle.Name = message
	winnerRaffle.Args = append(winnerRaffle.Args, Args{
		Method: "winnerRaffle",
		Params: map[string]interface{}{
			"channel": strings.ToLower(channel),
			"answer":  answer - 1,
		},
	})

	body, err := json.Marshal(winnerRaffle)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
