package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	// SupabaseのAuth IDを設定（自動生成なし）
	UserID        uuid.UUID      `gorm:"type:uuid;primarykey" json:"user_id"`
	LastName      string         `gorm:"not null;size:50" json:"last_name"`
	FirstName     string         `gorm:"not null;size:50" json:"first_name"`
	Age           int            `gorm:"check:age >= 0 AND age <= 150" json:"age"`
	University    string         `gorm:"size:50" json:"university"`
	Faculty       string         `gorm:"size:50" json:"faculty"`
	Grade         int            `gorm:"check:grade >= 1 AND grade <= 10" json:"grade"`
	TargetJobType string         `gorm:"not null;size:50" json:"target_job_type"`
	Services      pq.StringArray `gorm:"type:text[]" json:"services"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
