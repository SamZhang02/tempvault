package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"tempvault/config"

	fzf "github.com/junegunn/fzf/src"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(browseCmd)
}

var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "Browse the vault.",
	Run: func(cmd *cobra.Command, args []string) {
		vaultDir, err := config.GetTempVaultDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		inputChan := make(chan string)
		go func() {
			err := filepath.Walk(vaultDir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					inputChan <- filepath.Base(path)
				}
				return nil
			})
			if err != nil {
				fmt.Printf("An error occurred while walking through the vault directory, error: %s\n", err)
			}
			close(inputChan)
		}()

		outputChan := make(chan string)
		go func() {
			for s := range outputChan {
				fmt.Println("Got: " + s)
			}
		}()

		exit := func(code int, err error) {
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
			os.Exit(code)
		}

		// Build fzf.Options
		options, err := fzf.ParseOptions(
			true, // whether to load defaults ($FZF_DEFAULT_OPTS_FILE and $FZF_DEFAULT_OPTS)
			[]string{"--preview", "cat " + vaultDir + "/{}", "--multi"},
		)
		if err != nil {
			exit(fzf.ExitError, err)
		}

		// Set up input and output channels
		options.Input = inputChan
		options.Output = outputChan

		// Run fzf
		code, err := fzf.Run(options)
		exit(code, err)
	},
}
