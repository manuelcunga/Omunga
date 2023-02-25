package entities

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type User struct {
	ID        string  `gorm:"type:varchar(36);primary_key;default:gen_random_uuid()" valid:"-"`
	FirstName string  `json:"first_name"   gorm:"type:varchar(255)" valid:"required~O campo First Name é obrigatório"`
	LastName  string  `json:"last_name"  gorm:"type:varchar(255);" valid:"required~O campo Last Name é obrigatório"`
	Email     string  `json:"email" gorm:"type:varchar(100); unique_index"  valid:"required~O campo Email é obrigatório"`
	Phone     string  `json:"phone" gorm:"type:varchar(100);"  valid:"required~O campo Phone é obrigatório"`
	Password  string  `json:"password" gorm:"type:varchar(100)"  valid:"required~O campo Password é obrigatório"`
	Bio       string  `json:"bio"  gorm:"type:varchar(255)"   valid:"required~O campo Bio é obrigatório"`
	Photo     *string `json:"photo" gorm:"type:varchar(255)" valid:"-"`
	Base
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.New()
	user.ID = uuid.String()

	err = user.validate()
	if err != nil {
		return err
	}

	return

}

func (user *User) validate() error {

	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		return err
	}

	return nil
}
