package repositories

import (
	"github.com/manuelcunga/Omunga/src/modules/accounts/infra/entities"
	repositories "github.com/manuelcunga/Omunga/src/modules/accounts/repositories/interfaces"
	"gorm.io/gorm"
)

type userRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.ICreateUserRepository {
	return &userRepository{
		Db: db,
	}
}

// Create implements repositories.ICreateUserRepository
func (user *userRepository) Create(data *entities.User) error {
	err := user.Db.Create(&data).Error

	if err != nil {
		return err
	}

	return nil
}

// FindAll implements repositories.ICreateUserRepository
func (users *userRepository) FindAll() ([]*entities.User, error) {
	var allUsers []*entities.User

	err := users.Db.Find(&allUsers).Error

	if err != nil {
		return nil, err
	}

	return allUsers, nil
}
