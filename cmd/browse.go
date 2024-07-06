package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(browseCmd)
}

var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "Browse the vault.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Browse tempvault!")
	},
}
