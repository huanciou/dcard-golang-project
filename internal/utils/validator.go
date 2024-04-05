package utils

import (
	"dcard-golang-project/schemas"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/pariz/gountries"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
	Validate.RegisterValidation("countryValidator", countryValidator)
	Validate.RegisterValidation("ageValidator", ageValidator)
	Validate.RegisterValidation("countrySliceValidator", countrySliceValidator)
	Validate.RegisterValidation("genderSliceValidator", genderSliceValidator)
	Validate.RegisterValidation("platformSliceValidator", platformSliceValidator)
}

// get method struct
type GetAdValidation struct {
	Offset   int    `validate:"min=1,max=100" json:"offset"`
	Limit    int    `validate:"required" json:"limit"`
	Age      int    `validate:"ageValidator" json:"age"`
	Country  string `validate:"countryValidator" json:"country"`
	Gender   string `validate:"oneof=m f all" json:"gender"`
	Platform string `validate:"oneof=ios android web all" json:"platform"`
}

// post meethod struct
type PostAdValidation struct {
	Title    string             `validate:"required,max=255"`
	StartAt  time.Time          `validate:"required"`
	EndAt    time.Time          `validate:"required,gtefield=StartAt"`
	AgeStart int                `validate:"oneof= 1 20 24"`
	AgeEnd   int                `validate:"required,oneof=20 24 100,gtefield=AgeStart"`
	Country  []schemas.Country  `validate:"countrySliceValidator" json:"country"`
	Gender   []schemas.Gender   `validate:"genderSliceValidator" json:"gender"`
	Platform []schemas.Platform `validate:"platformSliceValidator" json:"platform"`
}

/* support ISO 3166-1 Alpha-2/Alpha-3 code */
func countryValidator(fl validator.FieldLevel) bool {
	query := gountries.New()
	value := fl.Field().String()

	if _, err := query.FindCountryByAlpha(value); err != nil {
		return false
	}

	return true
}

func ageValidator(fl validator.FieldLevel) bool {
	age := fl.Field().Int()
	return age == -1 || (age >= 1 && age <= 100)
}

func countrySliceValidator(fl validator.FieldLevel) bool {
	sliceValue := fl.Field()

	for i := 0; i < sliceValue.Len(); i++ {

		element := sliceValue.Index(i).Interface()

		country, ok := element.(schemas.Country)
		if !ok {
			return false
		}

		if country.Country != "tw" && country.Country != "cn" && country.Country != "jp" {
			return false
		}
	}

	return true
}

func genderSliceValidator(fl validator.FieldLevel) bool {
	sliceValue := fl.Field()

	for i := 0; i < sliceValue.Len(); i++ {

		element := sliceValue.Index(i).Interface()

		gender, ok := element.(schemas.Gender)
		if !ok {
			return false
		}

		if gender.Gender != "m" && gender.Gender != "f" {
			return false
		}
	}

	return true
}

func platformSliceValidator(fl validator.FieldLevel) bool {
	sliceValue := fl.Field()

	for i := 0; i < sliceValue.Len(); i++ {

		element := sliceValue.Index(i).Interface()

		platform, ok := element.(schemas.Platform)
		if !ok {
			return false
		}

		if platform.Platform != "ios" && platform.Platform != "android" && platform.Platform != "web" {
			return false
		}
	}

	return true
}
