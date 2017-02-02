package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/websocket"
	just "github.com/toby3d/hitGox/tools"
)

const (
	// APIEndpoint is a URL for all API requests.
	APIEndpoint = "https://api.hitbox.tv/%s"

	// ImagesEndpoint is a URL for all images resources.
	ImagesEndpoint = "https://edge.sf.hitbox.tv"

	message     = "message"
	textMessage = websocket.TextMessage
)

var (
	connPrefix = []byte("1::")
	echoPrefix = []byte("2::")
	msgPrefix  = []byte("5:::")
)

type (
	status struct {
		Success bool
		Error   bool
		Message string
	}

	choice struct {
		Text  string `json:"text"`
		Votes int    `json:"votes"`
	}

	// Message just a response message.
	Message struct {
		Name string `json:"name"`
		Args []Args `json:"args"`
	}

	// Args just args of response Message.
	Args struct {
		Method string `json:"method"`
		Params Params `json:"params"`
	}

	// Params just a params of Message.
	Params map[string]interface{}

	response struct {
		Name string   `json:"name"`
		Args []string `json:"args"`
	}

	Connection struct {
		*websocket.Conn
	}
)

// Connect try opened connection from client to chat server.
func Connect(serverIP string) (*Connection, *http.Response, error) {
	token, err := just.GET(fmt.Sprintf("https://%s/socket.io/1/", serverIP), nil)
	if err != nil {
		return nil, nil, err
	}

	wss := url.URL{
		Scheme: "wss",
		Host:   serverIP,
		Path:   fmt.Sprint("/socket.io/1/websocket/", strings.Split(string(token), ":")[0]),
	}

	var conn Connection
	dial, resp, err := websocket.DefaultDialer.Dial(wss.String(), nil)
	conn.Conn = dial
	return &conn, resp, err
}

// Read allow read any response messages and save opened connection.
func (ws *Connection) Read() (*Message, error) {
	_, source, err := ws.ReadMessage()
	if err != nil {
		return nil, err
	}

	switch {
	case bytes.HasPrefix(source, connPrefix):
		return nil, nil
	case bytes.HasPrefix(source, echoPrefix):
		return nil, ws.WriteMessage(websocket.TextMessage, source)
	case bytes.HasPrefix(source, msgPrefix):
		source = bytes.TrimPrefix(source, msgPrefix)
		var data response
		if err := json.Unmarshal(source, &data); err != nil {
			return nil, err
		}

		var args Args
		if err := json.Unmarshal([]byte(data.Args[0]), &args); err != nil {
			return nil, err
		}

		var msg Message
		msg.Name = data.Name
		msg.Args = append(msg.Args, args)

		return &msg, nil
	default:
		return nil, ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	}
}

/*
func checkArgs(args *Args) interface{} {
	switch args.Method {
	case "banList":
	case "chatLog":
	case "chatMsg":
	case "directMsg":
	case "infoMsg":
	case "loginMsg":
	case "mediaLog":
	case "motdMsg":
	case "pollMsg":
	case "raffleMsg":
	case "serverMsg":
	case "slowMsg":
	case "startPoll":
	case "userInfo":
	case "userList":
	case "winnerRaffle":
	default:
		return nil
	}
}
*/
