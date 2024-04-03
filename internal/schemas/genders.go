package schemas

type Gender struct {
	Id      int    `gorm:"primaryKey;autoIncrement;column:id"`
	Gender  string `gorm:"type:enum('m', 'f');not null;column:gender"`
	AdminId int    `gorm:"not null;column:admin_id"`
}

func (Gender) TableName() string {
	return "genders"
}
