package repositories

import (
	"github.com/manuelcunga/Omunga/src/modules/accounts/infra/entities"
)

type ICreateUserRepository interface {
	Create(data *entities.User) error
	FindAll() ([]*entities.User, error)
	FindById(id string) (error, *entities.User)
	FindByEmail(email string) (*entities.User, error)
	Update(user *entities.User) (error, *entities.User)
	Delete(id string) (error, *entities.User)
}
