package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

type SignUpInput struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func FormatValidationError(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			var msg string

			switch fieldErr.Tag() {
			case "required":
				msg = "Este campo é obrigatório"
			case "min":
				msg = fmt.Sprintf("Deve ter no mínimo %s caracteres", fieldErr.Param())
			case "email":
				msg = "email inválido"
			default:
				msg = "Valor inválido"
			}

			errors[fieldErr.Field()] = msg
		}
	}
	return errors
}


type SignInInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}