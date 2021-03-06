/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-28 21:47:23
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-29 00:39:03
 */
package validator

import (
	"errors"
	"strings"
	"sync"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtrans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/joshua-chen/go-commons/mvc/context/response"
	"github.com/kataras/golog"

)

//
type Validator struct {
	Code       int
	Message    string
	Messages   map[string]string
	Err        error
	Validate   *validator.Validate
	Translator ut.Translator
}

var (
	instance *Validator
	lock     *sync.Mutex = &sync.Mutex{}
)

//
func Singleton() *Validator {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Validator{
				Code:    response.StatusValidateFailed,
				Message: "验证未通过"}
		}
	}
	return instance
}

//
func Instance() *Validator {
	return Singleton()
}

//
func Error(err error, args ...interface{}) {
	Instance().Error(err, args...)
}

//
func ErrorS(err string, code ...int) {
	Instance().ErrorS(err, code...)
}
func GetMessage(err error, prefix ...string) string {
	return Instance().GetMessage(err, prefix...)
}
func (e *Validator) Error(err error, args ...interface{}) {

	if len(args) > 0 {
		e.Code = (args[0]).(int)
	}
	prefix := ""
	if len(args) > 1 {
		prefix = (args[1]).(string)
	}

	//e.Err = err
	msg := GetMessage(err, prefix)
	e.Message = msg
	newErr := errors.New(e.Message)
	e.Err = newErr
	golog.Errorf("Error[%d]: %s", e.Code, e.Message)
	panic(e.Err)
}

//
func (e *Validator) ErrorS(errMsg string, code ...int) {
	if len(code) > 0 {
		e.Code = code[0]
	}
	e.Message = errMsg
	golog.Errorf("Error[%d]: %s ", e.Code, errMsg)
	err := errors.New(errMsg)
	e.Err = err
	panic(err)
}
func New() *validator.Validate {
	return Singleton().New()
}

//
func (e *Validator) New() *validator.Validate {

	trans, _ := GetTranslator() //获取需要的语言
	e.Validate = validator.New()
	e.Translator = trans
	zhtrans.RegisterDefaultTranslations(e.Validate, trans)
	return e.Validate
}

func GetTranslator(language ...string) (trans ut.Translator, found bool) {
	zh := zh.New() //中文翻译器
	en := en.New() //英文翻译器
	// 第一个参数是必填，如果没有其他的语言设置，就用这第一个
	// 后面的参数是支持多语言环境（
	// uni := ut.New(en, en) 也是可以的
	// uni := ut.New(en, zh, tw)
	uni := ut.New(en, zh)
	lang := "zh"
	if len(language) > 0 {
		lang = language[0]
	}
	return uni.GetTranslator(lang)
}
func (e *Validator) GetMessage(err error, prefix ...string) string {

	//e.Err = err
	errs := err.(validator.ValidationErrors)
	trans := e.Translator
	if trans == nil {
		trans, _ = GetTranslator()
	}

	msgs := removeStructName(errs.Translate(trans))

	msg := msgToString(msgs) // fmt.Sprintf("%v ", msgs)
	e.Message = msg
	e.Messages = msgs
	if len(prefix) > 0 {
		msg = prefix[0] + msg
	}
	return msg
}

func removeStructName(fields map[string]string) map[string]string {
	result := map[string]string{}

	for field, err := range fields {
		result[field[strings.Index(field, ".")+1:]] = err
	}
	return result
}

func msgToString(msgs map[string]string) string {
	result := ""

	for key, msg := range msgs {
		if key != "" {
			result += `'` + key + `'`
		}
		result += msg + "\n"
	}

	if result != "" {
		result = result[0 : len(result)-1]
	}
	return result
}
