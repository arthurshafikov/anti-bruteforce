package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(blacklist)
	blacklist.AddCommand(blacklistAdd, blacklistRemove)
}

var blacklist = &cobra.Command{
	Use:     "blacklist",
	Aliases: []string{"b"},
	Short:   "blacklist command allows you to add and remove subnets from blacklist",
	Long:    "blacklist command allows you to add and remove subnets from blacklist",
}

var blacklistAdd = &cobra.Command{
	Use:   "add",
	Short: "add subnet",
	Long:  "add subnet to the table",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := appCli.AddToBlacklist(args[0])
		if err != nil {
			echoErrorAndExit(err)
		}

		fmt.Printf("Subnet %s was added to the blacklist\n", args[0])
	},
}

var blacklistRemove = &cobra.Command{
	Use:   "rm",
	Short: "remove subnet",
	Long:  "remove subnet from the table",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := appCli.RemoveFromBlacklist(args[0])
		if err != nil {
			echoErrorAndExit(err)
		}

		fmt.Printf("Subnet %s was removed from the blacklist\n", args[0])
	},
}
