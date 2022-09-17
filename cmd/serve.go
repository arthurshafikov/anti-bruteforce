package cmd

import (
	"github.com/arthurshafikov/anti-bruteforce/internal/app"
	"github.com/spf13/cobra"
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
		app.Run(config)
	},
}
