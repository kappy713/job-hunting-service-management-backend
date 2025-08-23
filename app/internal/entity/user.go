package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// User ユーザー情報
type User struct {
	UserID        uuid.UUID      `gorm:"type:uuid;primarykey" json:"user_id"`           // ユーザーID（主キー）
	LastName      string         `gorm:"not null;size:50" json:"last_name"`             // 姓
	FirstName     string         `gorm:"not null;size:50" json:"first_name"`            // 名
	BirthDate     *time.Time     `gorm:"type:date" json:"birth_date,omitempty"`         // 生年月日
	Age           int            `gorm:"check:age >= 0 AND age <= 150" json:"age"`      // 年齢
	University    string         `gorm:"size:50" json:"university"`                     // 大学名
	Category      string         `gorm:"size:100" json:"category"`                      // カテゴリ（学部系統など）
	Faculty       string         `gorm:"size:50" json:"faculty"`                        // 学部名
	Grade         int            `gorm:"check:grade >= 1 AND grade <= 10" json:"grade"` // 学年
	TargetJobType string         `gorm:"not null;size:50" json:"target_job_type"`       // 志望職種
	Services      pq.StringArray `gorm:"type:text[]" json:"services"`                   // 利用する就活サービス一覧
	CreatedAt     time.Time      `json:"created_at"`                                    // 作成日時
	UpdatedAt     time.Time      `json:"updated_at"`                                    // 更新日時
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

// CreateUser専用の構造体（services以外の項目）
type CreateUserData struct {
	LastName      string     `json:"last_name" binding:"required"`
	FirstName     string     `json:"first_name" binding:"required"`
	BirthDate     *time.Time `json:"birth_date,omitempty"`
	Age           int        `json:"age" binding:"required,min=0,max=150"`
	University    string     `json:"university" binding:"required"`
	Category      string     `json:"category" binding:"required"`
	Faculty       string     `json:"faculty" binding:"required"`
	Grade         int        `json:"grade" binding:"required,min=1,max=10"`
	TargetJobType string     `json:"target_job_type" binding:"required"`
}

type UpdateUserRequest struct {
	UserID string   `json:"user_id" binding:"required"`
	Data   UserData `json:"data" binding:"required"`
}

type CreateUserRequest struct {
	UserID string         `json:"user_id" binding:"required"`
	Data   CreateUserData `json:"data" binding:"required"`
}
