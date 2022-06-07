package routers

// import (
// 	"blog_go/middleware"

// 	"github.com/gin-gonic/gin"
// )

// type Option func(*gin.Engine)

// var options = []Option{}

// // 注册app的路由配置
// func Include(opts ...Option) {
// 	options = append(options, opts...)
// }

// // 初始化
// func Init() *gin.Engine {
// 	r := gin.New()
// 	//使用全局跨域
// 	r.Use(middleware.CORSMiddleware())
// 	for _, opt := range options {
// 		opt(r)
// 	}
// 	return r
// }