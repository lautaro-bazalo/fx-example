package user

type Repository interface {
	Create(user *User) error
	GetByID(id uint64) (*User, error)
}

type repository struct {
	dbRepo DBRepository
}

func (r *repository) Create(user *User) error {
	return r.dbRepo.Create(user)
}

func (r *repository) GetByID(id uint64) (*User, error) {
	user, err := r.dbRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewRepository(dbRepo DBRepository) Repository {
	return &repository{
		dbRepo: dbRepo,
	}
}
