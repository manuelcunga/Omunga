package usecases

import (
	"errors"

	"github.com/manuelcunga/Omunga/src/modules/accounts/dtos"
	"github.com/manuelcunga/Omunga/src/modules/accounts/infra/entities"
	repositories "github.com/manuelcunga/Omunga/src/modules/accounts/repositories/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type UpdateUserUseCases struct {
	userRepository repositories.ICreateUserRepository
}

func NewUpdateUserUseCases(userRepo repositories.ICreateUserRepository) UpdateUserUseCases {
	return UpdateUserUseCases{userRepository: userRepo}
}

func (userUseCase *UpdateUserUseCases) Execute(id string, data *dtos.UpdateUserDTO) (*entities.User, error) {
	user, err := userUseCase.userRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("Usuário não encontrado")
	}

	if data.FirstName != "" {
		user.FirstName = data.FirstName
	}
	if data.LastName != "" {
		user.LastName = data.LastName
	}
	if data.Email != "" {
		user.Email = data.Email
	}
	if data.Phone != "" {
		user.Phone = data.Phone
	}
	if data.Bio != "" {
		user.Bio = data.Bio
	}

	if data.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}

	err, updatedUser := userUseCase.userRepository.Update(user.ID, user)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}
