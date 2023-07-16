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
		StatusCode:   string(customError.Type()),
		ReasonPhrase: string(customError.Causes()[0].Code),
		Errors:       customError.Error(),
	}
}

func translateErrorCode(code errors.Error) int {
	switch code.Type() {
	case errors.NotFound:
		return http.StatusNotFound
	case errors.BadRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
