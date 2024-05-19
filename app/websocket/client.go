package websocket

import (
	"bytes"
	"log"
	"time"

	"github.com/gorilla/websocket"
	gorillaWS "github.com/gorilla/websocket"
)

type Client struct {
	hub  *Hub
	conn *gorillaWS.Conn
	send chan []byte
}

var (
	maxMessageSize = int64(512)
	pongWait       = 60 * time.Second
)

func (c *Client) readPump() {

	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, gorillaWS.CloseGoingAway, gorillaWS.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(message)

		c.hub.broadcast <- &Message{Data: message, Client: c}
	}
}

func (c *Client) writePump() {

	ticker := time.NewTicker(10 * time.Second)

	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(gorillaWS.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(gorillaWS.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.conn.WriteMessage(gorillaWS.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) Register() {

	c.hub.register <- c

	go c.writePump()

	go c.readPump()
}

func NewClient(hub *Hub, conn *gorillaWS.Conn) *Client {

	return &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, 256),
	}
}
