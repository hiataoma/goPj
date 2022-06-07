package dto

import (
	"blog_go/model"
	"time"
)

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	CreatedAt time.Time `json: "createdAt"`
}

//创建一个用户DTO
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}

//
