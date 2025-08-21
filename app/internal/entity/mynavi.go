package entity

import "github.com/google/uuid"

type Mynavi struct {
	ID            uuid.UUID `gorm:"type:uuid;primarykey" json:"id"`
	SelfPromotion string    `gorm:"size:1000" json:"self_promotion"`
	FuturePlan    string    `gorm:"size:300" json:"future_plan"`
}

func (Mynavi) TableName() string {
	return "mynavi"
}
