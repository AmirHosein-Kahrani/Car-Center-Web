package validations

import (
	"github.com/go-playground/validator/v10"
)

func PasswordValidator(fld validator.FieldLevel) bool {

	value, ok := fld.Field().Interface().(string)

	if !ok {
		fld.Param()
		return false
	}
	return passValidating(value)
}

func passValidating(val string) bool {
	switch {
	case len(val) <= 7:
		return false
	case len(val) >= 12:
		return false
	default:
		return true
	}
}
