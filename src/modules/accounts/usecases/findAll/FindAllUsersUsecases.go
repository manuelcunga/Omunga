package usecases

import (
	"github.com/manuelcunga/Omunga/src/modules/accounts/infra/entities"
	repositories "github.com/manuelcunga/Omunga/src/modules/accounts/repositories/interfaces"
)

type FindAllUsersUseCases struct {
	userRepository repositories.ICreateUserRepository
}

func NewFindAllUsersUseCases(userRepo repositories.ICreateUserRepository) FindAllUsersUseCases {
	return FindAllUsersUseCases{userRepository: userRepo}
}

func (users *FindAllUsersUseCases) FindAll() ([]*entities.User, error) {
	userList, err := users.userRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return userList, nil
}
