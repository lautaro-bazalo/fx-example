package user

import "fxdemo/internal/pkg/errors"

type Repository interface {
	Create(user *User) errors.Error
	UpdateByName(user *User) (*User, errors.Error)
	GetByID(id uint64) (*User, errors.Error)
	GetByName(userName string) (*User, errors.Error)
}

type repository struct {
	dbRepo DBRepository
}

func NewRepository(dbRepo DBRepository) Repository {
	return &repository{
		dbRepo: dbRepo,
	}
}
func (r *repository) Create(user *User) errors.Error {
	return r.dbRepo.Create(user)
}

func (r *repository) GetByID(id uint64) (*User, errors.Error) {
	user, err := r.dbRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) GetByName(userName string) (*User, errors.Error) {
	user, err := r.dbRepo.GetByName(userName)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *repository) UpdateByName(user *User) (*User, errors.Error) {
	return r.dbRepo.UpdateByName(user)
}
