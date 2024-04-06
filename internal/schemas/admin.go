package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Title    string     `json:"title" gorm:"type:varchar(255);not null;column:title"`
	StartAt  time.Time  `json:"startAt" gorm:"type:date;not null;column:start_at"`
	EndAt    time.Time  `json:"endAt" gorm:"type:date;not null;column:end_at"`
	AgeStart int        `json:"ageStart" gorm:"type:tinyint;not null;column:age_start"`
	AgeEnd   int        `json:"ageEnd" gorm:"type:tinyint;not null;column:age_end"`
	Country  []Country  `json:"country"`
	Gender   []Gender   `json:"gender"`
	Platform []Platform `json:"platform"`
}

func (Admin) TableName() string {
	return "admin"
}
