package chat

import (
	"encoding/json"
	"strings"
	"time"
)

func (ws *Connection) CreateGiveaway(channel, nameColor, question, prize string, choices []string, subscriberOnly, followerOnly bool, startTime *time.Time) error {
	poll := make([]choice, len(choices))
	for _, text := range choices {
		poll = append(poll, choice{text, 0})
	}

	var createRaffle Message
	createRaffle.Name = message
	createRaffle.Args = append(createRaffle.Args, Args{
		Method: "createRaffle",
		Params: map[string]interface{}{
			"channel":        strings.ToLower(channel),
			"question":       question,
			"prize":          prize,
			"choices":        poll,
			"subscriberOnly": subscriberOnly,
			"followerOnly":   followerOnly,
			"start_time":     startTime.Format(time.RFC3339Nano),
			"nameColor":      nameColor,
		},
	})

	body, err := json.Marshal(createRaffle)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
