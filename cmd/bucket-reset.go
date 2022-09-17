package cmd

import (
	"fmt"
	"log"

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
		err := appCli.ResetBucket()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Bucket was resetted successfully!")
	},
}
