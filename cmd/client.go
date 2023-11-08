package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type Task struct {
	From string `json:"from"`
	To   string `json:"to"`
	Msg  string `json:"msg"` // << used to be byte
}

type Client struct {
	user    string
	ws      *websocket.Conn
	server  *Server
	message chan *Task
}

func (c *Client) listen() {
	defer func() {
		c.server.unregister <- c
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		log.Println("------ client.listen ------")
		task := &Task{}
		err := c.ws.ReadJSON(task)
		log.Printf("To: %s, From: %s, Msg: %s", task.To, task.From, task.Msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		c.server.tasks <- task
	}
}

func (c *Client) write() {
	defer func() {
		c.ws.Close()
	}()
	for {
		select {
		case task, ok := <-c.message:
			log.Println("------ client.write ------")
			c.ws.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				log.Println("client.write channel closed")
				c.ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if task == nil {
				return
			}
			c.ws.WriteJSON(task)
		}
	}
}
