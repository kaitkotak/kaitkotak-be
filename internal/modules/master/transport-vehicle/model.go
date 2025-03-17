package transportvehicle

type TransportVehicle struct {
	ID            int    `json:"id"`
	DriverName    string `json:"driver_name"`
	VehicleNumber string `json:"vehicle_number"`
	PhoneNumber   string `json:"phone_number"`
}

type TransportVehicleRequestBody struct {
	DriverName    string `json:"driver_name" validate:"required,min=3"`
	VehicleNumber string `json:"vehicle_number" validate:"required,alphanum"`
	PhoneNumber   string `json:"phone_number" validate:"required"`
}
