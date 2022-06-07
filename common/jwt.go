package common

import (
	"blog_go/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwt权限中间
var jwtKey = []byte("a_secret_crect")

// 设置过期时间
type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 发放token
func ReleaseToken(user model.User) (string, error) {
	//设置过期时间 7小时候过期
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	//配置文件
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(), //token发放的时间
			Issuer:    "haitaoma",
			Subject:   "user token",
		},
	}
	//生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//加密token
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//转换token 后台解析token是否正确
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
