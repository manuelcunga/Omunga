package usecases

import (
	"fmt"

	"github.com/manuelcunga/Omunga/src/modules/accounts/dtos"
	repositories "github.com/manuelcunga/Omunga/src/modules/accounts/repositories/interfaces"
)

type ProfileUseCases struct {
	userRepository repositories.ICreateUserRepository
}

func NewProfileUseCases(userRepo repositories.ICreateUserRepository) ProfileUseCases {
	return ProfileUseCases{userRepository: userRepo}
}

func (profileUseCases *ProfileUseCases) Profile(userID string) (*dtos.ProfileDTO, error) {
	user, err := profileUseCases.userRepository.FindById(userID)

	if err != nil {
		return nil, fmt.Errorf("Não foi possível encontrar o usuário")
	}

	return &dtos.ProfileDTO{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Bio:       user.Bio,
		Photo:     user.Photo,
	}, nil
}
