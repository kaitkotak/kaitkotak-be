package file

import (
	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) UploadFile(c fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Gagal mengunggah berkas",
			"errors":  err.Error(),
		})
	}

	fileDetail, err := h.service.UploadFile(file, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Gagal mengunggah berkas",
			"errors":  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Berhasil menungunggah berkas",
		"data":    fileDetail,
		"errors":  nil,
	})
}

func (h *Handler) DownloadFile(c fiber.Ctx) error {
	fileName := c.Params("filename")
	filepath, err := h.service.DownloadFile(fileName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Gagal Mengunduh berkas",
			"errors":  err.Error(),
		})
	}

	return c.SendFile(filepath)
}
