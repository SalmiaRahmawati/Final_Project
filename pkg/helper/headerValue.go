package helper

import (
	"github.com/gin-gonic/gin"
)

var AppJson = "application/json"

func GetContentType(ctx *gin.Context) string {
	return ctx.Request.Header.Get("Content-Type")
}
