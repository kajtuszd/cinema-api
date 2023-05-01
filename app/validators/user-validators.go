package validators

import (
	"github.com/go-playground/validator/v10"
	"unicode"
)

func PasswordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) > 6 {
		for _, c := range password {
			if unicode.IsDigit(c) {
				return true
			}
		}
	}
	return false
}
