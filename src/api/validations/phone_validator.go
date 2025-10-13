package validations

import (
	"regexp"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/logging"
	"github.com/go-playground/validator/v10"
)

var logger = logging.NewLogger(config.GetConfig())

func IranianPhone_validator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)

	if !ok {
		return false
	}

	result, err := regexp.MatchString(`^0913\d*$`, value)
	if err != nil {
		logger.Fatal(logging.Validation, logging.MobileValidation, err.Error(), nil)
	}
	return result
}
