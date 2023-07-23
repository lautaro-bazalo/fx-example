package userfx

import (
	modelUser "fxdemo/internal/pkg/model/user"
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

func provideUserDBRepository(db *gorm.DB) modelUser.DBRepository {
	return modelUser.NewDBRepository(db)
}

func provideUserRepository(dbRepo modelUser.DBRepository) modelUser.Repository {
	return modelUser.NewRepository(dbRepo)
}

func provideUserUseCase(userRepo modelUser.Repository) user.UseCase {
	return user.NewUseCase(userRepo)
}

func provideUserRouter(userCtrl user.UseCase) user.Router {
	return user.NewRouter(userCtrl)
}
