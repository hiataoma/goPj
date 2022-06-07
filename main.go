package main

import (
	"blog_go/common"
	"blog_go/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

// @title bolg_go
// @version 1.0
// @description 博客项目
// @termsOfService http://swagger.io/terms/

// @contact.name marvin
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 这里写接口服务的host
// @BasePath 这里写base path
func main() {
	//go 入口文件.
	InitConfig() //读取配置文件
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	//r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))
	//routers.Include(test.Routers)
	//r := routers.Init()
	r = routes.CollectRoute(r)
	port := viper.GetString("server.port")
	r.Run(":" + port)
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
