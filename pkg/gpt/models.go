package gpt

import (
	"gogpt/internal/gogptconfig"

	openai "github.com/sashabaranov/go-openai"
)

// AllowedModels is a list of allowed models
var AllowedModels = []string{
	openai.GPT3Ada,
	openai.GPT3Babbage,
	openai.GPT3Curie,
	openai.GPT3CurieInstructBeta,
	openai.GPT3Davinci,
	openai.GPT3DavinciInstructBeta,
	openai.GPT3Dot5Turbo,
	openai.GPT3Dot5Turbo0301,
	openai.GPT3TextAda001,
	openai.GPT3TextBabbage001,
	openai.GPT3TextCurie001,
	openai.GPT3TextDavinci001,
	openai.GPT3TextDavinci002,
	openai.GPT3TextDavinci003,
	openai.GPT4,
	openai.GPT40314,
	openai.GPT432K,
}

// get model from string in OPENAI_API_MODEL config
func getModel() (string, error) {
	modelName := gogptconfig.GetConfig().GetString("OPENAI_API_MODEL")
	for _, model := range AllowedModels {
		if model == modelName {
			return model, nil
		}
	}
	// TODO: remove line once pr is merged
	return "openai.GPT3Dot5Turbo0301", nil
	// return "", fmt.Errorf("invalid model name: %s", modelName)
}
