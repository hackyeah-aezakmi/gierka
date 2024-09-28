package main

import (
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
}
