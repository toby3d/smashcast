package chat

import (
	"encoding/json"
	"strings"
)

// SendMessage send a message to the joined channel. The text value is limited to 300 characters.
//
// If accepted the chat server will send it to all users in the channel, including yourself.
//
// If rejected by the chat server due to slow or subscriber only mode, you will get back an infoMsg.
func (ws *Connection) SendMessage(channel, name, nameColor, text string) error {
	var joinChannel Message
	joinChannel.Name = message
	joinChannel.Args = append(joinChannel.Args, Args{
		Method: "chatMsg",
		Params: map[string]interface{}{
			"channel":   strings.ToLower(channel),
			"name":      name,
			"nameColor": nameColor,
			"text":      text,
		},
	})

	body, err := json.Marshal(joinChannel)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
