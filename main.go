package main

import (
	"fmt"
	"os"

	"tempvault/cmd"
	"tempvault/config"
)

func createTempVaultDir() error {
	tempVaultDir, err := config.GetTempVaultDir()
	if err != nil {
		return err
	}

	if _, err := os.Stat(tempVaultDir); os.IsNotExist(err) {
		err := os.MkdirAll(tempVaultDir, os.ModePerm)
		if err != nil {
			return err
		}
		fmt.Println("Created directory:", tempVaultDir)
	}

	return nil
}

func main() {
	err := createTempVaultDir()
	if err != nil {
		fmt.Println("Error creating temp vault directory:", err)
		os.Exit(1)
	}

	cmd.Execute()
}
