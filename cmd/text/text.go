/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package text

import (
	"gogpt/internal/logger"
	"gogpt/pkg/gpt"

	"github.com/spf13/cobra"
)

// txtCmd represents the txt command
var textCmd = &cobra.Command{
	Use:   `text "<prompt>"`,
	Short: `text "<prompt>"`,
	Long: `a prompt that contains the request for gpt. Typical use of Ada model

gogpt text "what's the capital of France?"`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.SetVerbose(cmd)
		log := logger.GetInstance()
		completion := gpt.NewCompletionService(log)
		err := completion.GetCompletion(args[0])
		if err != nil {
			log.Errorf("Unable ask endpoint: %v", err)
		}
	},
}

func NewCmdTextRun() *cobra.Command {
	return textCmd
}
