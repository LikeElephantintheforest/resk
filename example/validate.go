package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
)

type UserVL struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       uint8  `validate:"gte=0,lte=130"`
	Email     string `validate:"required,email"`
}

func main() {

	userVl := &UserVL{
		FirstName: "firstName",
		LastName:  "lastName",
		Age:       131,
		Email:     "qq.com",
	}

	validate := validator.New()
	//创建翻译器，用来转义错误
	cn := zh.New()
	uni := ut.New(cn, cn)
	trans, found := uni.GetTranslator("zh")

	if found {
		err := vtzh.RegisterDefaultTranslations(validate, trans)
		if nil != err {
			fmt.Print("翻译器装配失败")
		}
	} else {
		fmt.Print("查询翻译器失败")
	}

	err := validate.Struct(userVl)
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			//npe check
			fmt.Print(err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, err := range errs {
				fmt.Println(err.Translate(trans))
			}
		}
	}

}
