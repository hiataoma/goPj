package controller

import (
	"blog_go/common"
	"blog_go/dto"
	"blog_go/model"
	"blog_go/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//信息 获取用户信息

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	// 数据库去查询用户的信息
	// response.Success(http.StatusOK)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"user": dto.ToUserDto(user.(model.User)),
		},
	})
}

//注册功能实现

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	//获取表单参数的值
	
	var requestUser = model.User{}
	err := ctx.Bind(&requestUser)
	if err != nil {
		return 
	}
	//接收参数
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

	//数据验证.
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}

	//如果没有给名称随机给.
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	//判读手机号是否存在
	if isTelephoneExist(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已经存在"})
		return
	}
	//创建用户
	// 密码进行加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密失败"})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
	}
	//将一条记录添加到表中
	DB.Create(&newUser)

	token, err := common.ReleaseToken(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error:%v", err)
		return
	}
	//返回结果
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"token": token,
		},
		"msg": "注册成功",
	})
}

//注册功能的实现

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	//获取参数
	//使用map来获取参数
	//通过绑定来获取参数
	var requestUser = model.User{}
	err := ctx.Bind(&requestUser)
	if err != nil {
		return
	}

	telephone := requestUser.Telephone

	password := requestUser.Password

	//数据验证

	//号码验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	//密码验证
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	//判断手机号是否存在
	var user model.User

	// user 返回这条记录
	DB.Where("telephone = ?", telephone).First(&user)

	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error:%v", err)
		return
	}

	//返回结果
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"token": token,
		},
		"msg": "登陆成功",
	})
}

//get接口实现

//func Test(ctx *gin.Context) {
//	//接口测试
//	ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "测试接口"})
//}

//post接口实现

//func TestPos1(ctx *gin.Context) {
//	//
//	//var requestUser = model.User{}
//	//ctx.Bind(&requestUser)
//	//接口测试
//	//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "测试post接口"})
//}

//查询手机号是否存在
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	//测试.
	var user model.User

	db.Where("telephone = ?", telephone).First(&user)

	if user.ID != 0 {
		return true
	}
	return false
}
