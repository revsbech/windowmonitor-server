package main
import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Client struct {
	Ip      string
	conn    *websocket.Conn
}

// Send the value over the websocke.
// @ Handle if for some reasone, the connection is down
func (c *Client) send(p DataPoint) error {
	bytes, err := json.Marshal(p)
	if err != nil {
		return err
	}

	err = c.conn.WriteMessage(websocket.TextMessage, bytes)
	if err != nil {
		return err
	}
	return nil
}

type ClientView struct {
	IpAddress      string     `json:"ip"`
}


