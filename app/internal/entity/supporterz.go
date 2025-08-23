package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Supporterz struct {
	ID                           uuid.UUID      `gorm:"type:uuid;primarykey" json:"id"`
	CareerVision                 string         `gorm:"size:200" json:"career_vision"`
	SelfPromotion                string         `gorm:"size:5000" json:"self_promotion"`
	Skills                       pq.StringArray `gorm:"type:text[]" json:"skills"`
	SkillDescriptions            pq.StringArray `gorm:"type:text[]" json:"skill_descriptions"`
	InternExperiences            pq.StringArray `gorm:"type:text[]" json:"intern_experiences"`
	InternExperienceDescriptions pq.StringArray `gorm:"type:text[]" json:"intern_experience_descriptions"`
	Products                     pq.StringArray `gorm:"type:text[]" json:"products"`
	ProductTechStacks            pq.StringArray `gorm:"type:text[]" json:"product_tech_stacks"`
	ProductDescriptions          pq.StringArray `gorm:"type:text[]" json:"product_descriptions"`
	Researches                   pq.StringArray `gorm:"type:text[]" json:"researches"`
	ResearchDescriptions         pq.StringArray `gorm:"type:text[]" json:"research_descriptions"`
}

// リクエスト用の構造体
type SupporterzData struct {
	CareerVision                 string   `json:"career_vision"`
	SelfPromotion                string   `json:"self_promotion"`
	Skills                       []string `json:"skills"`
	SkillDescriptions            []string `json:"skill_descriptions"`
	InternExperiences            []string `json:"intern_experiences"`
	InternExperienceDescriptions []string `json:"intern_experience_descriptions"`
	Products                     []string `json:"products"`
	ProductTechStacks            []string `json:"product_tech_stacks"`
	ProductDescriptions          []string `json:"product_descriptions"`
	Researches                   []string `json:"researches"`
	ResearchDescriptions         []string `json:"research_descriptions"`
}

type CreateSupporterzRequest struct {
	UserID string         `json:"user_id" binding:"required"`
	Data   SupporterzData `json:"data" binding:"required"`
}

func (Supporterz) TableName() string {
	return "supporterz"
}
