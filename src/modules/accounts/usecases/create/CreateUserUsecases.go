package usecases

import (
	"errors"

	"github.com/manuelcunga/Omunga/src/modules/accounts/dtos"
	"github.com/manuelcunga/Omunga/src/modules/accounts/infra/entities"
	repositories "github.com/manuelcunga/Omunga/src/modules/accounts/repositories/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserUseCases struct {
	userRepository repositories.ICreateUserRepository
}

func NewUserUseCase(userRepo repositories.ICreateUserRepository) CreateUserUseCases {
	return CreateUserUseCases{userRepository: userRepo}
}

func (userUseCase *CreateUserUseCases) Execute(data *dtos.CreateUserDTO) (*entities.User, error) {

	existingUser, err := userUseCase.userRepository.FindByEmail(data.Email)

	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.New("Ups, este usuário já existe!")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, errors.New("falha ao criptografar a senha")
	}

	user := &entities.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Phone:     data.Phone,
		Password:  string(hashedPassword),
		Bio:       data.Bio,
		Photo:     nil,
	}

	err = userUseCase.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
