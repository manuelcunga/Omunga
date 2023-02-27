package usecases

import (
	"errors"

	"github.com/manuelcunga/Omunga/src/modules/accounts/infra/entities"
	repositories "github.com/manuelcunga/Omunga/src/modules/accounts/repositories/interfaces"
)

type FindOndeUserUseCases struct {
	userRepository repositories.ICreateUserRepository
}

func NewFindOndeUserUseCases(userRepo repositories.ICreateUserRepository) FindOndeUserUseCases {
	return FindOndeUserUseCases{userRepository: userRepo}
}

func (user *FindOndeUserUseCases) FindOne(userId string) (*entities.User, error) {
	foundUser, err := user.userRepository.FindById(userId)
	if err != nil {
		return nil, err
	}

	if foundUser == nil {
		return nil, errors.New("Usuário não encontrado")
	}

	return foundUser, nil
}
