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
	Use:   "chat",
	Short: "string to ask gpt",
	Long: `a string that contains the request for gpt:

gogpt chat "what's the capital of France?"`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.SetVerbose(cmd)
		log := logger.GetInstance()
		completion := gpt.NewCompletionService(log)
		err := completion.GetChatCompletion(args[0])
		if err != nil {
			log.Errorf("Unable ask gpt: %v", err)
		}
	},
}

func NewCmdChatRun() *cobra.Command {
	return chatCmd
}
