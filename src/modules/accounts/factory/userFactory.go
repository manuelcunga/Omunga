package factory

import (
	"github.com/manuelcunga/Omunga/src/infra/database"
	repositories "github.com/manuelcunga/Omunga/src/modules/accounts/repositories/implementations"
	usecases "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/create"
	findAllUsers "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/findAll"
)

func NewUserUseCase() usecases.CreateUserUseCases {
	repo := repositories.NewUserRepository(database.Connection())
	return usecases.NewUserUseCase(repo)
}

func NewListUsersUseCase() findAllUsers.FindAllUsersUseCases {
	repo := repositories.NewUserRepository(database.Connection())
	return findAllUsers.NewFindAllUsersUseCases(repo)
}

func NewListUsersUseCases() findAllUsers.FindAllUsersUseCases {
	repo := repositories.NewUserRepository(database.Connection())
	return findAllUsers.NewFindAllUsersUseCases(repo)
}
