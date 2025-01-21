package util

import (
	"github.com/go-playground/validator/v10"
)

func ValidateDTO(dto interface{}) (any, error) {

	var validate = validator.New()

	err := validate.Struct(dto)
	if err != nil {

		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			errs := make(map[string]string)
			for _, err2 := range validationErrs {
				errs[err2.Field()] = err2.Tag() // Field name and validation tag
			}
			return errs, err
		}

		return err.Error(), err
	}

	return nil, nil
}
