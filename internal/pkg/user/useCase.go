package user

import (
	"fxdemo/api"
	"fxdemo/internal/pkg/errors"
	"fxdemo/internal/pkg/model/user"
	"github.com/mitchellh/mapstructure"
)

type UseCase interface {
	CreateUser(request api.CreateUserRequest) (*user.User, errors.Error)
	GetUserByID(id uint64) (*user.User, errors.Error)
	updateByName(user api.CreateUserRequest) (*user.User, errors.Error)
	GetUserByName(userName string) (*user.User, errors.Error)
}
type useCase struct {
	userRepo user.Repository
}

func NewUseCase(userRepo user.Repository) UseCase {
	return &useCase{userRepo: userRepo}
}

func (uc *useCase) CreateUser(request api.CreateUserRequest) (*user.User, errors.Error) {
	userDB := &user.User{}

	if err := mapstructure.Decode(&request, userDB); err != nil {
		return nil, errors.InternalError.NewErrorF(errors.NewErrorCausef(errors.DecodingJsonCode, "", err))
	}

	if err := uc.userRepo.Create(userDB); err != nil {
		return nil, err
	}

	return userDB, nil
}

func (uc *useCase) GetUserByID(id uint64) (*user.User, errors.Error) {
	return uc.userRepo.GetByID(id)

}
func (uc *useCase) GetUserByName(userName string) (*user.User, errors.Error) {
	return uc.userRepo.GetByName(userName)
}

func (uc *useCase) updateByName(userRequest api.CreateUserRequest) (*user.User, errors.Error) {
	userDB := &user.User{}

	if err := mapstructure.Decode(&userRequest, userDB); err != nil {
		return nil, errors.InternalError.NewErrorF(errors.NewErrorCausef(errors.DecodingJsonCode, "", err))
	}

	return uc.userRepo.UpdateByName(userDB)
}
