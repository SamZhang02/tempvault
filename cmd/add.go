package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a file to the vault.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add file to tempvault!")
	},
}
