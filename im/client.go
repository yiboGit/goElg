package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 10 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 2048

	// staff id sync
	staffPrefix    = "$staff"
	staffPrefixLen = len(staffPrefix)
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
	roles   = map[string]bool{"$staff": true}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	ex *Exchanger

	// The websocket connection.
	conn  *websocket.Conn
	Role  string
	ID    int32
	Appid string
	meta  map[string]interface{}
	// Buffered channel of outbound messages.
	send chan []byte
	cc   *ChatContext
}

// BindChatContext  绑定cc
func (c *Client) BindChatContext(cc *ChatContext) {
	c.cc = cc
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.ex.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			return
		}
		if c.ID == 0 {
			error := c.bindRoleID(message)
			if error != nil {
				log.Printf("error bind role id: %v, message: %v\n", error, string(message))
				return
			}
			c.ex.bindStaff <- c
		} else {
			if string(message[:6]) != `"ping"` {
				c.ex.messages <- Message{c, message}
			}
		}
	}
}
func (c *Client) bindRoleID(message []byte) error {
	// $staff:12:appid
	error := json.Unmarshal(message, &c.meta)
	if error != nil {
		log.Printf("json parse error: %v \n", string(message))
		return errors.New("meta json parse fail")
	}
	ID, exist := c.meta["id"]
	if !exist {
		return errors.New("role id error, close conn")
	}
	c.ID = int32(ID.(float64))
	appid, exist := c.meta["appid"].(string)
	if !exist {
		return errors.New("appid not found close conn")
	}
	c.Appid = appid
	return nil
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			err := c.conn.WriteMessage(websocket.TextMessage, message)

			if err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
