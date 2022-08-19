package util

import (
	"time"
	"math/rand"
)

// RandStringRunes 生成长度为 length 随机数字字符串
func RandStringRunes(length int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")  

	rand.Seed(time.Now().UnixNano())
	
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	
	return string(b)
}


