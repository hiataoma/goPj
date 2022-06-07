package test

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/test", test1)
}
