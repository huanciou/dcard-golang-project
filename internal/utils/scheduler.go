package utils

import (
	"context"
	"dcard-golang-project/middlewares"
	"dcard-golang-project/models"
	"dcard-golang-project/schemas"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

/* 從 db 拿出 依照 Order by date (ASC) 後的所有廣告，建立 Bitmap */
func SetBitmaps() {

	ads, err := models.AdQueryWithDate()
	if err != nil {
		panic(&(middlewares.ServerInternalError{Message: err.Error()}))
	}

	for index, ad := range ads {
		idStr := strconv.Itoa(index + 1) // int to str
		adJSON, _ := json.Marshal(ad)

		models.Client.Set(models.Ctx, idStr, adJSON, 0) // id->ad 儲存廣告

		switch {
		case ad.AgeStart == 1 && ad.AgeEnd == 20:
			models.Client.SetBit(models.Ctx, "1to20", int64(index+1), 1)
		case ad.AgeStart == 1 && ad.AgeEnd == 24:
			models.Client.SetBit(models.Ctx, "1to24", int64(index+1), 1)
		case ad.AgeStart == 1 && ad.AgeEnd == 100:
			models.Client.SetBit(models.Ctx, "1to100", int64(index+1), 1)
		case ad.AgeStart == 20 && ad.AgeEnd == 24:
			models.Client.SetBit(models.Ctx, "20to24", int64(index+1), 1)
		case ad.AgeStart == 20 && ad.AgeEnd == 100:
			models.Client.SetBit(models.Ctx, "20to100", int64(index+1), 1)
		case ad.AgeStart == 24 && ad.AgeEnd == 100:
			models.Client.SetBit(models.Ctx, "24to100", int64(index+1), 1)
		}

		for _, countries := range ad.Country {
			switch countries.Country {
			case "tw":
				models.Client.SetBit(models.Ctx, countries.Country, int64(index+1), 1)
			case "cn":
				models.Client.SetBit(models.Ctx, countries.Country, int64(index+1), 1)
			case "jp":
				models.Client.SetBit(models.Ctx, countries.Country, int64(index+1), 1)
			}
		}

		for _, genders := range ad.Gender {
			switch genders.Gender {
			case "m":
				models.Client.SetBit(models.Ctx, genders.Gender, int64(index+1), 1)
			case "f":
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

func FilterResultsByConditions(params GetAdValidation) []schemas.Admin {
	ctx := context.Background()
	queryConditions := []string{}
	queryAge := []string{}

	/* setup query conditions */
	AgeRangeChecker(&queryAge, params.Age)
	OptionsChecker(&queryConditions, params)

	/* setup result bitmap with Lua script*/
	conditionsStr := strings.Join(queryConditions, ",")
	ageStr := strings.Join(queryAge, ",")

	result, err := models.Client.EvalSha(ctx, LuaHash1, nil, conditionsStr, ageStr, params.Offset, params.Limit).Result()

	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	bitMapData, ok := result.(string)
	if !ok {
		fmt.Println("Error: Result is not a string")
	}

	indexes := []string{}
	counter := 1
	flag := 1 + params.Limit*(params.Offset-1)

	for i := 0; len(indexes) < params.Limit; i++ {
		for j := 7; j >= 0 && len(indexes) < params.Limit; j-- { // from LSB
			if bitMapData[i]&(1<<j) != 0 && counter >= flag {
				indexes = append(indexes, strconv.Itoa((i*8)+(7-j)+1))
			} else if bitMapData[i]&(1<<j) != 0 {
				counter++
			}
		}
	}

	indexesStr := strings.Join(indexes, ",")
	result2, _ := models.Client.EvalSha(ctx, LuaHash2, nil, indexesStr, params.Offset, params.Limit).Result()

	values, ok := result2.([]interface{})
	if !ok {
		fmt.Println("Error: invalid result format")
		return nil
	}

	/* json struct unmarshal */
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
