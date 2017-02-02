package chat

import (
	"encoding/json"
	"strings"
)

func (ws *Connection) SubscriberOnlyMode(channel string, subscriber bool) error {
	var slowMode Message
	slowMode.Name = message
	slowMode.Args = append(slowMode.Args, Args{
		Method: "slowMode",
		Params: map[string]interface{}{
			"channel":    strings.ToLower(channel),
			"subscriber": subscriber,
			"rate":       0,
		},
	})

	body, err := json.Marshal(slowMode)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
