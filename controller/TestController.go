package controller

import (
	"blog_go/common"
	"blog_go/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
type Mahaitao struct {
	Name string `form:"name" json:"name" binding:"required"`
	Age string `form:"age" json:"age" binding:"required"`
}
func Testget(ctx *gin.Context) {
	var mahaitao Mahaitao
	if err := ctx.ShouldBind(&mahaitao); err == nil {
		//fmt.Println(mahaitao.Name)
		//fmt.Println(mahaitao.Age)
	}
	//fmt.Println(mahaitao)
	// 给query设置一个默认值
	//firstname := ctx.DefaultQuery("name", "Guest")

	// get获取某一条query的值
	//name := ctx.Query("name") // 是 c.Request.URL.Query().Get("lastname") 的简写

	// 返回字符串格式
	//ctx.String(http.StatusOK, firstname)

	// 返回状态
	//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": firstname})
	DB := common.GetDB()

	//var user model.User

	// 设置默认值
	//DB.Where("name = ?", name).Find(&user)
	//DB.Where(&User{Name: "willing"}).First(&aaa)
	//增删改查
	firstname := ctx.Query("name")
	fmt.Println(firstname)

	// 定义数据格式
	//UserDto
	var test model.User // 定义查询的表
	//var user dto.UserDto // 定义查询的表
	//DB.Take(&test)
	//DB.Find(&test)

	// 语句查询
	DB.Where("name = ?", firstname).Find(&test)

	ctx.JSON(http.StatusOK, gin.H{"code": 422,"msg": test})
	//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": test})
}

func TestPost(ctx *gin.Context) {
	//接口测试
	//获取表单参数
	//name := ctx.PostForm("name")
	//age := ctx.PostForm("age")
	//ctx.JSON(http.StatusOK, name+age)
	//var test model.Test

	var persion model.Person
	if ctx.Bind(&persion) == nil {
		//ctx.JSON(http.StatusOK, name+age)
		//接口传的参数与接收的方式是不一样的
		//接口传参的问题
		ctx.JSON(http.StatusOK, gin.H{
			"user": persion.Name,
			"pwd":  persion.Address,
		})
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": test.Name, "aa": test.Age})
	}
	//ctx.String(http.StatusOK, "测试")
	//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "测试post接口"})

}

