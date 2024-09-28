package socket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Client struct {
	ID     string
	GameID string
	Conn   *websocket.Conn
	Pool   *Pool
}

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return conn, nil
}

func (c *Client) HandleConn() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		log.Println("sending message....")
		c.Pool.Broadcast <- "hello world"
		time.Sleep(time.Second * 5)
	}
}
