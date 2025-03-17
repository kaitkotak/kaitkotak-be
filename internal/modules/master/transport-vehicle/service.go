package transportvehicle

import (
	"database/sql"

	"github.com/gofiber/fiber/v3/log"
	"github.com/kaitkotak-be/internal/shared"
)

type Service interface {
	GetAllTransportVehicles() ([]TransportVehicle, error)
	GetTransportVehicleById(id int) (*TransportVehicle, error)
	CreateTransportVehicle(body *TransportVehicleRequestBody) error
	UpdateTransportVehicle(body *TransportVehicleRequestBody, id int) error
	DeleteTransportVehicle(id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAllTransportVehicles() ([]TransportVehicle, error) {
	return s.repo.FindAll()
}

func (s *service) GetTransportVehicleById(id int) (*TransportVehicle, error) {
	return s.repo.FindByID(id)
}

func (s *service) CreateTransportVehicle(body *TransportVehicleRequestBody) error {
	if err := shared.ValidateStruct(body); err != nil {
		log.Error(err)
		return err
	}
	return s.repo.Create(body)
}

func (s *service) UpdateTransportVehicle(body *TransportVehicleRequestBody, id int) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		log.Error(err)
		return err
	}
	if existing == nil {
		return sql.ErrNoRows
	}

	return s.repo.Update(body, id)
}

func (s *service) DeleteTransportVehicle(id int) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		log.Error(err)
		return err
	}
	if existing == nil {
		return sql.ErrNoRows
	}

	return s.repo.Delete(id)
}
