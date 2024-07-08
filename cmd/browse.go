package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"tempvault/util"
	"tempvault/vault"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(browseCmd)
}

var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "Browse the vault.",
	Run: func(cmd *cobra.Command, args []string) {
		files, err := vault.SelectFilesFromVault()
		if err != nil {
			fmt.Printf("An error occurd browsing files. error: %s\n", err)
		}

		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting cwd, error: ", err)
		}

		vaultDir, err := vault.GetTempVaultDir()
		if err != nil {
			fmt.Println("Error getting vault dir, error: ", err)
		}

		for _, filename := range files {
			src := filepath.Join(vaultDir, filename)
			dst := filepath.Join(cwd, filename)

			var overwriteFile = true
			if util.FileExists(dst) {
				reader := bufio.NewReader(os.Stdin)

				fmt.Printf("File %s already exists in the current working directory. Do you want to overwrite it? (y/n): ", filename)
				response, _ := reader.ReadString('\n')
				response = strings.TrimSpace(response)

				if strings.ToLower(response) != "y" {
					overwriteFile = false
				}

			}

			if overwriteFile {
				util.CopyFile(src, dst)
				fmt.Println("Pasted file", filename, "into the current directory.")
			}
		}
	},
}
