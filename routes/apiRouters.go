package routes

import (
	"dcard-golang-project/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/v1/api")
	{
		apiRouters.GET("/admin", api.AD{}.Admin)
		apiRouters.GET("/broadcast", api.AD{}.Broadcast)
	}
}
