// Package cmd contains the commands for the gogpt tool.
package cmd

import (
	"fmt"
	"gogpt/cmd/chat"
	"gogpt/cmd/code"
	"gogpt/cmd/model"
	"gogpt/cmd/sh"
	"gogpt/cmd/text"
	"gogpt/internal/logger"
	"os"

	"github.com/spf13/cobra"
)

var (
	hash    string
	verbose bool

	rootCmd = &cobra.Command{
		Use:   "gogpt",
		Short: "gogpt",
		Long: `



	██████╗  ██████╗  ██████╗ ██████╗ ████████╗
	██╔════╝ ██╔═══██╗██╔════╝ ██╔══██╗╚══██╔══╝
	██║  ███╗██║   ██║██║  ███╗██████╔╝   ██║
	██║   ██║██║   ██║██║   ██║██╔═══╝    ██║
	╚██████╔╝╚██████╔╝╚██████╔╝██║        ██║
		╚═════╝  ╚═════╝  ╚═════╝ ╚═╝        ╚═╝
												`,
	}
)

// Execute adds all child commands to the root command.
func Execute(version, commit string) {
	rootCmd.Version = version
	hash = commit

	setVersion()

	if err := rootCmd.Execute(); err != nil {
		logger.GetInstance().Error(err)
		os.Exit(1)
	}
}

func setVersion() {
	template := fmt.Sprintf("gogpt version: %s commit: %s \n", rootCmd.Version, hash)
	rootCmd.SetVersionTemplate(template)
}

func init() {
	cobra.OnInitialize()
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "set logging level to verbose")
	rootCmd.AddCommand(chat.NewCmdChatRun())
	rootCmd.AddCommand(text.NewCmdTextRun())
	rootCmd.AddCommand(code.NewCmdCodeRun())
	rootCmd.AddCommand(model.NewCmdModelRun())
	rootCmd.AddCommand(sh.NewCmdShRun())
}
