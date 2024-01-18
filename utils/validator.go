package utils

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	trans ut.Translator
)

func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		lang := zh.New()
		uni := ut.New(lang, lang)
		trans, _ = uni.GetTranslator("zh")
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
	}
}

func GetValidatorTrans() ut.Translator {
	return trans
}
