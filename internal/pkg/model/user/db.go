package user

import (
	"errors"
	internalErrors "fxdemo/internal/pkg/errors"
	"gorm.io/gorm"
)

type DBRepository interface {
	GetByID(id uint64) (*User, error)
	Create(user *User) error
}

type dbRepo struct {
	db *gorm.DB
}

func (r *dbRepo) Create(user *User) error {
	return r.db.Create(user).Error

}

func (r *dbRepo) GetByID(id uint64) (user *User, err error) {
	u := &User{}
	if err := r.db.First(u, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, internalErrors.NotFound.NewErrorF(internalErrors.NewErrorCausef(internalErrors.UserNotFoundCode, "User id: %v not found", id))
		}
		return nil, internalErrors.InternalError.NewErrorF(internalErrors.NewErrorCausef(internalErrors.InternalServerErrorCode, err.Error()))
	}
	return u, nil
}

func NewDBRepository(db *gorm.DB) DBRepository {
	return &dbRepo{db: db}
}
