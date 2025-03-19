package salespeople

import "github.com/gofiber/fiber/v3"

func RegisterSalesPeopleRoutes(router fiber.Router, handler *Handler) {
	SalesPeopleGroup := router.Group("/sales_people")
	SalesPeopleGroup.Get("/", handler.GetSalesPeoples)
	SalesPeopleGroup.Get("/:id", handler.GetSalesPeopleById)
	SalesPeopleGroup.Post("/", handler.CreateSalesPeople)
	SalesPeopleGroup.Put("/:id", handler.UpdateSalesPeople)
	SalesPeopleGroup.Delete("/:id", handler.DeleteSalesPeople)
}
