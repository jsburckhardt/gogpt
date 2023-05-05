// Package gpt provides chat command
package gpt

import (
	"context"
	"fmt"

	markdown "github.com/MichaelMure/go-term-markdown"
	openai "github.com/sashabaranov/go-openai"
)

// GetChatCompletion returns the chat completion
func (s *Service) GetChatCompletion(prompt string, system string) error {
	client := newClient()
	model, err := getModel()
	if err != nil {
		return err
	}
	switch system {
	case "sh":
		prompt = GenerateShellPrompt(prompt)
	default:
		prompt = GeneratePrompt(system, prompt)
	}
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return err
	}

	source := resp.Choices[0].Message.Content
	if system == "sh" {
		result := source
		fmt.Println(result)
		return nil
	}

	result := markdown.Render(source, 80, 0)
	fmt.Println(string(result))
	return nil
}
