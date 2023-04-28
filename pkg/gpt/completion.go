package gpt

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

func (s *Service) GetCompletion(prompt string) error {
	openaiApiKey := s.config.GetString("OPENAI_API_KEY")
	client := openai.NewClient(openaiApiKey)
	resp, err := client.CreateCompletion(
		context.Background(),
		openai.CompletionRequest{
			Model:  openai.GPT3Ada,
			Prompt: prompt,
		},
	)
	if err != nil {
		return err
	}
	fmt.Println(resp.Choices[0].Text)
	return nil
}
