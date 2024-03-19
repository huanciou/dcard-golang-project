package models

type Broadcast struct {
	Title string
	EndAt string
}

func (Broadcast) TableName() string {
	return "broadcast"
}

func GetAd(ad map[string]interface{}) {

}
