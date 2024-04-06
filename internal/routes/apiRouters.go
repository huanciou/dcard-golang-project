package routes

import (
	"dcard-golang-project/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api/v1")
	{
		apiRouters.GET("/ad", api.GetAd)
		apiRouters.POST("/ad", api.PostAd)
		apiRouters.GET("/mock", api.MockData)
	}
}
