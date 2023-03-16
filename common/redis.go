package common

// RedisKeyConfiguration 存储在 Redis 中的键配置结构体
type RedisKeyConfiguration struct {
	PrefixTag            string `json:"prefix-tag"`
	TokenKeyPrefix       string `json:"token-key-prefix"`
	TokenExpireKeyPrefix string `json:"token-expire-key-prefix"`
}

// RedisKeys 实例化配置
var RedisKeys = RedisKeyConfiguration{
	PrefixTag:            ":",           // 分隔符
	TokenKeyPrefix:       "Token",       // 用户 Token Key 前缀
	TokenExpireKeyPrefix: "TokenExpire", // 用户 Token 超时 Key 前缀
}
