package utils

import (
	"context"
	"dcard-golang-project/models"
	"dcard-golang-project/schemas"
	"encoding/json"
	"fmt"
	"strconv"
)

/* 從 db 拿出 依照 Order by data (ASC) 後的所有廣告，建立 Bitmap */
func Scheduler(ads []schemas.Admin) {

	for index, ad := range ads {
		idStr := strconv.Itoa(index + 1) // int to str
		adJSON, _ := json.Marshal(ad)

		models.Client.Set(models.Ctx, idStr, adJSON, 0) // id->ad 儲存廣告

		for _, countries := range ad.Country {
			switch countries.Country {
			case "TW":
				models.Client.SetBit(models.Ctx, countries.Country, int64(index+1), 1)
			case "CN":
				models.Client.SetBit(models.Ctx, countries.Country, int64(index+1), 1)
			case "JP":
				models.Client.SetBit(models.Ctx, countries.Country, int64(index+1), 1)
			}
		}
		for _, genders := range ad.Gender {
			switch genders.Gender {
			case "M":
				models.Client.SetBit(models.Ctx, genders.Gender, int64(index+1), 1)
			case "F":
				models.Client.SetBit(models.Ctx, genders.Gender, int64(index+1), 1)
			}
		}
		for _, platforms := range ad.Platform {
			switch platforms.Platform {
			case "ios":
				models.Client.SetBit(models.Ctx, platforms.Platform, int64(index+1), 1)
			case "android":
				models.Client.SetBit(models.Ctx, platforms.Platform, int64(index+1), 1)
			case "web":
				models.Client.SetBit(models.Ctx, platforms.Platform, int64(index+1), 1)
			}
		}
	}
}

func A() []schemas.Admin {
	ctx := context.Background()
	key := "result"

	// cmd := redis.NewScript(Script)
	// result, err := cmd.Run(ctx, models.Client, []string{key}, 1, 3).Result()

	result, err := models.Client.EvalSha(ctx, LuaHash, []string{key}, 1, 3).Result()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	values, ok := result.([]interface{})
	if !ok {
		fmt.Println("Error: invalid result format")
		return nil
	}

	admins := []schemas.Admin{}
	for _, val := range values {
		if val != nil {
			var admin schemas.Admin
			err := json.Unmarshal([]byte(val.(string)), &admin)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			admins = append(admins, admin)
		}
	}

	return admins
}
