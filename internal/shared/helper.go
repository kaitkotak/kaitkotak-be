package shared

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err != nil {
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, err.Field()+" is invalid: "+err.Tag())
		}
		return errors.New(strings.Join(errorMessages, ", "))
	}
	return nil
}
