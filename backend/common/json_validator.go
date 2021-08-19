package common

import (
	"github.com/go-playground/validator"
	"regexp"
)

var JsonValidator = newValidator()

func newValidator() *validator.Validate {
	v := validator.New()

	emtErr := v.RegisterValidation("NotBlank", func(fl validator.FieldLevel) bool {
		re := regexp.MustCompile(`^[\s\t\n]+$`)
		str := fl.Field().String()

		if len(str) == 0 {
			return false
		}

		matches := re.FindAllString(str, -1)

		if len(matches) >= 1 {
			return false
		}
		return true
	})

	if emtErr != nil {
		panic(emtErr)
	}
	err := v.RegisterValidation("IsTag", func(fl validator.FieldLevel) bool {
		m := map[string]bool{
			"Book":   true,
			"Course": true,
			"Video":  true,
			"Pet":    true,
		}
		_, ok := m[fl.Field().String()]
		return ok
	})

	if err != nil {
		panic(err)
	}
	return v
}
