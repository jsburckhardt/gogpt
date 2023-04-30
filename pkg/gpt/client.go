package gpt

import (
	"gogpt/internal/gogptconfig"
	"gogpt/pkg/adapter"

	openai "github.com/sashabaranov/go-openai"
)

// Service interacts with the flight computer and returns stuff
type Service struct {
	logger *adapter.Logger
	config *gogptconfig.Config
}

func NewCompletionService(logger *adapter.Logger) *Service {
	return &Service{
		logger: logger,
		config: gogptconfig.GetConfig(),
	}
}

func newClient() *openai.Client {
	openaiApiKey := gogptconfig.GetConfig().GetString("OPENAI_API_KEY")
	// openaiModel := gogptconfig.GetConfig().GetString("OPENAI_API_MODEL")
	clientType := gogptconfig.GetConfig().GetString("OPENAI_API_TYPE")
	if clientType == "azure" {
		openaiHost := gogptconfig.GetConfig().GetString("OPENAI_API_HOST")
		openaiAzureEngine := gogptconfig.GetConfig().GetString("AZURE_OPENAI_ENGINE")
		config := openai.DefaultAzureConfig(openaiApiKey, openaiHost, openaiAzureEngine)
		return openai.NewClientWithConfig(config)

	}
	return openai.NewClient(openaiApiKey)
}
