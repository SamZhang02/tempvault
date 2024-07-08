package cmd

import (
	"fmt"
	"os"

	"tempvault/util"
	"tempvault/vault"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [files]",
	Short: "Add a file to the vault.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No files provided")
			cmd.Usage()
			os.Exit(1)
		}

		for _, fileName := range args {
			if !util.FileExists(fileName) {
				fmt.Printf("File does not exist, file: %s\n", fileName)
			}

			vault.PutFileInVault(fileName)
		}
	},
}
