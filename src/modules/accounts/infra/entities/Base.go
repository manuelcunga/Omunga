package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Base struct {
	CreatedAt time.Time `gorm:"autoCreateTime:true" json:"created_at" valid:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:true" json:"updated_at" valid:"-"`
}
