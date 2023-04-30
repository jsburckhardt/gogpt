/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package code

import (
	"gogpt/internal/logger"
	"gogpt/pkg/gpt"

	"github.com/spf13/cobra"
)

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code <prompt>",
	Short: "code <prompt>",
	Long: `a coding question for gpt. Typical use of Ada model:

gogpt code "how to search all files in bash for a filename?"`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.SetVerbose(cmd)
		log := logger.GetInstance()
		completion := gpt.NewCompletionService(log)
		err := completion.GetChatCompletion(args[0], "code")
		if err != nil {
			log.Errorf("Unable ask endpoint: %v", err)
		}
	},
}

func NewCmdCodeRun() *cobra.Command {
	return codeCmd
}
