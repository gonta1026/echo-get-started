package di

import (
	"echo-get-started/app/domain/model/user"
	"echo-get-started/app/infrastructure/dbconfig"
	"echo-get-started/app/infrastructure/persistence"
)

type infrastructure struct {
	repository repository
}

func initializeInfrastructure() infrastructure {
	return infrastructure{
		initializeRepository(),
	}
}

type repository struct {
	userRepository user.Repository
}

func initializeRepository() repository {
	return repository{
		persistence.NewUserPersistence(dbconfig.Master, dbconfig.Slave),
	}
}
