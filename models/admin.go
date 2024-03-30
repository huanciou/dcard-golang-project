package models

import (
	"dcard-golang-project/schemas"
)

type AdminSummary struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	EndAt string `json:"end_at"`
}

func FindAdminWithDetails(country, gender, platform string) ([]schemas.Admin, error) {
	var admin []schemas.Admin
	// var result []AdminSummary

	err := DB.Model(&schemas.Admin{}).
		Joins("JOIN countries ON countries.admin_id = admin.id").
		Joins("JOIN genders ON genders.admin_id = admin.id").
		Joins("JOIN platforms ON platforms.admin_id = admin.id").
		Where("countries.country = ? AND genders.gender = ? AND platforms.platform = ?", country, gender, platform).
		Preload("Country").Preload("Gender").Preload("Platform").
		Find(&admin).Error

	return admin, err
}
