package validations

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ValidationErrors struct {
	Property string `json:"Property"`
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Message  string `json:"message"`
}

func GetValidationErrors(err error) *[]ValidationErrors {
	var validationErrors []ValidationErrors

	var ve validator.ValidationErrors

	if errors.As(err, &ve){
		for _, err := range err.(validator.ValidationErrors){
			var el ValidationErrors 
			el.Property = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			validationErrors = append(validationErrors, el)
		}
		return &validationErrors
	}
	return nil

}
