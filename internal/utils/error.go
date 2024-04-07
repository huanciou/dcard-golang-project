package utils

import (
	"dcard-golang-project/middlewares"
	"errors"
)

func ErrorCheck() {
	err := errors.New("custom error for test")
	panic(&(middlewares.CustomizedError{Message: err.Error()}))
}
