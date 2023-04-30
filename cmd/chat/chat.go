/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
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
	Run: func(cmd *cobra.Command, args []string) {
		logger.SetVerbose(cmd)
		log := logger.GetInstance()
		completion := gpt.NewCompletionService(log)
		err := completion.GetChatCompletion(args[0], "chat")
		if err != nil {
			log.Errorf("Unable ask endpoint: %v", err)
		}
	},
}

func NewCmdChatRun() *cobra.Command {
	return chatCmd
}
