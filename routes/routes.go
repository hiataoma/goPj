package routes

import (
	"blog_go/controller"
	"blog_go/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	//使用全局的跨域
	r.Use(middleware.CORSMiddleware())
	//使用路由组进行设置
	v1 := r.Group("/api/v1")
	{
		v1.POST("/auth/register", controller.Register)
		v1.POST("/auth/login", controller.Login)
		v1.GET("/auth/info", middleware.AuthMiddleware(), controller.Info)
	}
	//r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//使用中间件
	//使用权限认证，只有有权限人才能调用此接口
	//POST接口
	//采用路由组的形式实现
	//采用权限认值，如果权限不足没有token无法获取此接口
	//在路由组中使用中间件
	v2 := r.Group("/api/v2")
	// 使用gin自带的日志

	v2.Use(gin.Logger())
	{
		v2.GET("/testget/:name", controller.Testget)
		v2.POST("/testpost", controller.TestPost)
		v2.POST("/upload", controller.Upload)
		v2.POST("/uploads", controller.Uploads)
	}
	return r
}
