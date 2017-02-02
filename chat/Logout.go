package chat

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

// Logout send the logout command, you won’t get a response but you will be logged out from the server. You can now close the connection.
//
// You cannot reuse the same connection.
func (ws *Connection) Logout(name string) error {
	var partChannel Message
	partChannel.Name = message
	partChannel.Args = append(partChannel.Args, Args{
		Method: "partChannel",
		Params: map[string]interface{}{
			"name": name,
		},
	})

	body, err := json.Marshal(partChannel)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.WriteMessage(websocket.TextMessage, data)
}
