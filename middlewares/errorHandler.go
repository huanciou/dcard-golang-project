package middlewares

import "github.com/gin-gonic/gin"

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch e := err.(type) {
				case *ValidationError:
					c.JSON(400, gin.H{"error": e.Error()})
					c.Abort()
				case error:
					c.JSON(400, gin.H{"error": e.Error()})
					c.Abort()
				default:
					c.JSON(500, gin.H{"error": "Unknown Error"})
					c.Abort()
				}
			}
		}()
		c.Next()
	}
}
