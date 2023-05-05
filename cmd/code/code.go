// Package code : provides code command
package code

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
				prompt = fmt.Sprintf("%s\n\n%s", bytes, prompt)
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

		err := completion.GetChatCompletion(prompt, "code")
		if err != nil {
			log.Errorf("Unable ask endpoint: %v", err)
		}
	},
}

// NewCmdCodeRun returns the code command
func NewCmdCodeRun() *cobra.Command {
	return codeCmd
}
