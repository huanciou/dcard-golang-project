package utils

import "github.com/go-playground/validator/v10"

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

func CountryValidate(fl validator.FieldLevel) bool {
	countries := []string{"US", "UK", "JP", "TW", "CN"}

	value := fl.Field().String()

	for _, country := range countries {
		if country == value {
			return true
		}
	}
	return false
}
