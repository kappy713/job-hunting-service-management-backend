package entity

import (
	"time"

	"github.com/google/uuid"
)

// Log ログ情報（各サービスのフィールド更新履歴）
type Log struct {
	ID          uuid.UUID `gorm:"type:uuid;primarykey" json:"id"`     // ログID（主キー）
	UserID      uuid.UUID `gorm:"type:uuid" json:"user_id"`           // ユーザーID
	TargetTable string    `gorm:"size:100" json:"target_table"`       // 対象テーブル名（どのサービスか）
	FieldName   string    `gorm:"size:100" json:"field_name"`         // 更新されたフィールド名
	UpdatedAt   time.Time `gorm:"type:timestamptz" json:"updated_at"` // 更新日時
}

func (Log) TableName() string {
	return "logs"
}

// LogResponse ログレスポンス用の構造体（サービス名 → フィールド名 → 更新日時のマップ）
type LogResponse map[string]map[string]time.Time
