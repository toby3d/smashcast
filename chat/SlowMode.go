package chat

import (
	"encoding/json"
	"strconv"
	"strings"
)

func (ws *Connection) SlowMode(channel string, time int) error {
	var slowMode Message
	slowMode.Name = message
	slowMode.Args = append(slowMode.Args, Args{
		Method: "slowMode",
		Params: map[string]interface{}{
			"channel": strings.ToLower(channel),
			"time":    strconv.Itoa(time),
		},
	})

	body, err := json.Marshal(slowMode)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
