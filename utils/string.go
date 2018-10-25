package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// RandmonCode 生成指定位数随机数
func RandmonCode(bitNumber int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var res string
	for i := 0; i < bitNumber; i++ {
		res += fmt.Sprintf("%d", r.Intn(10))
	}

	return res
}

//RandomString 生成随机字符串
func RandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := make([]byte, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
