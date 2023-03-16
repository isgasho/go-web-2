package gedis

import (
	"go-web/common"
)

// StringResult 字符串类型的返回
type StringResult struct {
	Result string
	Error  error
}

// NewStringResult 构造函数
func NewStringResult(result string, error error) *StringResult {
	return &StringResult{Result: result, Error: error}
}

// Unwrap 获取值
func (this *StringResult) Unwrap() string {
	if this.Error != nil {
		common.Logger.Debug("缓存未命中：", this.Error.Error())
	}
	return this.Result
}

// UnwrapWithDefault 获取值，如果不存在，则返回指定的默认值
func (this *StringResult) UnwrapWithDefault(v string) string {
	if this.Error != nil {
		return v
	}
	return this.Result
}

// UnwrapWithFunc 获取值，如果不存在，则触发指定函数
func (this *StringResult) UnwrapWithFunc(f func() string) string {
	if this.Error != nil {
		return f()
	}
	return this.Result
}
