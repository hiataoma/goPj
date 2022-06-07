package model

import (
	"github.com/jinzhu/gorm"
)

//创建一个数据库表  通过这种形式创建一个表
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}
