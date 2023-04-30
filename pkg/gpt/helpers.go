package gpt

import (
	"context"
	"gogpt/pkg/adapter"
)

// create a map of single word against a paragraph

var assistantInstructions = map[string]string{
	"chat": "You are a friendly assistant called gogpt. Having a nice chat with you.",
	"code": "You are coding assistant called gogpt, you get only questions and only reply code blocks.",
}

func getAssistantInstructions(key string) string {
	return assistantInstructions[key]
}

func ListModels(log *adapter.Logger) error {
	client := newClient()
	listModels, err := client.ListModels(context.Background())
	if err != nil {
		return err
	}
	for _, model := range listModels.Models {
		log.Infof(model.ID)
	}
	return nil
}
