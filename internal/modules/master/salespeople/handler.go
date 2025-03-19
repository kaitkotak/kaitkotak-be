package salespeople

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetSalesPeoples(c fiber.Ctx) error {
	SalesPeoples, err := h.service.GetAllSalesPeoples()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Gagal mengambil data sales",
			"errors":  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Berhasil mengambil data sales",
		"data":    SalesPeoples,
		"errors":  nil,
	})
}

func (h *Handler) GetSalesPeopleById(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ID sales tidak valid",
			"errors":  err.Error(),
		})
	}

	SalesPeople, err := h.service.GetSalesPeopleById(id)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"success": false,
			"message": "Gagal mengambil data sales",
			"errors":  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Berhasil mengambil data sales",
		"data":    SalesPeople,
	})
}

func (h *Handler) CreateSalesPeople(c fiber.Ctx) error {
	var body SalesPeopleRequestBody
	if err := c.Bind().JSON(&body); err != nil {
		log.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Gagal menambah data sales",
			"errors":  err.Error(),
		})
	}

	err := h.service.CreateSalesPeople(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Gagal menambah data sales",
			"errors":  err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Berhasil menambah data sales",
		"errors":  nil,
	})
}

func (h *Handler) UpdateSalesPeople(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ID sales tidak valid",
			"errors":  err.Error(),
		})
	}

	var body map[string]interface{}
	if err := c.Bind().JSON(&body); err != nil {
		log.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Gagal mengubah data sales",
			"errors":  err.Error(),
		})
	}

	err = h.service.UpdateSalesPeople(body, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Gagal mengubah data sales",
			"errors":  err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Berhasil mengubah data sales",
		"errors":  nil,
	})
}

func (h *Handler) DeleteSalesPeople(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ID sales tidak valid",
			"errors":  err.Error(),
		})
	}

	err = h.service.DeleteSalesPeople(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Gagal menghapus data sales",
			"errors":  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Berhasil menghapus data sales",
		"errors":  nil,
	})
}
