package chat

import (
	"encoding/json"
	"strings"
)

func (ws *Connection) UserInfo(channel, name string) error {
	var getChannelUser Message
	getChannelUser.Name = message
	getChannelUser.Args = append(getChannelUser.Args, Args{
		Method: "getChannelUser",
		Params: map[string]interface{}{
			"channel": strings.ToLower(channel),
			"name":    name,
		},
	})

	body, err := json.Marshal(getChannelUser)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
