package utils_test

import (
	"dcard-golang-project/utils"
	"reflect"
	"testing"
)

func TestAgeRangeChecker(t *testing.T) {
	var mockQueryAge []string
	utils.AgeRangeChecker(&mockQueryAge, 24)

	expectedValues := []string{"1to24", "1to100", "20to24", "20to100", "24to100"}
	if !reflect.DeepEqual(mockQueryAge, expectedValues) {
		t.Errorf("TestAgeRangeChecker failed, got: %v, want: %v", mockQueryAge, expectedValues)
	}
}

func TestOptionsChecker(t *testing.T) {
	var mockQueryConditions []string

	mockData := utils.GetAdValidation{
		Offset:   5,
		Limit:    3,
		Age:      24,
		Country:  "tw",
		Gender:   "m",
		Platform: "all",
	}

	utils.OptionsChecker(&mockQueryConditions, mockData)

	expectedValues := []string{"m", "tw"}
	if !reflect.DeepEqual(mockQueryConditions, expectedValues) {
		t.Errorf("TestOptionsChecker failed, got: %v, want: %v", mockQueryConditions, expectedValues)
	}

}
