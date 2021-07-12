package utils


import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"strings"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	en := en.New()
	uni = ut.New(en, en)

	trans, _ = uni.GetTranslator("en")

	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)
}

func ValidateRequest(data interface{}) map[string]string {

	validateEmptyString()


	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	err := validate.Struct(data)
	if err != nil {

		errs := err.(validator.ValidationErrors)

		var result = map[string]string{}

		for _, e := range errs {
			result[e.Field()] = e.Translate(trans)

		}

		return result
	}
	return nil
}

// this validates white space
func validateEmptyString()  {
	validate.RegisterValidation("notEmpty", func(fl validator.FieldLevel) bool {

		if len(fl.Field().String()) == 0 {
			return true
		}

		return len(strings.Trim(fl.Field().String(), " ")) > 0
	})

	validate.RegisterTranslation("notEmpty", trans, func(ut ut.Translator) error {
		return ut.Add("notEmpty", "{0} can not be empty", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("notEmpty", fe.Field())

		return t
	})
}
