package middlewares

import (
	"fmt"
	"net/http"

	"github.com/manuelcunga/Omunga/src/modules/accounts/infra/entities"
	repositories "github.com/manuelcunga/Omunga/src/modules/accounts/repositories/interfaces"
)

type CreateUserRepository struct {
	userRepository repositories.ICreateUserRepository
}

func (userRepo *CreateUserRepository) GetCurrentUser(w http.ResponseWriter, r *http.Request) (*entities.User, error) {
	// Obter o ID do usuário atual a partir do contexto da requisição
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		return nil, fmt.Errorf("User ID not found in context")
	}

	// Buscar as informações completas do usuário correspondente no banco de dados
	// user, err := userRepo.userRepository.FindById(userID)
	// if err != nil {
	// 	return err, nil
	// }

	// return nil, user
	foundUser, err := userRepo.userRepository.FindById(userID)
	if err != nil {
		return nil, err
	}
	return foundUser, nil
}
