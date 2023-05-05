// Package sh : provides sh command
package sh

import (
	"bufio"
	"fmt"
	"gogpt/internal/logger"
	"gogpt/pkg/gpt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

// shCmd represents the sh command
var shCmd = &cobra.Command{
	Use:   "sh <prompt>",
	Short: "sh <prompt>",
	Long: `a shell request. Typical use for retrieving a shell script:

gogpt sh "how to search all files in bash for a filename?"
git diff | gogpt sh "generage a commit meesage"`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.SetVerbose(cmd)
		log := logger.GetInstance()
		completion := gpt.NewCompletionService(log)

		var prompt string
		if len(args) > 0 {
			prompt = args[0]
			// check if something in Stdin
			stdinPassed := !terminal.IsTerminal(int(os.Stdin.Fd()))

			if stdinPassed {
				bytes, err := ioutil.ReadAll(os.Stdin)
				if err != nil {
					log.Errorf("can not read from stdin %v", err)
				}
				prompt = fmt.Sprintf("%s %s", prompt, bytes)
			}

		} else {
			reader := bufio.NewReader(os.Stdin)
			promptBytes, _, err := reader.ReadLine()
			if err != nil {
				log.Errorf("Unable to read from standard input: %v", err)
				return
			}
			prompt = string(promptBytes)
		}

		prompt = strings.TrimSpace(prompt)
		if prompt == "" {
			log.Error("Prompt cannot be empty")
			return
		}

		log.Debugf("prompt %s", prompt)
		err := completion.GetChatCompletion(prompt, "sh")
		if err != nil {
			log.Errorf("Unable to ask endpoint: %v", err)
		}
	},
}

// NewCmdShRun returns the sh command
func NewCmdShRun() *cobra.Command {
	return shCmd
}
