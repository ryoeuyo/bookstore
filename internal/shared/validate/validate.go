package validate

import "github.com/go-playground/validator/v10"

func IsNotZero(fl validator.FieldLevel) bool {
	return fl.Field().Int() > 0
}

func IsNotEmpty(fl validator.FieldLevel) bool {
	return fl.Field().String() != ""
}
