package models

type Admin struct {
	Title      string `json:"title"`
	StartAt    string `json:"startAt"`
	EndAt      string `json:"endAt"`
	Conditions `json:"conditions"`
}

type Conditions struct {
	Age      int      `json:"age"`
	Gender   []string `json:"gender"`
	Country  []string `json:"country"`
	Platform []string `json:"platform"`
}

func (Admin) TableName() string {
	return "admin"
}

func PostAd(p Admin) {
}

func AutoPostMockAd(p []Admin) ([]Admin, error) {
	return []Admin{}, nil
}
