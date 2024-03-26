package utils

import (
	"dcard-golang-project/models"
	"dcard-golang-project/schemas"
	"encoding/json"
	"strconv"
)

// var ads = []schemas.Admin{
// 	{
// 		Id:       4,
// 		Title:    "廣告3",
// 		StartAt:  time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC),
// 		EndAt:    time.Date(2024, time.March, 2, 0, 0, 0, 0, time.UTC),
// 		AgeStart: 1,
// 		AgeEnd:   100,
// 		Country: []schemas.Country{
// 			{
// 				Id:      6,
// 				Country: "CN",
// 				AdminId: 4,
// 			},
// 			{
// 				Id:      7,
// 				Country: "TW",
// 				AdminId: 4,
// 			},
// 		},
// 		Gender: []schemas.Gender{
// 			{
// 				Id:      5,
// 				Gender:  "F",
// 				AdminId: 4,
// 			},
// 			{
// 				Id:      6,
// 				Gender:  "M",
// 				AdminId: 4,
// 			},
// 		},
// 		Platform: []schemas.Platform{
// 			{
// 				Id:       5,
// 				Platform: "android",
// 				AdminId:  4,
// 			},
// 		},
// 	},
// }

func Scheduler(ads []schemas.Admin) {

	for _, ad := range ads {
		idStr := strconv.Itoa(ad.Id)
		idInt64 := int64(ad.Id)
		adJSON, _ := json.Marshal(ad)

		models.Client.Set(models.Ctx, idStr, adJSON, 0) // id->ad 儲存廣告

		for _, countries := range ad.Country {
			switch countries.Country {
			case "TW":
				models.Client.SetBit(models.Ctx, countries.Country, idInt64, 1)
			case "CN":
				models.Client.SetBit(models.Ctx, countries.Country, idInt64, 1)
			case "JP":
				models.Client.SetBit(models.Ctx, countries.Country, idInt64, 1)
			}
		}
		for _, genders := range ad.Gender {
			switch genders.Gender {
			case "M":
				models.Client.SetBit(models.Ctx, genders.Gender, idInt64, 1)
			case "F":
				models.Client.SetBit(models.Ctx, genders.Gender, idInt64, 1)
			}
		}
		for _, platforms := range ad.Platform {
			switch platforms.Platform {
			case "ios":
				models.Client.SetBit(models.Ctx, platforms.Platform, idInt64, 1)
			case "android":
				models.Client.SetBit(models.Ctx, platforms.Platform, idInt64, 1)
			case "web":
				models.Client.SetBit(models.Ctx, platforms.Platform, idInt64, 1)
			}
		}
	}
}
