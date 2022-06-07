package model

//创建一个数据库表  通过这种形式创建一个表
//type Test struct {
//	gorm.Model
//	Name      string `gorm:"type:varchar(20);not null"`
//	Telephone string `gorm:"type:varchar(11);not null;unique"`
//	Password  string `gorm:"size:255;not null"`
//}

// 如果定义字段在数据库中定义后可以查看出来

//数据库查出多个表中的数据
type Test struct {
	Name string `json:"name" form:"nameform"`
	Age  int    `json:"age" form:"ageform"`
	Aaaa string `json:"aaaa" form:"aaaaform"`
}

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
