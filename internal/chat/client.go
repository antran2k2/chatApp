package chat

import (
	"github.com/gorilla/websocket"
	"log"
)

// Client đại diện cho mỗi người dùng kết nối qua WebSocket
type Client struct {
	conn *websocket.Conn
	send chan Message
}

// ReadPump đọc tin nhắn từ client và gửi đến hub
func (c *Client) ReadPump() {
	defer func() {
		hub.Unregister <- c
		c.conn.Close()
	}()
	for {
		var msg Message
		if err := c.conn.ReadJSON(&msg); err != nil {
			log.Println("Error reading JSON:", err)
			break
		}
		hub.Broadcast <- msg
	}
}

// WritePump gửi tin nhắn tới client
func (c *Client) WritePump() {
	defer c.conn.Close()
	for msg := range c.send {
		if err := c.conn.WriteJSON(msg); err != nil {
			log.Println("Error writing JSON:", err)
			break
		}
	}
}
