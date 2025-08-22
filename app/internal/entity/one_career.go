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

func (OneCareer) TableName() string {
	return "one_career"
}
