package middleware

import (
	"errors"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go-web/common"
	"go-web/model"
	"go-web/pkg/dto"
	"go-web/pkg/response"
	"gorm.io/gorm"
	"time"
)

// JWTAuth JWT 认证中间件
func JWTAuth() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:      common.Config.JWT.Realm,                                 // JWT 标识
		Key:        []byte(common.Config.JWT.Key),                           // JWT 服务的密钥
		Timeout:    time.Hour * time.Duration(common.Config.JWT.Timeout),    // 超时时间
		MaxRefresh: time.Hour * time.Duration(common.Config.JWT.MaxRefresh), // Token 刷新时间
		// Login 中间件
		Authenticator: authenticator, // 用户认证
		PayloadFunc:   payloadFunc,   // Token 封装
		LoginResponse: loginResponse, // 登录成功相应
		Unauthorized:  unauthorized,  // 登录失败相应
		// Middleware 中间件
		IdentityHandler: identityHandler, // 解析 Token
		Authorizator:    authorizator,    // 验证 Token
		// Logout 中间件
		LogoutResponse: logoutResponse,                                     // 退出登录
		TokenLookup:    "header: Authorization, query: token, cookie: jwt", // Token 查找的字段
		TokenHeadName:  "Bearer",                                           // Token 请求头名称
		TimeFunc:       time.Now,
	})
}

// 隶属 Login 中间件，用于用户登录认证，替换自己写的认证函数，当调用 LoginHandler 就会触发。
// 通过从 ctx 中检索出数据进行验证，最终返回包含用户信息的 Map 或者 Struct
func authenticator(ctx *gin.Context) (interface{}, error) {
	// 获取用户传递的数据
	var req dto.Login
	_ = ctx.ShouldBindJSON(&req)

	// 查询用户信息
	var user model.User
	result := common.DB.Where("username = ? and password = ?", req.Username, req.Password).First(&user)

	// 判断用户是否存在
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New(response.UserLoginErrorMessage)
	} else {
		if *user.Status == 0 {
			// 判断用户是否被禁用
			return nil, errors.New(response.UserDisableMessage)
		} else if *user.Locked == 1 {
			// 判断用户是否被锁定
			return nil, errors.New(response.UserLockedMessage)
		} else {
			// 组装返回数据
			data := map[string]interface{}{
				"user": map[string]interface{}{
					"id":       user.Id,
					"username": user.Username,
					"nickname": user.Nickname,
				},
			}
			// 此处返回的数据会被传递给 PayloadFunc 函数继续处理
			return data, nil
		}
	}
}

// 接收 Authenticator 验证成功后传递过来的数据，进行封装成 Token
// MapClaims 必须包含 IdentityKey
// MapClaims 会被嵌入 Token 中，后续可以通过 ExtractClaims 对 Token 进行解析获取到
func payloadFunc(data interface{}) jwt.MapClaims {
	// 获取数据进行重新封装
	v, ok := data.(map[string]interface{})
	if ok {
		return jwt.MapClaims{
			jwt.IdentityKey: "JWTIdentityKey",
			"user":          v["user"],
		}
	}
	return jwt.MapClaims{}
}

// 接收 PayloadFunc 传递过来的 Token 信息，返回登录成功
func loginResponse(ctx *gin.Context, code int, token string, expire time.Time) {
	// 相应请求
	response.SuccessWithData(map[string]interface{}{
		"token":  token,
		"expire": expire,
	})
}

// 认证失败
func unauthorized(ctx *gin.Context, code int, message string) {
	response.FailedWithCodeAndMessage(code, message)
}

// 解析 Token，验证 Token 是否合法
func identityHandler(ctx *gin.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return map[string]interface{}{
		"IdentityKey": claims[jwt.IdentityKey],
		"user":        claims["user"],
	}
}

// 验证 Token
func authorizator(data interface{}, ctx *gin.Context) bool {
	v, ok := data.(map[string]interface{})
	if ok {
		ctx.Set("user", v["user"])
		return true
	}
	return false
}

// 退出登录
func logoutResponse(ctx *gin.Context, code int) {
	response.Success()
}
