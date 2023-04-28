package gpt

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

func (s *Service) GetChatCompletion(prompt string) error {
	client := newClient()
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
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

	fmt.Println(resp.Choices[0].Message.Content)
	return nil
}
