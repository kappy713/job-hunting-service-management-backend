package entity

import (
	"time"

	"github.com/google/uuid"
)

type Log struct {
	ID          uuid.UUID `gorm:"type:uuid;primarykey" json:"id"`
	UserID      uuid.UUID `gorm:"type:uuid" json:"user_id"`
	TargetTable string    `gorm:"size:100" json:"target_table"`
	FieldName   string    `gorm:"size:100" json:"field_name"`
	UpdatedAt   time.Time `gorm:"type:timestamptz" json:"updated_at"`
}

func (Log) TableName() string {
	return "logs"
}

// レスポンス用の構造体
type LogResponse map[string]map[string]time.Time
