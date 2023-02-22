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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.PASSWORD), bcrypt.DefaultCost)

	if err != nil {
		return nil, errors.New("failed to encrypt password")
	}
	user := &entities.User{
		FIRST_NAME: data.FIRST_NAME,
		LAST_NAME:  data.LAST_NAME,
		EMAIL:      data.EMAIL,
		PHONE:      data.PHONE,
		PASSWORD:   string(hashedPassword),
		BIO:        data.BIO,
		PHOTO:      nil,
	}

	err = userUseCase.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
