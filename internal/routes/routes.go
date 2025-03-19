package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kaitkotak-be/internal/modules/file"
	"github.com/kaitkotak-be/internal/modules/master"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/")

	master.RegisterMasterRoutes(api)
	file.RegisterFileRouter(api)
}
