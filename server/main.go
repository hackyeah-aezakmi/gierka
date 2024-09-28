package main

import (
	"log"
	"os"

	"github.com/hackyeah-aezakmi/gierka/transport/http"
	"github.com/hackyeah-aezakmi/gierka/transport/socket"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
