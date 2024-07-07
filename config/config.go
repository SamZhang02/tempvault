package config

import (
	"os"
	"path/filepath"
)

func GetTempVaultDir() (string, error) {
	var homeDir, err = os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, "tempvault"), nil
}
