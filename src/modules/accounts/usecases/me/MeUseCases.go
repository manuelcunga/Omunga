package usecases

import (
	"fmt"

	"github.com/manuelcunga/Omunga/src/modules/accounts/dtos"
	repositories "github.com/manuelcunga/Omunga/src/modules/accounts/repositories/interfaces"
)

type MeUseCases struct {
	userRepository repositories.ICreateUserRepository
}

func NewMeUseCases(userRepo repositories.ICreateUserRepository) MeUseCases {
	return MeUseCases{userRepository: userRepo}
}

func (meUseCases *MeUseCases) Me(userID string) (*dtos.ProfileDTO, error) {
	user, err := meUseCases.userRepository.FindById(userID)
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
