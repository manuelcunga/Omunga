package factory

import (
	"github.com/manuelcunga/Omunga/src/infra/database"
	repositories "github.com/manuelcunga/Omunga/src/modules/accounts/repositories/implementations"
	usecases "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/create"
)

func NewUserUseCase() usecases.CreateUserUseCases {
	repo := repositories.NewUserRepository(database.Connection())
	return usecases.NewUserUseCase(repo)
}
