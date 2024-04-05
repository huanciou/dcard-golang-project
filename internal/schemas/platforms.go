package schemas

type Platform struct {
	Id       int    `gorm:"primaryKey;autoIncrement;column:id"`
	Platform string `json:"platform" gorm:"type:enum('ios', 'android', 'web');not null;column:platform"`
	AdminId  int    `gorm:"not null;column:admin_id"`
}

func (Platform) TableName() string {
	return "platforms"
}
