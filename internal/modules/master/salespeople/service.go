package salespeople

import (
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v3/log"
	"github.com/kaitkotak-be/internal/shared/helper"
)

type Service interface {
	GetAllSalesPeoples() ([]SalesPeople, error)
	GetSalesPeopleById(id int) (*SalesPeople, error)
	CreateSalesPeople(body *SalesPeopleRequestBody) error
	UpdateSalesPeople(body map[string]interface{}, id int) error
	DeleteSalesPeople(id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAllSalesPeoples() ([]SalesPeople, error) {
	return s.repo.FindAll()
}

func (s *service) GetSalesPeopleById(id int) (*SalesPeople, error) {
	return s.repo.FindByID(id)
}

func (s *service) CreateSalesPeople(body *SalesPeopleRequestBody) error {
	if err := helper.ValidateStruct(body); err != nil {
		log.Error(err)
		return err
	}

	if body.Ktp_photo != "" {
		fileName, err := helper.ProcessUploadedFile(body.Ktp_photo, "salespeople-ktp")
		if err != nil {
			return err
		}
		body.Ktp_photo = fileName
	}

	if body.ProfilePhoto != "" {
		fileName, err := helper.ProcessUploadedFile(body.ProfilePhoto, "salespeople-profile")
		if err != nil {
			return err
		}
		body.ProfilePhoto = fileName
	}

	return s.repo.Create(body)
}

func (s *service) UpdateSalesPeople(body map[string]interface{}, id int) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		log.Error(err)
		return err
	}
	if existing == nil {
		return sql.ErrNoRows
	}

	updateData := make(map[string]interface{})
	if ktpPhoto, ok := body["ktp_photo"].(string); ok && ktpPhoto != "" {
		fileName, err := helper.ProcessUploadedFile(ktpPhoto, "salespeople-ktp")
		if err != nil {
			return err
		}
		updateData["ktp_photo"] = fileName
	}

	if profilePhoto, ok := body["profile_photo"].(string); ok && profilePhoto != "" {
		fileName, err := helper.ProcessUploadedFile(profilePhoto, "salespeople-profile")
		if err != nil {
			return err
		}
		updateData["profile_photo"] = fileName
	}

	allowedFields := []string{"full_name", "phone_number", "address", "ktp"}
	for _, field := range allowedFields {
		if value, ok := body[field]; ok && value != "" {
			updateData[field] = value
		}
	}

	if len(updateData) == 0 {
		return errors.New("tidak ada data yang diperbarui")
	}

	return s.repo.Update(updateData, id)
}

func (s *service) DeleteSalesPeople(id int) error {
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
