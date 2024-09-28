package ai

import (
	"bytes"
	"context"
	"html/template"
	"strconv"

	openai "github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

var promptQA = `You are an experienced games designer, known for creativity.

# Task
You will get category name and your job is to respond with {{.QuestionsCount}} questions related to that category and each question should have 4 predefined responses and indicator which response is correct.
The questions shouldn't be too easy, and some of them can be tricky.
Use European standard when giving measurements.

Category: {{.Category}}
`

func GetCategoryQnA(client openai.Client, category string, questionsCount int) (string, error) {
	tmpl, err := template.New("promptQA").Parse(promptQA)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, map[string]string{"Category": category, "QuestionsCount": strconv.Itoa(questionsCount)}); err != nil {
		return "", err
	}

	// Structured output JSON schema definition
	type Response struct {
		Questions []struct {
			Question           string   `json:"question"`
			Answers            []string `json:"answers"`
			CorrectAnswerIndex int      `json:"correctAnswerIndex"`
		} `json:"questions"`
	}
	var response Response
	schema, err := jsonschema.GenerateSchemaForType(response)
	if err != nil {
		return "", err
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT4o20240806,
			MaxTokens:   2000,
			Temperature: 1.0,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: buf.String(),
				},
			},
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
				JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
					Name:   "quiz_question",
					Schema: schema,
					Strict: true,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
