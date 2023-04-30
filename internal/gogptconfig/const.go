// Package gogptconfig is a helper for configs in gogpt
package gogptconfig

// ConfigName is the struct that enforces config names to be set in code
type ConfigName string

//revive:disable
const (
	APP_MODE            ConfigName = "APP_MODE"
	OPENAI_API_KEY      ConfigName = "OPENAI_API_KEY"
	OPENAI_API_MODEL    ConfigName = "OPENAI_API_MODEL"
	OPENAI_API_TYPE     ConfigName = "OPENAI_API_TYPE"
	OPENAI_API_HOST     ConfigName = "OPENAI_API_HOST"
	AZURE_OPENAI_ENGINE ConfigName = "AZURE_OPENAI_ENGINE"
)
