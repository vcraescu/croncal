package api

import (
	"gopkg.in/go-playground/validator.v9"
	enTranslations "gopkg.in/go-playground/validator.v9/translations/en"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
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

	enTranslations.RegisterDefaultTranslations(v, validatorTranslator)

	return &entityValidator{validator: v}
}
