package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func test1(ctx *gin.Context) {
	ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "路由工程化测试"})
}
