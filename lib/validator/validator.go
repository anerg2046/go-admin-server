package validator

//将验证器错误翻译成中文

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translation "github.com/go-playground/validator/v10/translations/zh"
)

var (
	trans ut.Translator
)

func NewValidator() {
	trans, _ = ut.New(zh.New()).GetTranslator("zh")
	Validate := binding.Validator.Engine().(*validator.Validate)
	Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})
	translation.RegisterDefaultTranslations(Validate, trans)
}

func Error(err error) (ret map[string][]string) {
	if _, ok := err.(validator.ValidationErrors); ok {
		ret = make(map[string][]string)
		for _, e := range err.(validator.ValidationErrors) {
			ret[e.StructField()] = append(ret[e.StructField()], e.Translate(trans))
		}
	}
	return
}
