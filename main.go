package main

import (
	_ "dcard-golang-project/docs"
	"dcard-golang-project/middlewares"
	"dcard-golang-project/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	r := gin.Default()

	/* statics */
	r.LoadHTMLGlob("templates/*")

	/* middlewares */

	r.Use(middlewares.LoggerToFile())
	r.Use(middlewares.ErrorHandler())

	/* rotues */
	routes.ApiRoutersInit(r)
	routes.SwaggerRoutersInit(r)

	PORT := os.Getenv("SERVER_PORT")
	r.Run(PORT)
}
