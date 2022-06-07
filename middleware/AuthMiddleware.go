package middleware

import (
	"blog_go/common"
	"blog_go/model"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//权限控制的中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		/**
		对于用户权限的认证
		*/
		//获取authorzation header
		//获取token的头部信息
		//获取调用接口的头部token，判断是否有权限进行访问
		tokenString := ctx.GetHeader("Authorization")

		//validate token formate  token 是否有效
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		//截取token后7位
		tokenString = tokenString[7:]

		//将token进行转化
		token, claims, err := common.ParseToken(tokenString)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		//验证通过后claim中的userId   取出userid用户的id
		userId := claims.UserId

		//打开数据库
		DB := common.GetDB()

		//检查数据库是否存在当前用户
		var user model.User
		DB.First(&user, userId)

		//用户
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		fmt.Print(ctx)
		//用户存在 将user的信息写入上下文
		ctx.Set("user", user)

		fmt.Print(ctx)
		//如何实现
		ctx.Next()
	}

}
