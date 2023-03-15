package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-web/common"
	"go-web/pkg/response"
	"strings"
)

type FieldTransInterface interface {
	FieldTrans() map[string]string
}

type FieldErrorsInterface interface {
	FieldErrors() map[string]string
}

// NewValidatorError 对验证错误进行重写
func NewValidatorError(err error, fieldTrans map[string]string, fieldError map[string]string) error {
	if err == nil {
		return nil
	}

	// 获取校验错误
	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		transStr := e.Translate(common.Translator)

		// 字段名称
		field := e.Field()

		// 自定义错误：先判断错误是否被重写
		v, ok := fieldError[field]
		if ok {
			// 返回自定义错误
			return errors.New(v)
		}

		// 系统错误：为重写错误信息的字段
		v, ok = fieldTrans[field]
		if ok {
			// 替换掉英文字段为中文
			return errors.New(strings.Replace(transStr, e.Field(), v, -1))
		}
		return errors.New(transStr)
	}
	return nil
}

// ShouldBindJSON 重写参数绑定方法，增加字段校验
func ShouldBindJSON(ctx *gin.Context, req interface{}) {
	err := ctx.ShouldBindJSON(req)

	// 判断是否有字段翻译
	fieldTrans := make(map[string]string, 0)
	if ft, ok := req.(FieldTransInterface); ok {
		fieldTrans = ft.FieldTrans()
	}

	// 判断是否有错误翻译
	fieldError := make(map[string]string, 0)
	if fe, ok := req.(FieldErrorsInterface); ok {
		fieldError = fe.FieldErrors()
	}

	// 对错误进行处理
	if err != nil {
		e := NewValidatorError(err, fieldTrans, fieldError)
		if e != nil {
			response.FailedWithMessage(e.Error())
			return
		}
		response.FailedWithMessage(err.Error())
		return
	}
}
