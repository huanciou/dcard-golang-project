package schemas

type Country struct {
	Id      int    `gorm:"primaryKey;autoIncrement;column:id"`
	Country string `json:"country" gorm:"type:enum('tw', 'cn', 'jp');not null;column:country"`
	AdminId int    `gorm:"not null;column:admin_id"`
}

func (Country) TableName() string {
	return "countries"
}
