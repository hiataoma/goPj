package common

//连接数据库
import (
	"blog_go/model"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {

	driverName := viper.GetString("datasource.driverName")
	//fmt.Print(driverName)
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	//全局禁用表明复数
	db.SingularTable(true)
	//数据库中自动创建表
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
