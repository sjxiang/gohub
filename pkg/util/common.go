package util

import (
	"fmt"
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





// 将 time.Duration（nano seconds 为单位）输出为小数点后 3 位数 ms（mircosecond 毫秒，千分之一秒）
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3f ms", float64(elapsed.Nanoseconds())/1e6)
}

