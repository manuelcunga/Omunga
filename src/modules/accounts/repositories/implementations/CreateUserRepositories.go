package repositories

import (
	"errors"
	"fmt"

	"github.com/manuelcunga/Omunga/src/modules/accounts/infra/entities"
	repositories "github.com/manuelcunga/Omunga/src/modules/accounts/repositories/interfaces"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.ICreateUserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (user *UserRepository) Create(data *entities.User) error {
	err := user.Db.Create(&data).Error

	if err != nil {
		return err
	}

	return nil
}

func (users *UserRepository) FindAll() ([]*entities.User, error) {
	var allUsers []*entities.User

	err := users.Db.Find(&allUsers).Error

	if err != nil {
		return nil, err
	}

	return allUsers, nil
}

func (userRepo *UserRepository) FindByEmail(email string) (*entities.User, error) {

	user := &entities.User{}

	result := userRepo.Db.Where("email = ?", email).First(user)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {

			return nil, fmt.Errorf("Ups, usuário não encontrado '%s' not found", email)
		}
		return nil, result.Error
	}
	return user, nil
}

func (userRepo *UserRepository) Delete(id string) (error, *entities.User) {

	user := &entities.User{}
	result := userRepo.Db.Where("id = ?", id).Delete(user)
	if result.Error != nil {
		return result.Error, nil
	}
	return nil, user
}

func (userRepo *UserRepository) FindById(id string) (error, *entities.User) {
	user := &entities.User{}
	result := userRepo.Db.Where("id = ?", id).First(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("user with id '%s' not found", id), nil
		}
		return result.Error, nil
	}
	return nil, user
}

func (userRepo *UserRepository) Update(user *entities.User) (error, *entities.User) {
	result := userRepo.Db.Save(user)
	if result.Error != nil {
		return result.Error, nil
	}
	return nil, user
}
