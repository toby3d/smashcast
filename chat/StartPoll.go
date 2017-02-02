package chat

import (
	"encoding/json"
	"strings"
	"time"
)

func (ws *Connection) StartPoll(channel, nameColor, question string, choices []string, subscriberOnly, followerOnly bool, startTime *time.Time) error {
	poll := make([]choice, len(choices))
	for _, text := range choices {
		poll = append(poll, choice{text, 0})
	}

	var createPoll Message
	createPoll.Name = message
	createPoll.Args = append(createPoll.Args, Args{
		Method: "createPoll",
		Params: map[string]interface{}{
			"channel":        strings.ToLower(channel),
			"question":       question,
			"choices":        poll,
			"subscriberOnly": subscriberOnly,
			"followerOnly":   followerOnly,
			"start_time":     startTime.Format(time.RFC3339Nano),
			"nameColor":      nameColor,
		},
	})

	body, err := json.Marshal(createPoll)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
