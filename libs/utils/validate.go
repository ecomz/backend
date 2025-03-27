package utils

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(data any) map[string]string {
	errs := validate.Struct(data)
	if errs != nil {
		errors := make(map[string]string)
		for _, err := range errs.(validator.ValidationErrors) {
			errors[err.Field()] = err.Tag()
		}
		return errors
	}

	return nil
}
