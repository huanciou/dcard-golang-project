package utils_test

import (
	"dcard-golang-project/schemas"
	"dcard-golang-project/utils"
	"testing"
	"time"
)

func TestPostAdValidation(t *testing.T) {
	mockPost := utils.PostAdValidation{
		Title:    "廣告字串",
		StartAt:  time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC),
		EndAt:    time.Date(2024, time.March, 2, 0, 0, 0, 0, time.UTC),
		AgeStart: 1,
		AgeEnd:   100,
		Country:  []schemas.Country{{Country: "tw"}, {Country: "jp"}, {Country: "cn"}},
		Gender:   []schemas.Gender{{Gender: "m"}, {Gender: "f"}},
		Platform: []schemas.Platform{{Platform: "ios"}, {Platform: "android"}, {Platform: "web"}},
	}

	err := utils.Validate.Struct(mockPost)
	if err != nil {
		t.Errorf("TestPostAdValidation failed: %v", err)
	}

}
