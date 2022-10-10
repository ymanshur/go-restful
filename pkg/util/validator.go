package util

import (
	"github.com/go-playground/validator/v10"
)

type (
	Validator interface {
		Validate(i interface{}) error
	}
	customValidator struct {
		Validator *validator.Validate
	}
)

func NewValidator() Validator {
	return &customValidator{
		Validator: validator.New(),
	}
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
