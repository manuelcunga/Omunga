package entities

import "time"

type Base struct {
	CreatedAt time.Time `gorm:"autoCreateTime:false" json:"created_at" valid:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:false" json:"updated_at" valid:"-"`
}
