package base

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
	"xffresk/infra"
)

var validate *validator.Validate
var translator ut.Translator

func Validate() *validator.Validate {
	return validate
}

func Transtate() ut.Translator {
	return translator
}

type ValidatorStarter struct {
	infra.BaseStarter
}

func (v *ValidatorStarter) Init(ctx infra.StarterContext) {
	validate = validator.New()

	//创建翻译器，用来转义错误
	cn := zh.New()
	uni := ut.New(cn, cn)
	var found bool
	trans, found := uni.GetTranslator("zh")

	if found {
		err := vtzh.RegisterDefaultTranslations(validate, trans)
		if nil != err {
			logrus.Error("翻译器装配失败")
		}
	} else {
		logrus.Error("查询翻译器失败")
	}
	//
	//err := validate.Struct(userVl)
	//if err != nil {
	//	_, ok := err.(*validator.InvalidValidationError)
	//	if ok {
	//		//npe check
	//		fmt.Print(err)
	//	}
	//	errs, ok := err.(validator.ValidationErrors)
	//	if ok {
	//		for _, err := range errs {
	//			fmt.Println(err.Translate(trans))
	//		}
	//	}
	//}
}
