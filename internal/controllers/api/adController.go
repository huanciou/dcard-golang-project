package api

import (
	"dcard-golang-project/middlewares"
	"dcard-golang-project/models"
	"dcard-golang-project/schemas"
	"dcard-golang-project/utils"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary post
// @Description post
// @Tags ad
// @Accept json
// @Produce json
// @Success 200
// @Router /api/v1/ad [post]
func PostAd(c *gin.Context) { // POST
	post := schemas.Admin{}

	/* receive JSON post reqest */
	if err := c.ShouldBindJSON(&post); err != nil {
		panic(&(middlewares.ValidationError{Message: err.Error()}))
	}

	/* lowercase */
	for i := range post.Country {
		post.Country[i].Country = strings.ToLower(post.Country[i].Country)
	}

	for i := range post.Gender {
		post.Gender[i].Gender = strings.ToLower(post.Gender[i].Gender)
	}

	for i := range post.Platform {
		post.Platform[i].Platform = strings.ToLower(post.Platform[i].Platform)
	}

	/* validation */
	postAd := utils.PostAdValidation{
		Title:    post.Title,
		StartAt:  post.StartAt,
		EndAt:    post.EndAt,
		AgeStart: post.AgeStart,
		AgeEnd:   post.AgeEnd,
		Country:  post.Country,
		Gender:   post.Gender,
		Platform: post.Platform,
	}

	if err := utils.Validate.Struct(postAd); err != nil {
		panic(&(middlewares.ValidationError{Message: err.Error()}))
	}

	/* store in queue*/
	utils.Enqueue(post)

	c.JSON(200, gin.H{
		"result": "ok",
	})
}

// @Summary get
// @Description get
// @Tags ad
// @Accept json
// @Produce json
// @Success 200
// @Router /api/v1/ad [get]
func GetAd(c *gin.Context) { // GET

	/* receive query params */
	offsetStr := c.DefaultQuery("offset", "5")
	ageStr := c.DefaultQuery("age", "-1")
	gender := c.DefaultQuery("gender", "all")
	country := c.DefaultQuery("country", "all")
	platform := c.DefaultQuery("platform", "all")

	/* convert string to int */
	offset, _ := strconv.Atoi(offsetStr)
	age, _ := strconv.Atoi(ageStr)

	/* validation */
	getAd := utils.GetAdValidation{
		Offset:   offset,
		Limit:    3,
		Age:      age,
		Gender:   strings.ToLower(gender),
		Country:  strings.ToLower(country),
		Platform: strings.ToLower(platform),
	}
	if err := utils.Validate.Struct(getAd); err != nil {
		panic(&(middlewares.ValidationError{Message: err.Error()}))
	}

	/* redis query */
	result := utils.FilterResultsByConditions(getAd)

	/* response */
	c.JSON(200, result)
}

/* auto generate mock Ads */

func MockData(c *gin.Context) {
	var mockDataSet []schemas.Admin
	var existingDataCount int64

	if err := models.DB.Model(&schemas.Admin{}).Count(&existingDataCount).Error; err != nil {
		panic(&(middlewares.ServerInternalError{Message: err.Error()}))
	}

	targetData := 3000
	numFakeData := targetData - int(existingDataCount)

	if numFakeData <= 0 {
		c.JSON(200, gin.H{
			"result": "already 3000+ Ads exist",
		})
		return
	}

	var mockCountries = []string{
		"tw",
		"jp",
		"cn",
	}

	var mockGenders = []string{
		"m",
		"f",
	}

	var mockPlatforms = []string{
		"ios",
		"android",
		"web",
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numFakeData; i++ {
		var countries []schemas.Country
		var genders []schemas.Gender
		var platforms []schemas.Platform

		numCountries := rand.Intn(len(mockCountries))
		numGenders := rand.Intn(len(mockGenders))
		numPlatforms := rand.Intn(len(mockPlatforms))

		countrySet := make(map[string]bool)
		for i := 0; i <= numCountries; i++ {
			country := mockCountries[rand.Intn(len(mockCountries))]
			countrySet[country] = true
		}

		genderSet := make(map[string]bool)
		for i := 0; i <= numGenders; i++ {
			gender := mockGenders[rand.Intn(len(mockGenders))]
			genderSet[gender] = true
		}

		platformSet := make(map[string]bool)
		for i := 0; i <= numPlatforms; i++ {
			platform := mockPlatforms[rand.Intn(len(mockPlatforms))]
			platformSet[platform] = true
		}

		for country := range countrySet {
			countries = append(countries, schemas.Country{Country: country})
		}
		for gender := range genderSet {
			genders = append(genders, schemas.Gender{Gender: gender})
		}
		for platform := range platformSet {
			platforms = append(platforms, schemas.Platform{Platform: platform})
		}

		startAt := time.Now()
		endAt := utils.GenerateRandomDate()
		for endAt.Before(startAt) {
			endAt = utils.GenerateRandomDate()
		}

		data := schemas.Admin{
			Title:    fmt.Sprintf("廣告%v", i),
			StartAt:  startAt,
			EndAt:    endAt,
			AgeStart: 1,
			AgeEnd:   100,
			Country:  countries,
			Gender:   genders,
			Platform: platforms,
		}
		mockDataSet = append(mockDataSet, data)
	}

	models.DB.Create(&mockDataSet)

	c.JSON(200, gin.H{
		"result": fmt.Sprintf("Inserted %v advertisements", numFakeData),
	})
}
