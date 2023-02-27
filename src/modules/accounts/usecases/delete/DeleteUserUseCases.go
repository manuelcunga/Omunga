package usecases

import (
	"errors"

	repositories "github.com/manuelcunga/Omunga/src/modules/accounts/repositories/interfaces"
)

type DeleteUserUseCases struct {
	userRepository repositories.ICreateUserRepository
}

func NewDeleteUserUseCases(userRepo repositories.ICreateUserRepository) DeleteUserUseCases {
	return DeleteUserUseCases{userRepository: userRepo}
}

func (deleteUseruseCase *DeleteUserUseCases) Execute(id string) error {
	user, err := deleteUseruseCase.userRepository.FindById(id)

	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("Usuário não encontrado")
	}

	err = deleteUseruseCase.userRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
