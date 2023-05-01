// Package sh : provides sh command
package sh

import (
	"gogpt/internal/logger"
	"gogpt/pkg/gpt"

	"github.com/spf13/cobra"
)

// shCmd represents the sh command
var shCmd = &cobra.Command{
	Use:   "sh <prompt>",
	Short: "sh <prompt>",
	Long: `a shell request. Typical use for retrieving a shell script:

gogpt sh "how to search all files in bash for a filename?"`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.SetVerbose(cmd)
		log := logger.GetInstance()
		completion := gpt.NewCompletionService(log)
		err := completion.GetChatCompletion(args[0], "sh")
		if err != nil {
			log.Errorf("Unable ask endpoint: %v", err)
		}
	},
}

// NewCmdShRun returns the sh command
func NewCmdShRun() *cobra.Command {
	return shCmd
}
