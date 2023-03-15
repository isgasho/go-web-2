package initialize

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"go-web/common"
	"log"
	"regexp"
)

// Validator 初始化校验器
func Validator() {
	// 中文翻译器
	zhTrans := zh.New()
	// 第一个是备用的语言，后面的才是支持的语言
	uni := ut.New(zhTrans, zhTrans)
	trans, _ := uni.GetTranslator("zh")
	v := validator.New()

	// 修改 gin 框架的 validator 引擎属性，实现自定义
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		// 注册自定义验证
		err := v.RegisterValidation("is-username", validateUsername)
		if err != nil {
			errMsg := "注册自定义验证失败：is-username"
			log.Fatalln(errMsg)
		}
	}

	_ = zhTranslations.RegisterDefaultTranslations(v, trans)
	common.Validate = v
	common.Translator = trans
	log.Println("校验器 v10 初始化完成！")
}

// 用户名认证
func validateUsername(fl validator.FieldLevel) bool {
	username, ok := fl.Field().Interface().(string)
	if ok {
		reg := `^[a-z][a-z0-9]{3,20}$`
		rgx := regexp.MustCompile(reg)
		return rgx.MatchString(username)
	}
	return false
}
