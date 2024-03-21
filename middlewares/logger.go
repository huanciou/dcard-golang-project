package middlewares

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() *logrus.Logger {
	now := time.Now()
	logFilePath := ""

	/* os.Getwd() = pwd */
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/log/"
	}

	/* mkdir */
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}

	logFileName := now.Format("2006-01-02") + ".log"

	/* fileName = dirName + timestamp */
	fileName := path.Join(logFilePath, logFileName)

	/* make a file, if not exist returns an error */
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}

	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	logger := logrus.New()
	logger.Out = src

	logger.SetLevel(logrus.InfoLevel)

	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}

func LoggerToFile() gin.HandlerFunc {
	logger := Logger()
	return func(c *gin.Context) {

		startTime := time.Now()
		fmt.Println("logger here")
		c.Next()

		fmt.Println("logger continue")
		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)

		reqMethod := c.Request.Method

		reqUri := c.Request.RequestURI

		statusCode := c.Writer.Status()

		clientIP := c.ClientIP()

		/* throw a Info level log */
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
