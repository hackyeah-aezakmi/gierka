package main

import (
	"fmt"
	"github.com/hackyeah-aezakmi/gierka/transport/http"
	"github.com/hackyeah-aezakmi/gierka/transport/socket"
	"log"
	"os"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	openaiApiKey := os.Getenv("OPENAI_API_KEY")
	if openaiApiKey == "" {
		log.Fatal("Environment variable OPENAI_API_KEY is required")
	}
	_ = openai.NewClient(openaiApiKey) // OpenAI Client

	pool := socket.NewPool()
	go pool.Start()
	h := http.NewHandler(pool)
	if err := h.Serve(); err != nil {
		log.Fatalf("http.Serve: %s", err)
	}
}