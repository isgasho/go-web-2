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
	// 如果没有报错，则直接跳过
	if err == nil {
		return nil
	}

	// 判断错误是否是校验错误
	errs, ok := err.(validator.ValidationErrors)
	if ok {
		// 遍历错误，可能有多个字段报错
		for _, e := range errs {
			tran := e.Translate(common.Translator)
			// 字段名称
			fd := e.Field()

			// 获取自定义的错误
			v, isExist := fieldError[fd]
			if isExist {
				return errors.New(v)
			}

			// 如果不是自定义的，则是系统的错误，就需要重写字段
			v, isExist = fieldTrans[fd]
			if isExist {
				// 需要进行翻译
				return errors.New(strings.Replace(tran, e.Field(), v, -1))
			}
			return errors.New(tran)
		}
	}
	return nil
}

// ShouldBindJSON 重写参数绑定方法，在内部增加字段校验
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
