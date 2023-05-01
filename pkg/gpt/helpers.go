package gpt

import (
	"context"
	"gogpt/pkg/adapter"
)

// ListModels prints the list of available models
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
