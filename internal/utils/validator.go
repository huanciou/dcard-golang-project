package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/pariz/gountries"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
	Validate.RegisterValidation("CountryValidator", CountryValidator)
	Validate.RegisterValidation("AgeValidator", AgeValidator)
}

type Params struct {
	Offset   int    `validate:"min=1,max=100" json:"offset"`
	Limit    int    `validate:"required" json:"limit"`
	Age      int    `validate:"AgeValidator" json:"age"`
	Gender   string `validate:"oneof=m f nil" json:"gender"`
	Platform string `validate:"oneof=ios android web nil" json:"platform"`
	Country  string `validate:"CountryValidator" json:"country"`
}

/* support ISO 3166-1 Alpha-2/Alpha-3 code */
func CountryValidator(fl validator.FieldLevel) bool {
	query := gountries.New()
	value := fl.Field().String()

	if _, err := query.FindCountryByAlpha(value); err != nil {
		return false
	}

	return true
}

func AgeValidator(fl validator.FieldLevel) bool {
	age := fl.Field().Int()
	return age == -1 || (age >= 1 && age <= 100)
}
