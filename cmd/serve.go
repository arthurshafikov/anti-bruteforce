package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thewolf27/anti-bruteforce/internal/launcher"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:     "serve",
	Aliases: []string{"s"},
	Short:   "serve command runs HTTP/GRPC servers",
	Long:    "serve command initialize app and runs HTTP/GRPC servers",
	Run: func(cmd *cobra.Command, args []string) {
		launcher.Run(config)
	},
}
