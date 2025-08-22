package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	UserID        uuid.UUID      `gorm:"type:uuid;primarykey" json:"user_id"`
	LastName      string         `gorm:"not null;size:50" json:"last_name"`
	FirstName     string         `gorm:"not null;size:50" json:"first_name"`
	BirthDate     *time.Time     `gorm:"type:date" json:"birth_date,omitempty"`
	Age           int            `gorm:"check:age >= 0 AND age <= 150" json:"age"`
	University    string         `gorm:"size:50" json:"university"`
	Category      string         `gorm:"size:100" json:"category"`
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

// リクエスト用の構造体
type UserData struct {
	LastName      string     `json:"last_name"`
	FirstName     string     `json:"first_name"`
	BirthDate     *time.Time `json:"birth_date,omitempty"`
	Age           int        `json:"age" binding:"min=0,max=150"`
	University    string     `json:"university"`
	Category      string     `json:"category"`
	Faculty       string     `json:"faculty"`
	Grade         *int       `json:"grade,omitempty" binding:"omitempty,min=1,max=10"`
	TargetJobType string     `json:"target_job_type"`
	Services      []string   `json:"services"`
}

type UpdateUserRequest struct {
	UserID string   `json:"user_id" binding:"required"`
	Data   UserData `json:"data" binding:"required"`
}
