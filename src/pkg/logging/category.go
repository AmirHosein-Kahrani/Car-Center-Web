package logging

type Category string
type SubCategory string
type ExtraKey string

const (
	General         Category = "General"
	Internal        Category = "internal"
	Postgres        Category = "Postgres"
	Redis           Category = "Redis"
	Validation      Category = "Validation"
	RequestResponse Category = "RequestResponse"
)

const (
	//General
	StartUp          SubCategory = "StartUp"
	ExternalServices SubCategory = "ExternalServices"

	// Postgres
	Select   SubCategory = "Select"
	Rollback SubCategory = "Rollback"
	Update   SubCategory = "Update"
	Delete   SubCategory = "Delete"
	Insert   SubCategory = "Insert"

	// Internal
	Api                 SubCategory = "Api"
	HashPassword        SubCategory = "HashPassword"
	DefaultRoleNotFound SubCategory = "DefaultRoleNotFound"

	// Validation
	MobileValidation   SubCategory = "MobileValidation"
	PasswordValidation SubCategory = "PasswordValidation"
)

const (
	AppName      ExtraKey = "AppName"
	LoggerName   ExtraKey = "LoggerName"
	ClientIP     ExtraKey = "ClientIP"
	HostIp       ExtraKey = "HostIp"
	Method       ExtraKey = "Method"
	StatusCode   ExtraKey = "StatusCode"
	BodySize     ExtraKey = "BodySize"
	Path         ExtraKey = "Path"
	Latency      ExtraKey = "Latency"
	Body         ExtraKey = "Body"
	ErrorMessage ExtraKey = "ErrorMessage"
	RequestBody  ExtraKey = "RequestBody"
	ResponseBody ExtraKey = "ResponseBody"
)
