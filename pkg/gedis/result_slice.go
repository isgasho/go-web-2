package gedis

import (
	"go-web/common"
)

// SliceResult 切片类型返回值
type SliceResult struct {
	Result []interface{}
	Error  error
}

// NewSliceResult 构造函数
func NewSliceResult(result []interface{}, error error) *SliceResult {
	return &SliceResult{Result: result, Error: error}
}

// Unwrap 获取值
func (this *SliceResult) Unwrap() []interface{} {
	if this.Error != nil {
		common.Logger.Warn("缓存未命中：", this.Error.Error())
	}
	return this.Result
}

// UnwrapWithDefault 获取值，没有返回默认值
func (this *SliceResult) UnwrapWithDefault(v []interface{}) []interface{} {
	if this.Error != nil {
		return v
	}
	return this.Result
}

// Iter 迭代
func (this *SliceResult) Iter() *Iterator {
	return NewIterator(this.Result)
}

// 迭代用法
//func demo() {
//	var conn = NewStringOperation()
//	var res = conn.Mget("name", "age", "gender").Iter()
//	for res.HasNext() {
//		fmt.Println(res.Next())
//	}
//}
