package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tempvault",
	Short: "tempvault is a command line tool to rapidly access and paste your template files.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Browse templates!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
