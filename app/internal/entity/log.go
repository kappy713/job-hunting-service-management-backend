package entity

import (
	"time"

	"github.com/google/uuid"
)

type Log struct {
	ID          uuid.UUID `gorm:"type:uuid;primarykey" json:"id"`
	TargetTable string    `gorm:"size:100" json:"target_table"`
	FieldName   string    `gorm:"size:100" json:"field_name"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Log) TableName() string {
	return "logs"
}
