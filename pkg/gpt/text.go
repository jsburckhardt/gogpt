package gpt

import (
	"context"
	"fmt"

	markdown "github.com/MichaelMure/go-term-markdown"
	openai "github.com/sashabaranov/go-openai"
)

func (s *Service) GetCompletion(prompt string) error {
	client := newClient()
	model, err := getModel()
	if err != nil {
		return err
	}
	resp, err := client.CreateCompletion(
		context.Background(),
		openai.CompletionRequest{
			Model:  model,
			Prompt: prompt,
		},
	)
	if err != nil {
		return err
	}

	source := resp.Choices[0].Text
	result := markdown.Render(source, 120, 6)
	fmt.Println(string(result))
	return nil
}
