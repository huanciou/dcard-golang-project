package routes

import (
	_ "dcard-golang-project/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerRoutersInit(r *gin.Engine) {
	swaggerRouters := r.Group("/swagger")
	{
		swaggerRouters.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
