package tool

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint           `gorm:"primaryKey" json:"id,omitempty" binding:"-"`
	CreatedAt time.Time      `gorm:"" json:"created_at,omitempty" binding:"-"`
	UpdatedAt time.Time      `gorm:"" json:"updated_at,omitempty" binding:"-"`
	DeletedAt gorm.DeletedAt `gorm:"" json:"deleted_at,omitempty" binding:"-"`
}
