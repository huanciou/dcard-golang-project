package api

import (
	"dcard-golang-project/middlewares"
	"dcard-golang-project/models"
	"dcard-golang-project/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Ad struct{}

/* manually create Ad */

func (Ad) Admin(c *gin.Context) { // POST
	post := models.Admin{}

	if err := c.ShouldBindJSON(&post); err != nil {
		panic(&(middlewares.ValidationError{Message: err.Error()}))
	}
}

/* broadcasting */

func (Ad) Broadcast(c *gin.Context) { // GET

	var offsetStr string = c.DefaultQuery("offset", "5")
	var ageStr string = c.DefaultQuery("age", "-1")
	var gender string = c.DefaultQuery("gender", "-1")
	var country string = c.DefaultQuery("country", "-1")
	var platform string = c.DefaultQuery("platform", "-1")

	/* convert string to int */

	offset, _ := strconv.Atoi(offsetStr)
	age, _ := strconv.Atoi(ageStr)

	params := utils.Params{
		Offset:   offset,
		Limit:    3,
		Age:      age,
		Gender:   gender,
		Country:  country,
		Platform: platform,
	}

	if err := utils.Validate.Struct(params); err != nil {
		panic(&(middlewares.ValidationError{Message: "Validation Error"}))
	} else {
		c.JSON(200, params)
	}

}

/* auto generate mock Ads */

func (Ad) MockData(c *gin.Context) {
	var mockDataSet []models.Admin

	for i := 0; i < 1000; i++ {
		age := i + 1
		data := models.Admin{
			Title:   fmt.Sprintf("Ad%v", i),
			StartAt: "0301",
			EndAt:   "0302",
			Conditions: models.Conditions{
				Age:      &age,
				Gender:   []string{"M", "F"},
				Country:  []string{"TW", "JP"},
				Platform: []string{"IOS", "Android", "Web"},
			},
		}
		mockDataSet = append(mockDataSet, data)
	}
	c.JSON(200, mockDataSet)
	// if data, err := models.AutoPostMockAd(mockDataSet); err != nil {
	// 	panic(&(middlewares.ServerInternalError{Message: "Ad Insertion error"}))
	// } else {
	// 	c.JSON(200, data)
	// }
}
