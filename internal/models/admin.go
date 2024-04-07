package models

import (
	"dcard-golang-project/schemas"
	"time"
)

type AdminSummary struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	EndAt string `json:"end_at"`
}

func AdQueryWithDate() ([]schemas.Admin, error) {
	var admin []schemas.Admin
	// var result []AdminSummary

	now := time.Now()

	err := DB.Model(&schemas.Admin{}).
		Joins("JOIN countries ON countries.admin_id = admin.id").
		Joins("JOIN genders ON genders.admin_id = admin.id").
		Joins("JOIN platforms ON platforms.admin_id = admin.id").
		Preload("Country").Preload("Gender").Preload("Platform").
		Where("admin.deleted_at IS NULL").
		Where("admin.start_at < ? AND admin.end_at > ?", now, now).
		Order("end_at ASC").
		Distinct(). // 確保一對多結構在 Prelaod 的時候不會重複返回多個 Admin
		Find(&admin).Error

	return admin, err
}
