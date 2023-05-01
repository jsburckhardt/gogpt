package gpt

import (
	"context"
	"gogpt/pkg/adapter"
)

// create a map of single word against a paragraph
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
