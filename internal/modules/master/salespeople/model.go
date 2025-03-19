package salespeople

type SalesPeople struct {
	ID           int     `json:"id"`
	FullName     string  `json:"full_name"`
	PhoneNumber  string  `json:"phone_number"`
	Address      *string `json:"address"`
	Ktp          string  `json:"ktp"`
	Ktp_photo    *string `json:"ktp_photo"`
	ProfilePhoto *string `json:"profile_photo"`
}

type SalesPeopleRequestBody struct {
	FullName     string `json:"full_name" validate:"required"`
	PhoneNumber  string `json:"phone_number" validate:"required"`
	Address      string `json:"address"`
	Ktp          string `json:"ktp" validate:"required"`
	Ktp_photo    string `json:"ktp_photo"`
	ProfilePhoto string `json:"profile_photo"`
}
