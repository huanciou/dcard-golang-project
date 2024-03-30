package schemas

import "time"

type Admin struct {
	Id       int       `gorm:"primaryKey;autoIncrement;column:id"`
	Title    string    `gorm:"type:varchar(255);not null;column:title"`
	StartAt  time.Time `gorm:"type:date;not null;column:start_at"`
	EndAt    time.Time `gorm:"type:date;not null;column:end_at"`
	AgeStart int       `gorm:"type:tinyint;not null;column:age_start"`
	AgeEnd   int       `gorm:"type:tinyint;not null;column:age_end"`
	Country  []Country
	Gender   []Gender
	Platform []Platform
}

func (Admin) TableName() string {
	return "admin"
}
