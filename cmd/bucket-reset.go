package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(bucketResetCmd)
}

var bucketResetCmd = &cobra.Command{
	Use:     "bucket-reset",
	Aliases: []string{"br"},
	Short:   "bucket-reset resets login/password/ip bucket",
	Long:    "bucket-reset resets login/password/ip bucket",
	Run: func(cmd *cobra.Command, args []string) {
		err := app.ResetBucket()
		if err != nil {
			echoErrorAndExit(err)
		}

		fmt.Println("Bucket was resetted successfully!")
	},
}
