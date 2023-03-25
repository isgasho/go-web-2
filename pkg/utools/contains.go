package utools

// ContainsUint 判断 Uint 数据是否在 Array 中
func ContainsUint(arr []uint, item uint) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}
