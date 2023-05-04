// Package model : provides model command
package model

import (
	"gogpt/internal/logger"
	"gogpt/pkg/gpt"

	"github.com/spf13/cobra"
)

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "model",
	Long: `a helper to validate models and connection:


gogpt model --list`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.SetVerbose(cmd)
		log := logger.GetInstance()
		log.Infof("model triggered. Select an option")

		if list, _ := cmd.Flags().GetBool("list"); list {
			err := gpt.ListModels(log)
			if err != nil {
				log.Errorf("Unable ask endpoint: %v", err)
			}
			return
		}
	},
}

func init() {
	modelCmd.Flags().BoolP("list", "l", false, "List available model")
}

// NewCmdModelRun returns the model command
func NewCmdModelRun() *cobra.Command {
	return modelCmd
}
