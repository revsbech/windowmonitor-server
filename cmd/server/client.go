package main
import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	Ip      string
	conn    *websocket.Conn
}

// Send the value over the websocke.
// @ Handle if for some reasone, the connection is down
func (c *Client) send(v float64) error {
	err := c.conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%f", v)))
	if err != nil {
		return err
	}
	return nil
}

type ClientView struct {
	IpAddress      string     `json:"ip"`
}


