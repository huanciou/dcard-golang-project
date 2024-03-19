package routes

import (
	"dcard-golang-project/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api/v1")
	{
		apiRouters.GET("/ad", api.Ad{}.Broadcast)
		apiRouters.POST("/ad", api.Ad{}.Admin)
		apiRouters.GET("/mock", api.Ad{}.MockData)
	}
}
