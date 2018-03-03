package api

import (
	"gopkg.in/go-playground/validator.v9"
	enTranslations "gopkg.in/go-playground/validator.v9/translations/en"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"reflect"
	"strings"
	"github.com/gorhill/cronexpr"
)

var validatorTranslator ut.Translator

func init() {
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	validatorTranslator, _ = uni.GetTranslator("en")
}

type entityValidator struct {
	validator *validator.Validate
}

func (v *entityValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func newEntityValidator() *entityValidator {
	v := validator.New()
	v.RegisterValidation("interval", intervalFieldValidator)

	enTranslations.RegisterDefaultTranslations(v, validatorTranslator)
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	v.RegisterTranslation("required", validatorTranslator, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", strings.Title(fe.Field()))

		return t
	})

	v.RegisterTranslation("interval", validatorTranslator, func(ut ut.Translator) error {
		return ut.Add("interval", "{0} is invalid", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("interval", strings.Title(fe.Field()))

		return t
	})

	return &entityValidator{validator: v}
}

func intervalFieldValidator(fl validator.FieldLevel) bool {
	parts := strings.Split(fl.Field().String(), " ")
	if len(parts) < 5 {
		return false
	}

	for _, part := range parts {
		if !intervalFieldRangeValidator(part) {
			return false
		}
	}

	return true
}

func intervalFieldRangeValidator(v string) bool {
	if v == "" {
		return false
	}

	_, err := cronexpr.Parse(v)
	return err != nil
}
