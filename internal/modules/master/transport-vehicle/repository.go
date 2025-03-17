package transportvehicle

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/jackc/pgx/v5"
	"github.com/kaitkotak-be/internal/database"
)

type Repository interface {
	FindAll() ([]TransportVehicle, error)
	FindByID(id int) (*TransportVehicle, error)
	Create(body *TransportVehicleRequestBody) error
	Update(body *TransportVehicleRequestBody, id int) error
	Delete(id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() ([]TransportVehicle, error) {
	rows, err := database.DB.Query(context.Background(), "SELECT id, driver_name, vehicle_number, phone_number FROM transport_vehicles")
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	var transportvehicles []TransportVehicle
	for rows.Next() {
		var transportvehicle TransportVehicle
		if err := rows.Scan(&transportvehicle.ID, &transportvehicle.DriverName, &transportvehicle.VehicleNumber, &transportvehicle.PhoneNumber); err != nil {
			log.Error(err)
			return nil, err
		}
		transportvehicles = append(transportvehicles, transportvehicle)
	}

	return transportvehicles, nil
}

func (r *repository) FindByID(id int) (*TransportVehicle, error) {
	row := database.DB.QueryRow(context.Background(), "SELECT id, driver_name, vehicle_number, phone_number FROM transport_vehicles WHERE id=$1", id)

	var transportvehicle TransportVehicle
	err := row.Scan(&transportvehicle.ID, &transportvehicle.DriverName, &transportvehicle.VehicleNumber, &transportvehicle.PhoneNumber)
	if err == pgx.ErrNoRows {
		return nil, fiber.NewError(fiber.StatusNotFound, "Vehicle not found")
	} else if err != nil {
		log.Error(err)
		return nil, err
	}

	return &transportvehicle, nil
}

func (r *repository) Create(body *TransportVehicleRequestBody) error {
	_, err := database.DB.Exec(
		context.Background(),
		"INSERT INTO transport_vehicles (driver_name, vehicle_number, phone_number) VALUES ($1, $2, $3)",
		body.DriverName, body.VehicleNumber, body.PhoneNumber,
	)
	return err
}

func (r *repository) Update(body *TransportVehicleRequestBody, id int) error {
	_, err := database.DB.Exec(
		context.Background(),
		"UPDATE transport_vehicles SET driver_name=$1, vehicle_number=$2, phone_number=$3 WHERE id=$4",
		body.DriverName, body.VehicleNumber, body.PhoneNumber, id,
	)
	return err
}

func (r *repository) Delete(id int) error {
	_, err := database.DB.Exec(context.Background(), "DELETE FROM transport_vehicles WHERE id=$1", id)
	return err
}
