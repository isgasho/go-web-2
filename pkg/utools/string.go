package utools

import (
	"fmt"
	"strconv"
)

// String2Uint 字符串转 uint
func String2Uint(str string) uint {
	num, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0
	}
	return uint(num)
}

// Uint2String Uint 转字符串转
func Uint2String(num uint) string {
	return fmt.Sprintf("%d", num)
}
