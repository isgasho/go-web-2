package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 响应状态码
const (
	OK                  = 200
	NotOK               = 400
	Unauthorized        = 401
	Forbidden           = 403
	ParamError          = 406
	InternalServerError = 500
	UserLoginError      = 1001
	UserDisable         = 1002
)

// 状态码对应的信息
const (
	OKMessage                  = "操作成功"
	NotOKMessage               = "操作失败"
	UnauthorizedMessage        = "登录过期，请重新登录"
	ForbiddenMessage           = "无权限访问该资源"
	ParamErrorMessage          = "参数错误"
	InternalServerErrorMessage = "服务器内部错误，请联系管理员"
	UserLoginErrorMessage      = "用户名或密码错误"
	UserDisableMessage         = "用户已经被禁用，请联系管理员"
)

// CustomMessage 状态码和信息绑定
var CustomMessage = map[int]string{
	OK:                  OKMessage,
	NotOK:               NotOKMessage,
	Unauthorized:        UnauthorizedMessage,
	Forbidden:           ForbiddenMessage,
	ParamError:          ParamErrorMessage,
	InternalServerError: InternalServerErrorMessage,
	UserLoginError:      UserLoginErrorMessage,
	UserDisable:         UserDisableMessage,
}

// JSON 封装 JSON 响应方法，状态码通过返回的数据给定
func JSON(ctx *gin.Context, resp interface{}) {
	ctx.JSON(http.StatusOK, resp)
}

// ResponseInfo 响应数据统一格式
type ResponseInfo struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Result 生成响应格式的数据
func Result(code int, status bool, data interface{}) {
	panic(ResponseInfo{
		Code:    code,
		Status:  status,
		Message: CustomMessage[code],
		Data:    data,
	})
}

// ResultWithMessage 生成自定义 message 的响应数据
func ResultWithMessage(code int, status bool, message string, data interface{}) {
	panic(ResponseInfo{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
	})
}

// Success 成功的响应
func Success() {
	Result(OK, true, map[string]interface{}{})
}

// SuccessWithData 成功的响应，带数据
func SuccessWithData(data interface{}) {
	Result(OK, true, data)
}

// Failed 失败的响应
func Failed() {
	Result(NotOK, true, map[string]interface{}{})
}

// FailedWithCode 失败的响应，带状态码
func FailedWithCode(code int) {
	Result(code, true, map[string]interface{}{})
}

// FailedWithMessage 失败的响应，带消息提示
func FailedWithMessage(message string) {
	ResultWithMessage(NotOK, true, message, map[string]interface{}{})
}

// FailedWithCodeAndMessage 失败的响应，带状态码和消息提示
func FailedWithCodeAndMessage(code int, message string) {
	ResultWithMessage(code, true, message, map[string]interface{}{})
}
