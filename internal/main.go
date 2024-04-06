package main

import (
	"dcard-golang-project/middlewares"
	"dcard-golang-project/models"
	"dcard-golang-project/routes"
	"dcard-golang-project/utils"
	"dcard-golang-project/workers"
	"fmt"
	"os"

	_ "dcard-golang-project/docs"

	"github.com/gin-contrib/pprof"

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

	env := os.Getenv("ENV")
	var envFile string

	switch env {
	case "production":
		envFile = ".env.prod"
		fmt.Println("using production mode")
	default:
		envFile = ".env.dev"
		fmt.Println("using dev mode")
	}

	if err := godotenv.Load(envFile); err != nil {
		fmt.Println("Error loading .env file")
	}

	/* db initiation */
	models.DBInit()

	/* redis initiation */
	models.RedisInit()

	/* Load Lua Script*/
	utils.LoadLuaScript()

	/* 分配一個 goroutine 進行 cron job */
	go workers.CronJob()

	/* 當每次重啟時，更新 redis Bitmap */
	go utils.SetBitmaps()
}

func main() {
	r := gin.Default()

	/* statics */
	// r.LoadHTMLGlob("templates/*")

	/* middlewares */
	// r.Use(middlewares.LoggerToFile())
	r.Use(middlewares.ErrorHandler())

	/* rotues */
	routes.ApiRoutersInit(r)
	routes.SwaggerRoutersInit(r)

	pprof.Register(r)

	PORT := os.Getenv("SERVER_PORT")
	r.Run(PORT)
}
