package chat

import (
	"encoding/json"
	"strings"
	"time"
)

func (ws *Connection) StickyMessage(channel, name, nameColor, text string, startTime *time.Time) error {
	var motdMsg Message
	motdMsg.Name = message
	motdMsg.Args = append(motdMsg.Args, Args{
		Method: "motdMsg",
		Params: map[string]interface{}{
			"channel":   strings.ToLower(channel),
			"name":      name,
			"nameColor": nameColor,
			"text":      text,
			"time":      startTime.Format(time.RFC3339Nano),
		},
	})

	body, err := json.Marshal(motdMsg)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
