package file

import (
	"github.com/gofiber/fiber/v3"
)

func RegisterFileRouter(router fiber.Router) {
	fileRepo := NewRepository()
	fileService := NewService(fileRepo)
	fileHandler := NewHandler(fileService)

	fileGroup := router.Group("/file")
	fileGroup.Post("/upload", fileHandler.UploadFile)
	fileGroup.Get("/download/:filename", fileHandler.DownloadFile)
}
