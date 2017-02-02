package chat

import (
	"encoding/json"
	"strings"
)

// Login connect client to channel chat.
//
// If you don’t have a login to use, you can login as a guest role with UnknownSoldier and provide null (not a string) as a token. If you ever get back UnknownSoldier when you didn’t send it, that means you provided a incorrect login and should disconnect and try again.
//
// Sending hideBuffered as true will not send back a backlog of chat messages.
//
// The client must wait until it gets login confirmation before doing anything else.
//
// If you don’t get a loginMsg response within 10 seconds the chat server is not responding, you should disconnect from the server and try another one.
//
// Once you’ve joined a channel, you may be sent a backlog of chatLog and chatMsg, these will contain a buffer boolean property. The last message of each backlog will include a buffersent boolean property to indicate sending has finished.
func (ws *Connection) Login(channel, name, token string, hideBuffered bool) error {
	var joinChannel Message
	joinChannel.Name = message
	joinChannel.Args = append(joinChannel.Args, Args{
		Method: "joinChannel",
		Params: map[string]interface{}{
			"channel":      strings.ToLower(channel),
			"name":         name,
			"token":        token,
			"hideBuffered": hideBuffered,
		},
	})

	body, err := json.Marshal(joinChannel)
	if err != nil {
		return nil
	}

	data := append(msgPrefix, body...)
	return ws.Conn.WriteMessage(textMessage, data)
}
