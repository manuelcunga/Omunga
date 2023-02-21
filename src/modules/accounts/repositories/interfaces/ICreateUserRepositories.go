package repositories

import (
	"github.com/manuelcunga/Omunga/src/modules/accounts/infra/entities"
)

type ICreateUserRepository interface {
	Create(data *entities.User) error
	FindAll() ([]*entities.User, error)
}
