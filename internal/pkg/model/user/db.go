package user

import (
	internalErrors "fxdemo/internal/pkg/errors"
	"gorm.io/gorm"
	"strconv"
)

type DBRepository interface {
	GetByID(id uint64) (*User, internalErrors.Error)
	GetByName(userName string) (*User, internalErrors.Error)
	Create(user *User) internalErrors.Error
	UpdateByName(user *User) (*User, internalErrors.Error)
}

type dbRepo struct {
	db *gorm.DB
}

func NewDBRepository(db *gorm.DB) DBRepository {
	return &dbRepo{db: db}
}

func (db *dbRepo) Create(user *User) internalErrors.Error {
	if err := db.db.Create(user).Error; err != nil {
		return customError(user.Name, err)
	}
	return nil
}

func (db *dbRepo) GetByID(id uint64) (*User, internalErrors.Error) {
	u := &User{}
	if err := db.db.First(u, id).Error; err != nil {
		return nil, customError(strconv.FormatUint(id, 10), err)
	}
	return u, nil
}

func (db *dbRepo) UpdateByName(userDB *User) (*User, internalErrors.Error) {

	if err := db.db.Where("user.name = ?", userDB.Name).Updates(userDB).Error; err != nil {
		return nil, customError(userDB.Name, err)
	}

	return userDB, nil
}

func (db *dbRepo) GetByName(userName string) (*User, internalErrors.Error) {
	u := &User{}
	if err := db.db.First(u, "user.name = ?", userName).Error; err != nil {
		return nil, customError(userName, err)
	}
	return u, nil
}
func customError(userName string, err error) internalErrors.Error {
	switch err {
	case gorm.ErrRecordNotFound:
		return internalErrors.NotFound.NewErrorF(internalErrors.NewErrorCausef(internalErrors.UserNotFoundCode, "User with name userName: %v not found. %v", userName, err))
	case gorm.ErrInvalidData:
		return internalErrors.BadRequest.NewErrorF(internalErrors.NewErrorCausef(internalErrors.UserInvalidData, err.Error()))
	case gorm.ErrMissingWhereClause:
		return internalErrors.InternalError.NewErrorF(internalErrors.NewErrorCausef(internalErrors.SQLError, err.Error()))
	default:
		return internalErrors.InternalError.NewErrorF(internalErrors.NewErrorCausef(internalErrors.InternalServerErrorCode, err.Error()))
	}
}
