package chat

import (
	"encoding/json"
	"strings"
)

func (ws *Connection) UserList(channel string) error {
	var getChannelUserList Message
	getChannelUserList.Name = message
	getChannelUserList.Args = append(getChannelUserList.Args, Args{
		Method: "getChannelUserList",
		Params: map[string]interface{}{
			"channel": strings.ToLower(channel),
		},
	})

	body, err := json.Marshal(getChannelUserList)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
