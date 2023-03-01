package factory

import (
	"github.com/manuelcunga/Omunga/src/infra/database"
	repositories "github.com/manuelcunga/Omunga/src/modules/accounts/repositories/implementations"
	usecases "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/create"
	deleteUser "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/delete"
	findAllUsers "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/findAll"
	findOneUsers "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/findOne"
	LoginUser "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/login"
	meUsecase "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/me"
	userProfile "github.com/manuelcunga/Omunga/src/modules/accounts/usecases/profile"
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

func NewDeleteUserUseCases() deleteUser.DeleteUserUseCases {
	repo := repositories.NewUserRepository(database.Connection())
	return deleteUser.NewDeleteUserUseCases(repo)
}

func NewLoginUseCases() LoginUser.LoginUseCases {
	repo := repositories.NewUserRepository(database.Connection())
	return LoginUser.NewLoginUseCases(repo)
}

func NewProfileUseCases() userProfile.ProfileUseCases {
	repo := repositories.NewUserRepository(database.Connection())
	return userProfile.NewProfileUseCases(repo)
}

func NewMeUseCases() meUsecase.MeUseCases {
	repo := repositories.NewUserRepository(database.Connection())
	return meUsecase.NewMeUseCases(repo)
}
