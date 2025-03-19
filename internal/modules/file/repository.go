package file

import (
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/kaitkotak-be/internal/config"
	filehelper "github.com/kaitkotak-be/internal/shared/file-helper"
	"github.com/kaitkotak-be/internal/shared/helper"
)

type Repository interface {
	UploadFile(file *multipart.FileHeader, c fiber.Ctx) (*File, error)
	DownloadFile(fileName string) (string, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) UploadFile(file *multipart.FileHeader, c fiber.Ctx) (fileDetail *File, err error) {
	temporaryPath, err := filehelper.GetAbsolutePath(config.Value.String("upload-file.temp-dir"))
	if err != nil {
		log.Errorf("Failed to get absolute path: %v", err)
		return nil, err
	}
	fileExt := strings.TrimPrefix(filepath.Ext(file.Filename), ".")
	tempFileName := helper.GetTempFileName(fileExt)
	fullPath := filepath.Join(temporaryPath, tempFileName)
	err = c.SaveFile(file, fullPath)
	if err != nil {
		log.Errorf("Failed to get absolute path: %v", err)
		return nil, err
	}

	fileDetail = &File{
		Name: tempFileName,
	}
	return fileDetail, nil
}

func (r *repository) DownloadFile(fileName string) (string, error) {
	filepath, err := filehelper.GetAbsolutePath(filepath.Join(config.Value.String("upload-file.upload-dir"), fileName))
	if err != nil {
		log.Errorf("Failed to get absolute path: %v", err)
		return "", err
	}

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		log.Error(err)
		return "", err
	}

	return filepath, err
}
