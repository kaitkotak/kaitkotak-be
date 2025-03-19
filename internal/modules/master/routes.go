package master

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kaitkotak-be/internal/modules/master/salespeople"
	transportvehicle "github.com/kaitkotak-be/internal/modules/master/transport-vehicle"
)

func RegisterMasterRoutes(router fiber.Router) {
	masterGroup := router.Group("/master")

	// Transport Vehicle Routes
	transportVehicleRepo := transportvehicle.NewRepository()
	transportVehicleService := transportvehicle.NewService(transportVehicleRepo)
	transportVehicleHandler := transportvehicle.NewHandler(transportVehicleService)
	transportvehicle.RegisterTransportVehicleRoutes(masterGroup, transportVehicleHandler)

	// Sales People Routes
	salesPeopleRepo := salespeople.NewRepository()
	salesPeopleService := salespeople.NewService(salesPeopleRepo)
	salesPeopleHandler := salespeople.NewHandler(salesPeopleService)
	salespeople.RegisterSalesPeopleRoutes(masterGroup, salesPeopleHandler)
}
