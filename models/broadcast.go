package models

type Broadcast struct {
	Offset int
	Limit  int
	Title  string
	EndAt  string
}

func (Broadcast) TableName() string {
	return "broadcast"
}

func GetAd(ad map[string]interface{}) {

}
