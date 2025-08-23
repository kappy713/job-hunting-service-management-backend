package entity

import "github.com/google/uuid"

// Mynavi マイナビ用のプロフィール情報
type Mynavi struct {
	ID            uuid.UUID `gorm:"type:uuid;primarykey" json:"id"`  // ユーザーID（主キー）
	SelfPromotion string    `gorm:"size:1000" json:"self_promotion"` // 自己PR文
	FuturePlan    string    `gorm:"size:300" json:"future_plan"`     // 将来の計画・目標
}

// MynaviData マイナビ用のリクエストデータ構造体
type MynaviData struct {
	SelfPromotion string `json:"self_promotion"` // 自己PR文
	FuturePlan    string `json:"future_plan"`    // 将来の計画・目標
}

// CreateMynaviRequest マイナビ作成リクエスト
type CreateMynaviRequest struct {
	UserID string     `json:"user_id" binding:"required"` // ユーザーID
	Data   MynaviData `json:"data" binding:"required"`    // マイナビデータ
}

func (Mynavi) TableName() string {
	return "mynavi"
}
