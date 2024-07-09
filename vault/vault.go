package vault

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	fzf "github.com/junegunn/fzf/src"
	"tempvault/util"
)

func GetTempVaultDir() (string, error) {
	var homeDir, err = os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, "tempvault"), nil

}

func PutFileInVault(src string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	vaultDir, err := GetTempVaultDir()
	if err != nil {
		return err
	}

	filename := filepath.Base(src)
	dst := filepath.Join(vaultDir, filename)

	var overwriteFile = true
	if util.FileExists(dst) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Printf("File %s already exists in the vault. Do you want to overwrite it? (y/n): ", filename)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(response)

		if strings.ToLower(response) != "y" {
			overwriteFile = false
		}
	}

	var destinationFile *os.File
	if overwriteFile {
		var err error
		destinationFile, err = os.Create(dst)
		if err != nil {
			return err
		}
		defer destinationFile.Close()
	}

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	err = destinationFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

func SelectFilesFromVault() ([]string, error) {
	vaultDir, err := GetTempVaultDir()
	if err != nil {
		return nil, err
	}

	inputChan := make(chan string)
	outputChan := make(chan string)
	var selectedFiles []string

	go func() {
		defer close(inputChan)
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
			fmt.Fprintf(os.Stderr, "An error occurred while walking through the vault directory, error: %s\n", err)
		}
	}()

	go func() {
		for selectedFile := range outputChan {
			selectedFiles = append(selectedFiles, selectedFile)
		}
	}()

	options, err := fzf.ParseOptions(
		true,
		[]string{
			"--preview",
			"cat " + vaultDir + "/{}",
			"--multi",
			"--border",
			"--header=TAB to select an item\nENTER to paste selected items into cwd\nCTRL-c or ESC to quit",
		},
	)
	if err != nil {
		return nil, err
	}

	options.Input = inputChan
	options.Output = outputChan

	code, err := fzf.Run(options)
	if err != nil {
		return nil, err
	}

	if code != 0 && code != 130 {
		return nil, fmt.Errorf("fzf exited with code %d", code)
	}
	return selectedFiles, nil
}
