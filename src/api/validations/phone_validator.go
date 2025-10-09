package validations

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func IranianPhone_validator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)

	if !ok {
		return false
	}

	result, err := regexp.MatchString(`^0913\d*$`, value)
	if err != nil {
		log.Print(err.Error())
	}
	return result
}
