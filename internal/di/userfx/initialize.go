package userfx

import (
	user3 "fxdemo/internal/pkg/model/user"
	"fxdemo/internal/pkg/user"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Provide(
	provideUserDBRepository,
	provideUserRepository,
	provideUserRouter,
	provideUserUseCase,
)

func provideUserDBRepository(db *gorm.DB) user3.DBRepository {
	return user3.NewDBRepository(db)
}

func provideUserRepository(dbRepo user3.DBRepository) user3.Repository {
	return user3.NewRepository(dbRepo)
}

func provideUserUseCase(userRepo user3.Repository) user.UseCase {
	return user.NewUseCase(userRepo)
}

func provideUserRouter(userCtrl user.UseCase) user.Router {
	return user.NewRouter(userCtrl)
}
