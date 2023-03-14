package gedis

import (
	"fmt"
	"time"
)

// OperationAttr 用于数据操作的参数设置
type OperationAttr struct {
	Name  string
	Value interface{}
}

// NewOperationAttr 构造函数
func NewOperationAttr(name string, value interface{}) *OperationAttr {
	return &OperationAttr{Name: name, Value: value}
}

// WithExpire 设置超时时间
func WithExpire(t time.Duration) *OperationAttr {
	return &OperationAttr{
		Name:  "expr",
		Value: t,
	}
}

// WithNX 设置 NX 锁，当 key 不存在时才能设置 key/value
func WithNX() *OperationAttr {
	return &OperationAttr{
		Name:  "nx",
		Value: struct{}{},
	}
}

// WithXX 设置 XX 锁，当 key 存在时才能设置 key/value
func WithXX() *OperationAttr {
	return &OperationAttr{
		Name:  "xx",
		Value: struct{}{},
	}
}

type OperationAttrs []*OperationAttr

// Find 查找传递的参数
func (this OperationAttrs) Find(name string) *InterfaceResult {
	for _, attr := range this {
		if attr.Name == name {
			return NewInterfaceResult(attr.Value, nil)
		}
	}
	return NewInterfaceResult(nil, fmt.Errorf("参数错误：", name))
}
