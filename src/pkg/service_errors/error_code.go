package service_errors

const (

	// Token
	UnexpectedError = "unexpected error"
	ClaimsNotFound  = "cslaims not found"
	TokenRequired   = "token required"
	TokenExpired    = "token expired"
	TokenInvalid    = "token invalid"

	// OTP
	OtpExist    = "otp_exist"
	OtpUsed     = "otp_used"
	OtpNotValid = "otp_not_valid"

	// User
	EmailExists    = "Email Exists"
	UserNameExists = "UserName Exists"
)
