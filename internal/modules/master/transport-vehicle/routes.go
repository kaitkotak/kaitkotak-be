package transportvehicle

import "github.com/gofiber/fiber/v3"

func RegisterTransportVehicleRoutes(router fiber.Router, handler *Handler) {
	transportVehicleGroup := router.Group("/transport_vehicle")
	transportVehicleGroup.Get("/", handler.GetTransportVehicles)
	transportVehicleGroup.Get("/:id", handler.GetTransportVehicleById)
	transportVehicleGroup.Post("/", handler.CreateTransportVehicle)
	transportVehicleGroup.Put("/:id", handler.UpdateTransportVehicle)
	transportVehicleGroup.Delete("/:id", handler.DeleteTransportVehicle)
}
