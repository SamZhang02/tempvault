package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"tempvault/config"

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
			if !fileExists(fileName) {
				fmt.Printf("File does not exist, file: %s\n", fileName)
			}

			putFileInVault(fileName)
		}
	},
}

func fileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func putFileInVault(src string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	vaultPath, err := config.GetTempVaultDir()
	if err != nil {
		return err
	}

	filename := filepath.Base(src)
	destinationPath := filepath.Join(vaultPath, filename)

	var overwriteFile = true
	if fileExists(destinationPath) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Printf("File %s already exists in the vault. Do you want to overwrite it? (y/n): ", filename)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(response)

		if strings.ToLower(response) != "y" {
			overwriteFile = false
		}
	}

	if overwriteFile {
		destinationFile, err := os.Create(destinationPath)
		if err != nil {
			return err
		}
		defer destinationFile.Close()
	}

	return nil
}
