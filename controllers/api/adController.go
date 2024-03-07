package api

import (
	"dcard-golang-project/middlewares"

	"github.com/gin-gonic/gin"
)

type AD struct{}

func (a AD) Admin(c *gin.Context) {
	panic(&(middlewares.ValidationError{Message: "Validation error"}))
}

func (a AD) Broadcast(c *gin.Context) {
	c.String(200, "hi, Broadcast controller says")
}
