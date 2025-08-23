package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type OneCareer struct {
	ID                           uuid.UUID      `gorm:"type:uuid;primarykey" json:"id"`
	Skills                       pq.StringArray `gorm:"type:text[]" json:"skills"`
	SkillDescriptions            pq.StringArray `gorm:"type:text[]" json:"skill_descriptions"`
	Researches                   pq.StringArray `gorm:"type:text[]" json:"researches"`
	ResearchDescriptions         pq.StringArray `gorm:"type:text[]" json:"research_descriptions"`
	InternExperiences            pq.StringArray `gorm:"type:text[]" json:"intern_experiences"`
	InternExperienceDescriptions pq.StringArray `gorm:"type:text[]" json:"intern_experience_descriptions"`
	Products                     pq.StringArray `gorm:"type:text[]" json:"products"`
	ProductDescriptions          pq.StringArray `gorm:"type:text[]" json:"product_descriptions"`
	EngineerAspiration           string         `gorm:"size:1000" json:"engineer_aspiration"`
}

// リクエスト用の構造体
type OneCareerData struct {
	Skills                       []string `json:"skills"`
	SkillDescriptions            []string `json:"skill_descriptions"`
	Researches                   []string `json:"researches"`
	ResearchDescriptions         []string `json:"research_descriptions"`
	InternExperiences            []string `json:"intern_experiences"`
	InternExperienceDescriptions []string `json:"intern_experience_descriptions"`
	Products                     []string `json:"products"`
	ProductDescriptions          []string `json:"product_descriptions"`
	EngineerAspiration           string   `json:"engineer_aspiration"`
}

type CreateOneCareerRequest struct {
	UserID string        `json:"user_id" binding:"required"`
	Data   OneCareerData `json:"data" binding:"required"`
}

func (OneCareer) TableName() string {
	return "one_career"
}
