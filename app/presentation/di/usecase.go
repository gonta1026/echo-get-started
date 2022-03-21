package di

import (
	"echo-get-started/app/application/usecase/user"
)

type usecase struct {
	userUsecases userUsecases
}

type userUsecases struct {
	user.UsersAccessor
}

func initializeUsecase(infra infrastructure) usecase {
	return usecase{
		userUsecases: userUsecases{
			user.NewUsers(infra.repository.userRepository),
		},
	}
}
