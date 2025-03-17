package transportvehicle

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

func (h *Handler) GetTransportVehicles(c fiber.Ctx) error {
	transportVehicles, err := h.service.GetAllTransportVehicles()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Gagal mengambil data kendaraan",
			"errors":  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Berhasil mengambil data kendaraan",
		"data":    transportVehicles,
		"errors":  nil,
	})
}

func (h *Handler) GetTransportVehicleById(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ID kendaraan tidak valid",
			"errors":  err.Error(),
		})
	}

	transportVehicle, err := h.service.GetTransportVehicleById(id)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"success": false,
			"message": "Gagal mengambil data kendaraan",
			"errors":  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Berhasil mengambil data kendaraan",
		"data":    transportVehicle,
	})
}

func (h *Handler) CreateTransportVehicle(c fiber.Ctx) error {
	var body TransportVehicleRequestBody
	if err := c.Bind().JSON(&body); err != nil {
		log.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Gagal menambah data kendaraan",
			"errors":  err.Error(),
		})
	}

	err := h.service.CreateTransportVehicle(&body)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Gagal menambah data kendaraan",
			"errors":  err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Berhasil menambah data kendaraan",
		"errors":  nil,
	})
}

func (h *Handler) UpdateTransportVehicle(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ID kendaraan tidak valid",
			"errors":  err.Error(),
		})
	}

	var body TransportVehicleRequestBody
	if err := c.Bind().JSON(&body); err != nil {
		log.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Gagal mengubah data kendaraan",
			"errors":  err.Error(),
		})
	}

	err = h.service.UpdateTransportVehicle(&body, id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Gagal mengubah data kendaraan",
			"errors":  err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Berhasil mengubah data kendaraan",
		"errors":  nil,
	})
}

func (h *Handler) DeleteTransportVehicle(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "ID kendaraan tidak valid",
			"errors":  err.Error(),
		})
	}

	err = h.service.DeleteTransportVehicle(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Gagal menghapus data kendaraan",
			"errors":  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Berhasil menghapus data kendaraan",
		"errors":  nil,
	})
}
