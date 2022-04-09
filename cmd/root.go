package cmd

import (
	"fmt"
	"os"

	"github.com/arthurshafikov/anti-bruteforce/internal/cli"
	configPkg "github.com/arthurshafikov/anti-bruteforce/internal/config"
	"github.com/spf13/cobra"
)

var (
	configFolder string
	config       *configPkg.Config
	app          *cli.AppCli

	rootCmd = &cobra.Command{
		Use:   "anti-bruteforce",
		Short: "anti-bruteforce command line interface",
		Long:  "anti-bruteforce command line interface",
	}
)

func Execute() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig, initGRPCClient)
	rootCmd.PersistentFlags().StringVarP(&configFolder, "configFolder", "c", "./configs", "path to config folder")
}

func initConfig() {
	config = configPkg.NewConfig(configFolder)
}

func initGRPCClient() {
	app = cli.NewAppCli(config.GrpcServerConfig.Address)
}

func echoErrorAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
