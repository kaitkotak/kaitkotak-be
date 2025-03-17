package master

import (
	"github.com/gofiber/fiber/v3"
	transportvehicle "github.com/kaitkotak-be/internal/modules/master/transport-vehicle"
)

func RegisterMasterRoutes(router fiber.Router) {
	masterGroup := router.Group("/master")

	// Transport Vehicle Routes
	transportVehicleRepo := transportvehicle.NewRepository()
	transportVehicleService := transportvehicle.NewService(transportVehicleRepo)
	transportVehicleHandler := transportvehicle.NewHandler(transportVehicleService)
	transportvehicle.RegisterTransportVehicleRoutes(masterGroup, transportVehicleHandler)
}
