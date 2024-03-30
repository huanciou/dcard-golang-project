package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/pariz/gountries"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
	Validate.RegisterValidation("CountryValidate", CountryValidate)
}

type Params struct {
	Offset   int    `validate:"min=1,max=100,required" json:"offset"`
	Limit    int    `validate:"required" json:"limit"`
	Age      int    `validate:"min=1,max=100" json:"age"`
	Gender   string `validate:"oneof=M F,required" json:"gender"`
	Platform string `validate:"oneof=IOS Android Web,required" json:"platform"`
	Country  string `validate:"CountryValidate,required" json:"country"`
}

/* support ISO 3166-1 Alpha-2/Alpha-3 code */
func CountryValidate(fl validator.FieldLevel) bool {
	query := gountries.New()
	value := fl.Field().String()

	if _, err := query.FindCountryByAlpha(value); err != nil {
		return false
	}

	return true
}
