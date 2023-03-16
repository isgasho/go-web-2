package gedis

import "go-web/common"

// InterfaceResult Interface 类型的数据返回
type InterfaceResult struct {
	Result interface{}
	Error  error
}

// NewInterfaceResult 构造函数
func NewInterfaceResult(result interface{}, error error) *InterfaceResult {
	return &InterfaceResult{Result: result, Error: error}
}

// Unwrap 获取值
func (this *InterfaceResult) Unwrap() interface{} {
	if this.Error != nil {
		common.Logger.Debug("缓存未命中：", this.Error.Error())
	}
	return this.Result
}

// UnwrapWithDefault 如果获取值不存在，则返回传入的值
func (this *InterfaceResult) UnwrapWithDefault(v interface{}) interface{} {
	if this.Error != nil {
		return v
	}
	return this.Result
}
