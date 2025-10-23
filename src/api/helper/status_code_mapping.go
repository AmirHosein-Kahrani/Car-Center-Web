package helper

import (
	"net/http"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/pkg/service_errors"
)

var StatusCodeMapping = map[string]int{
	// Otp
	service_errors.OtpExist:    409,
	service_errors.OtpUsed:     409,
	service_errors.OtpNotValid: 400,
	// User
	service_errors.EmailExists:    409,
	service_errors.UserNameExists: 409,
	service_errors.RecordNotFound: 404,
}

func TranslateErrorToStatusCode(err error) int {
	value, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}
	return value
}
