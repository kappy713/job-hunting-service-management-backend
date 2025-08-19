package entity

import (
	"time"

	"gorm.io/gorm"
)

type SampleUser struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Name      string         `gorm:"not null;size:100" json:"name"`
	Email     string         `gorm:"uniqueIndex;not null;size:255" json:"email"`
	Age       int            `gorm:"check:age >= 0 AND age <= 150" json:"age"`
	IsActive  bool           `gorm:"default:true" json:"is_active"`
	Bio       string         `gorm:"size:500" json:"bio,omitempty"`
	Website   string         `gorm:"size:255" json:"website,omitempty"`
}

func (SampleUser) TableName() string {
	return "sample_users"
}
