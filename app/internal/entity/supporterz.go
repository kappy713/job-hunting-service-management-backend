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

func (Supporterz) TableName() string {
	return "supporterz"
}
