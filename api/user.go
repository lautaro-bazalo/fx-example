package api

import (
	user2 "fxdemo/internal/pkg/model/user"
)

type CreateUserRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Age         int32  `json:"age"`
}

type CreateUserResponse struct {
	User user2.User
}

type updateUser struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Age         int32  `json:"age,omitempty"`
}
