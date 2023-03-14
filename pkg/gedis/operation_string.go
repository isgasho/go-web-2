package gedis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-web/common"
	"time"
)

// StringOperation 操作 String 类型数据
type StringOperation struct {
	Redis *redis.Client
	Ctx   context.Context
}

// NewStringOperation 构造函数
func NewStringOperation() *StringOperation {
	return &StringOperation{Redis: common.Redis, Ctx: context.Background()}
}

// Set 设置字符串的 Key/Value
func (this *StringOperation) Set(key string, value interface{}, attrs ...*OperationAttr) *InterfaceResult {
	// 获取过期时间，如果参数不存在，就设置为 0，永不过期
	expr := OperationAttrs(attrs).Find("expr").UnwrapWithDefault(time.Second * 0).(time.Duration)

	// 获取 NX 锁，两种锁不能同时存在
	nx := OperationAttrs(attrs).Find("nx").UnwrapWithDefault(nil)
	if nx != nil {
		return NewInterfaceResult(this.Redis.SetNX(this.Ctx, key, value, expr).Result())
	}

	// 获取 XX 锁
	xx := OperationAttrs(attrs).Find("xx").UnwrapWithDefault(nil)
	if xx != nil {
		return NewInterfaceResult(this.Redis.SetXX(this.Ctx, key, value, expr).Result())
	}

	// 如果没有锁则走默认设置
	return NewInterfaceResult(this.Redis.Set(this.Ctx, key, value, expr).Result())
}

// 设置用法：
// gedis.Set("key", "value", gedis.WithExpire(time.Second * 10), gedis.WithNX())

// Get 获取单个值
func (this *StringOperation) Get(key string) *StringResult {
	return NewStringResult(this.Redis.Get(this.Ctx, key).Result())
}

// Mget 获取多个值
func (this *StringOperation) Mget(keys ...string) *SliceResult {
	return NewSliceResult(this.Redis.MGet(this.Ctx, keys...).Result())
}
