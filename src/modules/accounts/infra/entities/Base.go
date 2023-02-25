package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// func (Base) init() {
// 	govalidator.CustomTypeTagMap.Set("time.Time", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
// 		_, ok := i.(time.Time)
// 		return ok
// 	}))
// }

type Base struct {
	CreatedAt time.Time `gorm:"autoCreateTime:true" json:"created_at" valid:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:true" json:"updated_at" valid:"-"`
}
