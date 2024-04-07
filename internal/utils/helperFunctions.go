package utils

import (
	"math/rand"
	"time"
)

func AgeRangeChecker(queryAge *[]string, age int) {

	if age >= 1 && age <= 20 {
		*queryAge = append(*queryAge, "1to20")
	}
	if age >= 1 && age <= 24 {
		*queryAge = append(*queryAge, "1to24")
	}
	if age >= 1 && age <= 100 {
		*queryAge = append(*queryAge, "1to100")
	}
	if age >= 20 && age <= 24 {
		*queryAge = append(*queryAge, "20to24")
	}
	if age >= 20 && age <= 100 {
		*queryAge = append(*queryAge, "20to100")
	}
	if age >= 24 && age <= 100 {
		*queryAge = append(*queryAge, "24to100")
	}
}

func OptionsChecker(queryConditions *[]string, params GetAdValidation) {
	if params.Gender != "all" {
		*queryConditions = append(*queryConditions, params.Gender)
	}
	if params.Platform != "all" {
		*queryConditions = append(*queryConditions, params.Platform)
	}
	if params.Country != "all" {
		*queryConditions = append(*queryConditions, params.Country)
	}
}

func GenerateRandomDate() time.Time {

	min := time.Now()
	max := time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)

	delta := max.Unix() - min.Unix()
	sec := rand.Int63n(delta) + min.Unix()

	return time.Unix(sec, 0)
}
