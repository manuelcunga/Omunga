package factory

import (
	"github.com/manuelcunga/Omunga/src/infra/database"
	repositories "github.com/manuelcunga/Omunga/src/modules/accounts/repositories/implementations"
	usecases "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/create"
	findAllUsers "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/findAll"
	findOneUsers "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/findOne"
	updateUser "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/update"
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

func NewFindOneUserUseCases() findOneUsers.FindOndeUserUseCases {
	repo := repositories.NewUserRepository(database.Connection())
	return findOneUsers.NewFindOndeUserUseCases(repo)
}

func NewUpdateUserUsecases() updateUser.UpdateUserUseCases {
	repo := repositories.NewUserRepository(database.Connection())
	return updateUser.NewUpdateUserUseCases(repo)
}
