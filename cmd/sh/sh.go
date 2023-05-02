// Package sh : provides sh command
package sh

import (
	"gogpt/internal/logger"
	"gogpt/pkg/gpt"
	"io/ioutil"
	"os"

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

		// Check if input is provided via command line arguments
		if len(args) > 0 {
			err := completion.GetChatCompletion(args[0], "sh")
			if err != nil {
				log.Errorf("Unable to ask endpoint: %v", err)
			}
			return
		}

		// Read from stdin
		input, err := os.Stdin.Stat()
		if err != nil {
			log.Errorf("Unable to read from stdin: %v", err)
			return
		}
		if input.Mode()&os.ModeCharDevice != 0 {
			log.Errorf("No input provided")
			return
		}
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Errorf("Unable to read from stdin: %v", err)
			return
		}
		inputStr := string(bytes)
		err = completion.GetChatCompletion(inputStr, "sh")
		if err != nil {
			log.Errorf("Unable to ask endpoint: %v", err)
		}
	},
}

// NewCmdShRun returns the sh command
func NewCmdShRun() *cobra.Command {
	return shCmd
}
