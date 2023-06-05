package registry

import (
	"go-auth/pkg-v2/adapter/controller"
	"go-auth/pkg-v2/adapter/repository"
	"go-auth/pkg-v2/usecase/usecase"
)

func (r *registry) NewUserController() controller.User {
	u := usecase.NewUserUsecase(
		repository.NewUserRepository(r.db),
		repository.NewDBRepository(r.db),
	)

	return controller.NewUserController(u)
}
