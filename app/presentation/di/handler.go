package di

import (
	"echo-get-started/app/presentation/handler"
)

type Handlers struct {
	User handler.User
}

func InitializeHandler() *Handlers {

	infra := initializeInfrastructure()
	application := initializeApplication(infra)

	return &Handlers{
		User: *handler.NewUser(
			application.usecase.userUsecases.UsersAccessor,
		),
	}
}
