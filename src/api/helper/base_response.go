package helper

import "github.com/AmirHosein-Kahrani/Car-Center-Web/api/validations"

// "github.com/AmirHosein-Kahrani/Car-Center-Web/"

type BaseHttpResponse struct {
	Result           any                             `json:"result"`
	Succuss          bool                            `json:"success"`
	ResultCode       ResultCode                      `json:"resultCode"`
	ValidationErrors *[]validations.ValidationErrors `json:"validationErrors"`
	Error            any                             `json:"error"`
}

func GenerateBaseResponse(result any, success bool, resultCode ResultCode) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Succuss:    success,
		ResultCode: resultCode}
}

func GenerateBaseResponseWithError(result any, success bool, resultCode ResultCode, err error) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Succuss:    success,
		ResultCode: resultCode,
		Error:      err.Error()}

}

func GenerateBaseResponseWithValidationError(result any, success bool, resultCode ResultCode, err error) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Succuss:          success,
		ResultCode:       resultCode,
		ValidationErrors: validations.GetValidationErrors(err)}
}

func GenerateBaseResponseWithAnyError(result any, success bool, resultCode ResultCode, err interface{}) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Succuss:    success,
		ResultCode: resultCode,
		Error:      err,
	}
}
