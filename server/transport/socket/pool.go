package socket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan string
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan string),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("client connected")))
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("client disconnected")))
			}
			break
		case message := <-pool.Broadcast:
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
					log.Printf("pool.Start: message: client.Conn.WriteMessage: %v", err)
					return
				}
			}
		}
	}
}
