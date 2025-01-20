package util

import (
	"github.com/go-playground/validator/v10"
)

func ValidateDTO(dto interface{}) (any, error) {

	var validate = validator.New()

	if err := validate.Struct(&dto); err != nil {
		errors := make(map[string]string)
		for _, err2 := range err.(validator.ValidationErrors) {
			errors[err2.Field()] = err2.Tag() // Field name and validation tag
		}
		return errors, err
	}

	return nil, nil
}
