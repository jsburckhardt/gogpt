// Package chat : provides chat command
package chat

import (
	"gogpt/internal/logger"
	"gogpt/pkg/gpt"

	"github.com/spf13/cobra"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat <prompt>",
	Short: "chat <prompt>",
	Long: `a prompt you ask gpt for. Typical use of GPT model:

gogpt chat "what's the capital of France?",
gogpt chat "fizz buzz in bash"`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.SetVerbose(cmd)
		log := logger.GetInstance()
		completion := gpt.NewCompletionService(log)
		err := completion.GetChatCompletion(args[0], "chat")
		if err != nil {
			log.Errorf("Unable ask endpoint: %v", err)
			return err
		}
		return nil
	},
}

// NewCmdChatRun returns the chat command
func NewCmdChatRun() *cobra.Command {
	return chatCmd
}
