package utils

import (
	"github.com/go-playground/validator/v10"
)

type ValidationUtil struct {
	validator *validator.Validate
}

func NewValidationUtil() *ValidationUtil {
	return &ValidationUtil{validator: validator.New()}
}

func (v *ValidationUtil) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
