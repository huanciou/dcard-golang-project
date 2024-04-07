package utils

import (
	"context"
	"dcard-golang-project/middlewares"
	"dcard-golang-project/models"
	"dcard-golang-project/schemas"
	"encoding/json"
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
		models.Client.SetBit(models.Ctx, "default", int64(index+1), 1)

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

	/* when it fits all conditons */
	if len(queryConditions) == 0 {
		queryConditions = append(queryConditions, "default")
	}
	if len(queryAge) == 0 {
		queryAge = append(queryAge, "default")
	}

	/* setup result bitmap with Lua script*/
	conditionsStr := strings.Join(queryConditions, ",")
	ageStr := strings.Join(queryAge, ",")

	result, err := models.Client.EvalSha(ctx, LuaHash1, nil, conditionsStr, ageStr, params.Offset, params.Limit).Result()

	if err != nil {
		panic(&(middlewares.ServerInternalError{Message: err.Error()}))
	}

	bitMapData, ok := result.(string)
	if !ok {
		panic(&(middlewares.ServerInternalError{Message: err.Error()}))
	}

	indexes := []string{}
	counter := 1
	flag := 1 + params.Limit*(params.Offset-1)

	for i := 0; i < len(bitMapData); i++ {
		for j := 7; j >= 0; j-- { // from LSB
			if len(indexes) >= params.Limit {
				break
			}
			if bitMapData[i]&(1<<j) != 0 && counter >= flag {
				indexes = append(indexes, strconv.Itoa((i*8)+(7-j)))
			} else if bitMapData[i]&(1<<j) != 0 {
				counter++
			}
		}
		if len(indexes) >= params.Limit {
			break
		}
	}

	if len(indexes) == 0 {
		panic(&(middlewares.CustomizedError{Message: "Already reached the last page"}))
	}

	indexesStr := strings.Join(indexes, ",")
	result2, _ := models.Client.EvalSha(ctx, LuaHash2, nil, indexesStr, len(indexes)).Result()

	values, ok := result2.([]interface{})
	if !ok {
		panic(&(middlewares.ServerInternalError{Message: err.Error()}))
	}

	/* json struct unmarshal */
	admins := []schemas.Admin{}
	for _, val := range values {
		if val != nil {
			var admin schemas.Admin
			err := json.Unmarshal([]byte(val.(string)), &admin)
			if err != nil {
				panic(&(middlewares.ServerInternalError{Message: err.Error()}))
			}
			admins = append(admins, admin)
		}
	}

	return admins
}

func Enqueue(post schemas.Admin) {

	jsonPost, err := json.Marshal(post)
	if err != nil {
		panic(&(middlewares.ServerInternalError{Message: err.Error()}))
	}

	ctx := context.Background()

	postLength, err := models.Client.LLen(ctx, "post_queue").Result()
	if postLength >= 3000 {
		panic(&(middlewares.CustomizedError{Message: "Reached the daily quota for creating advertisements"}))
	} else if err != nil {
		panic(&(middlewares.ServerInternalError{Message: err.Error()}))
	}

	if err := models.Client.RPush(ctx, "post_queue", jsonPost).Err(); err != nil {
		panic(&(middlewares.ServerInternalError{Message: err.Error()}))
	}
}

func Dequeue() ([]schemas.Admin, bool) {

	ctx := context.Background()
	var result []schemas.Admin

	resultArr, err := models.Client.LRange(ctx, "post_queue", 0, -1).Result()
	if err != nil {
		panic(&(middlewares.ServerInternalError{Message: err.Error()}))
	}

	/* 當 array 為空，不存入 db */
	if len(resultArr) == 0 {
		return nil, false
	}

	for _, jsonStr := range resultArr {
		var decodedData schemas.Admin
		if err := json.Unmarshal([]byte(jsonStr), &decodedData); err != nil {
			panic(&(middlewares.ServerInternalError{Message: err.Error()}))
		}
		result = append(result, decodedData)
	}

	return result, true
}
