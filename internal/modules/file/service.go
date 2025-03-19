package file

import (
	"errors"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/kaitkotak-be/internal/config"
	"github.com/kaitkotak-be/internal/shared/helper"
)

type Service interface {
	UploadFile(file *multipart.FileHeader, c fiber.Ctx) (*File, error)
	DownloadFile(fileName string) (string, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) UploadFile(file *multipart.FileHeader, c fiber.Ctx) (fileDetail *File, err error) {
	if file.Size > config.Value.Int64("upload-file.max-size")*1024 {
		return nil, errors.New("File size exceeds 5MB")
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !helper.ContainsString(config.Value.Strings("upload-file.allowed-types"), ext) {
		return nil, errors.New("invalid file type. only images are allowed")
	}

	return s.repo.UploadFile(file, c)
}

func (s *service) DownloadFile(fileName string) (string, error) {
	return s.repo.DownloadFile(fileName)
}
