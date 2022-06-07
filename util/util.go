package util

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

//随机生成数字
func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwetyuiopQWERTYUIOPASDFGHJKLZXCVBNM")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func test(n int) {
	fmt.Print("测试案例")
	log.Print(n)
}
