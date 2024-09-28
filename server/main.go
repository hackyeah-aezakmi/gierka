package main

import (
	"fmt"
	"github.com/hackyeah-aezakmi/gierka/transport/http"
	"github.com/hackyeah-aezakmi/gierka/transport/socket"
	"log"
)

func main() {
	fmt.Println("elo")

	pool := socket.NewPool()
	go pool.Start()
	h := http.NewHandler(pool)
	if err := h.Serve(); err != nil {
		log.Fatalf("http.Serve: %s", err)
	}
}
