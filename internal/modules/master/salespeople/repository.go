package salespeople

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/jackc/pgx/v5"
	"github.com/kaitkotak-be/internal/database"
)

type Repository interface {
	FindAll() ([]SalesPeople, error)
	FindByID(id int) (*SalesPeople, error)
	Create(body *SalesPeopleRequestBody) error
	Update(body map[string]interface{}, id int) error
	Delete(id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() ([]SalesPeople, error) {
	rows, err := database.DB.Query(context.Background(), "SELECT id, full_name, phone_number, address, ktp, ktp_photo, profile_photo FROM salespeople")
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	var SalesPeoples []SalesPeople
	SalesPeoples = []SalesPeople{}
	for rows.Next() {
		var SalesPeople SalesPeople
		if err := rows.Scan(&SalesPeople.ID, &SalesPeople.FullName, &SalesPeople.PhoneNumber, &SalesPeople.Address, &SalesPeople.Ktp, &SalesPeople.Ktp_photo, &SalesPeople.ProfilePhoto); err != nil {
			log.Error(err)
			return nil, err
		}
		SalesPeoples = append(SalesPeoples, SalesPeople)
	}

	return SalesPeoples, nil
}

func (r *repository) FindByID(id int) (*SalesPeople, error) {
	row := database.DB.QueryRow(context.Background(), "SELECT id, full_name, phone_number, address, ktp, ktp_photo, profile_photo FROM salespeople WHERE id=$1", id)

	var SalesPeople SalesPeople
	err := row.Scan(&SalesPeople.ID, &SalesPeople.FullName, &SalesPeople.PhoneNumber, &SalesPeople.Address, &SalesPeople.Ktp, &SalesPeople.Ktp_photo, &SalesPeople.ProfilePhoto)
	if err == pgx.ErrNoRows {
		return nil, fiber.NewError(fiber.StatusNotFound, "Vehicle not found")
	} else if err != nil {
		log.Error(err)
		return nil, err
	}

	return &SalesPeople, nil
}

func (r *repository) Create(body *SalesPeopleRequestBody) error {
	_, err := database.DB.Exec(
		context.Background(),
		"INSERT INTO salespeople (full_name, phone_number, address, ktp, ktp_photo, profile_photo) VALUES ($1, $2, $3, $4, $5, $6)",
		&body.FullName, &body.PhoneNumber, &body.Address, &body.Ktp, &body.Ktp_photo, &body.ProfilePhoto,
	)
	return err
}

func (r *repository) Update(updateData map[string]interface{}, id int) error {
	if len(updateData) == 0 {
		return errors.New("there is no data to update")
	}

	var queryParts []string
	var values []interface{}
	i := 1

	for key, value := range updateData {
		queryParts = append(queryParts, fmt.Sprintf("%s=$%d", key, i))
		values = append(values, value)
		i++
	}

	query := fmt.Sprintf("UPDATE salespeople SET %s WHERE id=$%d", strings.Join(queryParts, ", "), i)
	values = append(values, id)

	_, err := database.DB.Exec(context.Background(), query, values...)
	return err
}

func (r *repository) Delete(id int) error {
	_, err := database.DB.Exec(context.Background(), "DELETE FROM salespeople WHERE id=$1", id)
	return err
}
