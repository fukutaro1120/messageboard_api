package controller

import "github.com/gin-gonic/gin"

func NewError(err error) interface{} {
	return gin.H{
		"error": err.Error(),
	}
}
