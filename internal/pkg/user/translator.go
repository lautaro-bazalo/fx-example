package user

import (
	"fxdemo/api"
	"fxdemo/internal/pkg/errors"
	"fxdemo/internal/pkg/model/user"
	"net/http"
)

func ToAPI(user user.User) api.CreateUserResponse {
	return api.CreateUserResponse{
		User: user,
	}
}

func toApiError(customError errors.Error) api.Error {
	return api.Error{
		StatusCode:   translateErrorCode(customError.Type()),
		ReasonPhrase: string(customError.Type()),
		Errors:       customError.Error(),
	}
}

func translateErrorCode(code errors.ErrorType) int {
	switch code {
	case errors.NotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
