package repositories

import (
	"errors"

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

	var user entities.User
	result := userRepo.Db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (users *UserRepository) FindById(userId string) (*entities.User, error) {
	user := &entities.User{}

	err := users.Db.First(user, "id = ?", userId).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userRepo *UserRepository) Update(userId string, user *entities.User) (error, *entities.User) {

	result := userRepo.Db.Save(user)
	if result.Error != nil {
		return result.Error, nil
	}
	return nil, user
}

func (userRepo *UserRepository) Delete(id string) error {
	result := userRepo.Db.Where("id = ?", id).Delete(&entities.User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
