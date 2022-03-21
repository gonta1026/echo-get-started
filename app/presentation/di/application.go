package di

type application struct {
	usecase usecase
}

func initializeApplication(infra infrastructure) application {
	return application{
		usecase: initializeUsecase(infra),
	}
}
