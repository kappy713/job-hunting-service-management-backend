package entity

import "github.com/google/uuid"

type Mynavi struct {
	ID            uuid.UUID `gorm:"type:uuid;primarykey" json:"id"`
	SelfPromotion string    `gorm:"size:1000" json:"self_promotion"`
	FuturePlan    string    `gorm:"size:300" json:"future_plan"`
}

// リクエスト用の構造体
type MynaviData struct {
	SelfPromotion string `json:"self_promotion"`
	FuturePlan    string `json:"future_plan"`
}

type CreateMynaviRequest struct {
	UserID string     `json:"user_id" binding:"required"`
	Data   MynaviData `json:"data" binding:"required"`
}

func (Mynavi) TableName() string {
	return "mynavi"
}
