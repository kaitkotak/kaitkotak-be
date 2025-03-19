package helper

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/kaitkotak-be/internal/config"
	filehelper "github.com/kaitkotak-be/internal/shared/file-helper"
)

var validate = validator.New()

func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err != nil {
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, err.Field()+" is invalid: "+err.Tag())
		}
		return errors.New(strings.Join(errorMessages, ", "))
	}
	return nil
}

func GetTempFileName(extension string) string {
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(1e9)) // Secure random number
	return fmt.Sprintf("temp_%d.%s", randomNumber, extension)
}

func ContainsString(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func MoveUploadedFile(tempFilePath, newFilePath string) error {
	destFolder := filepath.Dir(newFilePath)
	if err := EnsureDirExists(destFolder); err != nil {
		return err
	}

	if err := os.Rename(tempFilePath, newFilePath); err != nil {
		return errors.New("no such file temporary in directory")
	}

	return nil
}

func EnsureDirExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}
	return nil
}

func ProcessUploadedFile(filePath, prefix string) (string, error) {
	if filePath == "" {
		return "", nil
	}

	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(1e9))
	fileExt := filepath.Ext(filePath)

	tempDir, err := filehelper.GetAbsolutePath(config.Value.String("upload-file.temp-dir"))
	if err != nil {
		return "", err
	}

	uploadedPath, err := filehelper.GetAbsolutePath(config.Value.String("upload-file.upload-dir"))
	if err != nil {
		return "", err
	}

	newFileName := fmt.Sprintf("%s_%d%s", prefix, randomNumber, fileExt)
	if err := MoveUploadedFile(filepath.Join(tempDir, filePath), filepath.Join(uploadedPath, newFileName)); err != nil {
		return "", err
	}

	return newFileName, nil
}
