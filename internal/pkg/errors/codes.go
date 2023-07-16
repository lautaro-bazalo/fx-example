package errors

type ErrorCode string

const (
	// Generic
	InternalServerErrorCode ErrorCode = "internal_server_error"
	DecodingJsonCode        ErrorCode = "json_decoding_error"
	SQLError                ErrorCode = "sql_error"
	// UserTarget
	UserNotFoundCode      ErrorCode = "user_not_found"
	UserAlreadyExistsCode ErrorCode = "user_already_exists"
	UserInvalidData       ErrorCode = "user_invalid_data"
)
