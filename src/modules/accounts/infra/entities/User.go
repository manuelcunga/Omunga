package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         string  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	FIRST_NAME string  `gorm:"type:varchar(100) not null" json:"first_name"`
	LAST_NAME  string  `gorm:"type:varchar(100) not null" json:"last_name"`
	EMAIL      string  `gorm:"type:varchar(100) not null" json:"email"`
	PHONE      string  `gorm:"type:varchar(100) not null" json:"phone"`
	PASSWORD   string  `gorm:"type:varchar(100) not null" json:"password"`
	BIO        string  `gorm:"type:varchar(100) not null" json:"bio"`
	PHOTO      *string `gorm:"type:varchar(100) " json:"photo"`

	Base
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New().String()

	u.ID = uuid

	return nil
}
