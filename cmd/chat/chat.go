// Package chat : provides chat command
package chat

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
				return err
			}
			prompt = string(promptBytes)
		}

		prompt = strings.TrimSpace(prompt)
		if prompt == "" {
			log.Errorf("Prompt cannot be empty")
			return fmt.Errorf("Prompt cannot be empty")
		}

		log.Debugf("prompt %s", prompt)
		err := completion.GetChatCompletion(prompt, "chat")
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
