package util

import (
	"time"
	mathrand "math/rand"
	// "crypto/rand"
)

// RandStringRunes 生成长度为 length 随机数字字符串
func RandStringRunes(length int) string {
	mathrand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")  

	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[mathrand.Intn(len(letterRunes))]
	}
	
	return string(b)
}



