package gpt

import (
	"context"
	"fmt"

	markdown "github.com/MichaelMure/go-term-markdown"
	openai "github.com/sashabaranov/go-openai"
)

func (s *Service) GetChatCompletion(prompt string, system string) error {
	assistantInstruction := getAssistantInstructions(system)
	client := newClient()
	model, err := getModel()
	if err != nil {
		return err
	}
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: assistantInstruction,
				},
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
	result := markdown.Render(source, 80, 6)
	fmt.Println(string(result))
	return nil
}
