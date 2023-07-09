package user

import (
	"fxdemo/api"
	"fxdemo/internal/pkg/errors"
	"fxdemo/internal/pkg/model/user"
	"github.com/mitchellh/mapstructure"
)

type UseCase interface {
	CreateUser(request api.CreateUserRequest) (*user.User, error)
	GetUserByID(id uint64) (*user.User, errors.Error)
}
type useCase struct {
	userRepo user.Repository
}

func NewUseCase(userRepo user.Repository) UseCase {
	return &useCase{userRepo: userRepo}
}

func (uc *useCase) CreateUser(request api.CreateUserRequest) (*user.User, error) {
	userDB := &user.User{}

	if err := mapstructure.Decode(&request, userDB); err != nil {
		return nil, err
	}

	if err := uc.userRepo.Create(userDB); err != nil {
		return nil, err
	}

	return userDB, nil
}

func (uc *useCase) GetUserByID(id uint64) (*user.User, errors.Error) {
	userDB, err := uc.userRepo.GetByID(id)
	if err != nil {
		customError := errors.NewError()
		if err := mapstructure.Decode(err, &customError); err != nil {
			return nil, errors.InternalError.NewErrorF(errors.NewErrorCausef(errors.DecodingJsonCode, "", err))
		}
		return nil, customError
	}
	return userDB, nil

}
