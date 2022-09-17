package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(whitelist)
	whitelist.AddCommand(whitelistAdd, whitelistRemove)
}

var whitelist = &cobra.Command{
	Use:     "whitelist",
	Aliases: []string{"w"},
	Short:   "whitelist command allows you to add and remove subnets from whitelist",
	Long:    "whitelist command allows you to add and remove subnets from whitelist",
}

var whitelistAdd = &cobra.Command{
	Use:   "add",
	Short: "add subnet",
	Long:  "add subnet to the table",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := appCli.AddToWhitelist(args[0])
		if err != nil {
			echoErrorAndExit(err)
		}

		fmt.Printf("Subnet %s was added to the whitelist\n", args[0])
	},
}

var whitelistRemove = &cobra.Command{
	Use:   "rm",
	Short: "remove subnet",
	Long:  "remove subnet from the table",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := appCli.RemoveFromWhitelist(args[0])
		if err != nil {
			echoErrorAndExit(err)
		}

		fmt.Printf("Subnet %s was removed from the whitelist\n", args[0])
	},
}
