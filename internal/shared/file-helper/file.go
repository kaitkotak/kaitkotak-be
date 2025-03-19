package filehelper

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetAbsolutePath(subPath string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return "", err
	}

	return filepath.Join(wd, subPath), nil
}
